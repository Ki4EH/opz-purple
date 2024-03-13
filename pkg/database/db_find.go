package database

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/models"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func SearchData(segments []int64, price models.RequestPrice) models.ResponsePrice {
	ansCh := make(chan models.ResponsePrice)
	wg.Add(len(segments))
	for id, v := range segments {
		if id == 0 {
			s := fmt.Sprintf("baseline_matrix_%d", v)
			go Connection.SearchInTable(s, price.LocationId, price.MicrocategoryId, ansCh)
			continue
		}
		s := fmt.Sprintf("discount_matrix_%d", v)
		go Connection.SearchInTable(s, price.LocationId, price.MicrocategoryId, ansCh)
	}
	go func() {
		wg.Wait()
		close(ansCh)
	}()
	var baseline models.ResponsePrice
	var discount models.ResponsePrice
	//TODO: нужно рассмотреть крайние случаи, когда в дисконтах несколько скидок одинаковых (как пример)
	fmt.Println(ansCh)
	for v := range ansCh {
		if v.MatrixId == segments[0] {
			baseline = v
		}
		discount = v
		return discount
	}
	return baseline
}

func (s *Storage) SearchInTable(table string, lc, mc int, ans chan<- models.ResponsePrice) {
	sqlQuery := fmt.Sprintf("SELECT * FROM %s WHERE (location_id = %d AND microcategory_id = %d);", table, lc, mc)
	var mc1, lc1, price1 int
	err := s.db.QueryRow(sqlQuery).Scan(&mc1, &lc1, &price1)
	if err != nil {
		return
	}
	tableSplit := strings.Split(table, "_")
	id, _ := strconv.Atoi(tableSplit[len(tableSplit)-1])
	answer := models.ResponsePrice{
		Price:           int64(price1),
		LocationId:      lc1,
		MicrocategoryId: mc1,
		MatrixId:        int64(id),
	}
	defer wg.Done()
	mu.Lock()
	ans <- answer
	defer mu.Unlock()
}
func (s *Storage) AddNewPrice(data models.RequestAddPrice) error {
	var exist bool

	s.db.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema='public' AND table_name='%s');", data.Matrix)).Scan(&exist)
	if !exist {
		return fmt.Errorf("неверное название матрицы")
	}

	sqlQuery := fmt.Sprintf("INSERT INTO %s (microcategory_id, location_id, price) VALUES (%d, %d, %d);", data.Matrix, data.MicrocategoryId, data.LocationId, data.Price)
	s.db.QueryRow(sqlQuery)
	return nil
}

func (s *Storage) UpdatePrice(data models.RequestAddPrice) error {
	var exist bool

	existChecker := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema='public' AND table_name='%s');", data.Matrix)
	s.db.QueryRow(existChecker).Scan(&exist)
	if exist == false {
		return fmt.Errorf("неверное название матрицы")
	}
	existChecker = fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE microcategory_id=%d AND location_id=%d);", data.Matrix, data.MicrocategoryId, data.LocationId)
	s.db.QueryRow(existChecker).Scan(&exist)
	if exist == false {
		return fmt.Errorf("неверный номер локации или микрокатегории")
	}

	sqlQuery := fmt.Sprintf("UPDATE %s price SET %d WHERE location_id=%d AND microcategory_id=%d;", data.Matrix, data.Price, data.LocationId, data.MicrocategoryId)
	s.db.QueryRow(sqlQuery)
	return nil
}
