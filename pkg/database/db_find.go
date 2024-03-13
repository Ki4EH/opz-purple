package database

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/models"
	"github.com/Ki4EH/opz-purple/pkg/treebase/location"
	"github.com/Ki4EH/opz-purple/pkg/treebase/microcategory"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func SearchData(segments []int64, price models.RequestPrice) models.ResponsePrice {
	ansCh := make(chan models.ResponsePrice)

	mc1 := microcategory.GetCategoryParent(price.MicrocategoryId)
	lc1 := location.GetLocationParent(price.LocationId)

	for id, v := range segments {
		s := fmt.Sprintf("baseline_matrix_%d", v)
		if id != 0 {
			s = fmt.Sprintf("discount_matrix_%d", v)
		}

		wg.Add(4)
		go Connection.SearchInTable(s, price.LocationId, price.MicrocategoryId, ansCh)
		go Connection.SearchInTable(s, price.LocationId, mc1, ansCh)
		go Connection.SearchInTable(s, lc1, price.MicrocategoryId, ansCh)
		go Connection.SearchInTable(s, lc1, mc1, ansCh)
	}

	go func() {
		wg.Wait()
		close(ansCh)
	}()

	var baseline models.ResponsePrice
	//var discount models.ResponsePrice

	for v := range ansCh {
		fmt.Println(v)
		//if v.MatrixId == segments[0] {
		//	baseline = v
		//}
		//discount = v
		//return discount
	}

	return baseline
}

func (s *Storage) SearchInTable(table string, lc, mc int64, ans chan<- models.ResponsePrice) {
	defer wg.Done()

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

	ans <- answer
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

func (s *Storage) UpdatePrice(data models.RequestAddPrice) {
	//var exist bool

	//existChecker := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_schema='public' AND table_name='%s');", data.Matrix)
	//s.db.QueryRow(existChecker).Scan(&exist)
	//if exist == false {
	//	return fmt.Errorf("неверное название матрицы")
	//}
	//existChecker = fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE microcategory_id=%d AND location_id=%d);", data.Matrix, data.MicrocategoryId, data.LocationId)
	//s.db.QueryRow(existChecker).Scan(&exist)
	//if exist == false {
	//	return fmt.Errorf("неверный номер локации или микрокатегории")
	//}

	sqlQuery := fmt.Sprintf("UPDATE %s price SET %d WHERE location_id=%d AND microcategory_id=%d;", data.Matrix, data.Price, data.LocationId, data.MicrocategoryId)
	s.db.QueryRow(sqlQuery)
}

func (s *Storage) UpdateManyPrices(data models.RequestWithPercentage) {
	var updater string
	switch {
	case data.LocationId == 0 && data.MicrocategoryId != 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=%d WHERE microcategory_id=%d;", data.Matrix, data.Price, data.MicrocategoryId)
	case data.LocationId != 0 && data.MicrocategoryId == 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=%d WHERE location_id=%d;", data.Matrix, data.Price, data.LocationId)
	case data.LocationId != 0 && data.MicrocategoryId != 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=%d WHERE location_id=%d AND microcategory_id=%d;", data.Matrix, data.Price, data.LocationId, data.MicrocategoryId)
	default:
		updater = fmt.Sprintf("UPDATE %s  SET price=%d;", data.Matrix, data.Price)
	}
	s.db.QueryRow(updater)
}

func (s *Storage) UpdateManyPricesWithPercentage(data models.RequestWithPercentage) {
	var updater string
	switch {
	case data.LocationId == 0 && data.MicrocategoryId != 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=price * %f WHERE microcategory_id=%d;", data.Matrix, 1-data.Percent/100, data.MicrocategoryId)
	case data.LocationId != 0 && data.MicrocategoryId == 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=price * %f WHERE location_id=%d;", data.Matrix, 1-data.Percent/100, data.LocationId)
	case data.LocationId != 0 && data.MicrocategoryId != 0:
		updater = fmt.Sprintf("UPDATE %s  SET price=price * %f WHERE location_id=%d AND microcategory_id=%d;", data.Matrix, 1-data.Percent/100, data.LocationId, data.MicrocategoryId)
	default:
		updater = fmt.Sprintf("UPDATE %s  SET price=price * %f;", data.Matrix, 1-data.Percent/100)
	}
	s.db.QueryRow(updater)
}

func (s *Storage) CreateNewTable(data models.RequestCreate) error {
	creator := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (microcategory_id int PRIMARY KEY, location_id int, price int);", data.Matrix)
	s.db.QueryRow(creator)
	for _, row := range data.Rows {
		err := s.AddNewPrice(models.RequestAddPrice{Matrix: data.Matrix, LocationId: row.LocationId, MicrocategoryId: row.MicrocategoryId, Price: row.Price})
		if err != nil {
			return err
		}
	}
	return nil
}
