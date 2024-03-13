package database

import (
	"github.com/Ki4EH/opz-purple/internal/models"
)

//func (s *Storage) SearchData(segments []int64) models.ResponsePrice {
//	sqlQuery := fmt.Sprintf("SELECT COALESCE(d.price, b.price")
//}

func (s *Storage) AddNewPrice(data models.RequestAddPrice) error {
	//sqlStatement := ``
	//var la []string
	//s.db.QueryRow("SELECT * FROM p;").Scan(&la)
	//fmt.Println(la)

	//var a interface{}
	//
	//s.db.QueryRow("INSERT INTO discount_matrix_3 (microcategory_id, location_id, price) VALUES (?, ?, ?);", data.MicrocategoryId, data.LocationId, data.Price).Scan(&a)
	//fmt.Println("ура")
	//fmt.Println(a)
	//defer s.db.Close()
	return nil
}
