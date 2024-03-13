package models

type RequestPrice struct {
	LocationId      int `json:"location_id"`
	MicrocategoryId int `json:"microcategory_id"`
	UserId          int `json:"user_id"`
}

type ResponsePrice struct {
	Price           int64 `json:"price"`
	LocationId      int   `json:"location_id"`
	MicrocategoryId int   `json:"microcategory_id"`
	MatrixId        int   `json:"matrix_id"`
	UserSegmentId   int64 `json:"user_segment_id"`
}

type RequestAddPrice struct {
	Matrix          string `json:"matrix_name"`
	LocationId      int    `json:"location_id"`
	MicrocategoryId int    `json:"microcategory_id"`
	Price           int    `json:"price"`
}
