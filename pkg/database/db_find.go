package database

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/models"
)

//func (s *Storage) SearchData(segments []int64) models.ResponsePrice {
//	sqlQuery := fmt.Sprintf("SELECT COALESCE(d.price, b.price")
//}

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
