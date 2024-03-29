package models

type RequestPrice struct {
	LocationId      int64 `json:"location_id"`
	MicrocategoryId int64 `json:"microcategory_id"`
	UserId          int64 `json:"user_id"`
}

type ResponsePrice struct {
	Price           int64 `json:"price"`
	LocationId      int   `json:"location_id"`
	MicrocategoryId int   `json:"microcategory_id"`
	MatrixId        int64 `json:"matrix_id"`
	UserSegmentId   int64 `json:"user_segment_id"`
}

type RequestAddPrice struct {
	Matrix          string `json:"matrix_name"`
	LocationId      int    `json:"location_id"`
	MicrocategoryId int    `json:"microcategory_id"`
	Price           int    `json:"price"`
}

type RequestWithPercentage struct {
	Matrix          string  `json:"matrix_name"`
	LocationId      int     `json:"location_id"`
	MicrocategoryId int     `json:"microcategory_id"`
	Price           int     `json:"price"`
	Percent         float64 `json:"percent"`
}

type Row struct {
	LocationId      int `json:"location_id"`
	MicrocategoryId int `json:"microcategory_id"`
	Price           int `json:"price"`
}

type RequestCreate struct {
	Matrix string `json:"matrix_name"`
	Rows   []Row  `json:"rows"`
}

type TableName struct {
	TableNameArr []string `json:"tables"`
}

type TableId struct {
	ID string `json:"id"`
}
