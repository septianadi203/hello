package usersdto

type UserResponse struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ListAsRole string `json:"listasrole" gorm:"type: varchar(225)"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Image      string `json:"image"`
}
