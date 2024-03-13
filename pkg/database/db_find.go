package database

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/models"
)

//func (s *Storage) SearchData(segments []int64) models.ResponsePrice {
//	sqlQuery := fmt.Sprintf("SELECT COALESCE(d.price, b.price")
//}

func (s *Storage) AddNewPrice(data models.RequestAddPrice) error {
	//sqlStatement := ``
	//s.db.QueryRow("SELECT * FROM p;").Scan(&la)
	//
	fmt.Println(s.db.Ping())
	sqlQuery := fmt.Sprintf("INSERT INTO %s (microcategory_id, location_id, price) VALUES (%d, %d, %d);", "discount_matrix_3", data.MicrocategoryId, data.LocationId, data.Price)
	s.db.QueryRow(sqlQuery)
	defer s.db.Close()
	return nil
}
