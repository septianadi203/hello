package models

type ListAs struct {
	ID   int    `json:"id"`
	Role string `json:"role" grom:"type: varchar(255);primaryKey"`
}

type ListAsResponse struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

func (ListAsResponse) TableName() string {
	return "list_as"
}
