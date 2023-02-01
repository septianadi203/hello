package usersdto

type RequestUser struct {
	
	Fullname   string `json:"fullname" gorm:"type : varchar(255)" validate:"required"`
	Username   string `json:"username" gorm:"type : varchar(255)" validate:"required"`
	Email      string `json:"email" gorm:"type : varchar(255)" validate:"required"`
	Password   string `json:"password" gorm:"type : varchar(255)" validate:"required"`
	ListAsRole string `json:"listasrole" gorm:"type : varchar(255)" validate:"required"`
	Gender     string `json:"gender" gorm:"type : varchar(255)" validate:"required"`
	Phone      string `json:"phone" gorm:"type : varchar(255)" validate:"required"`
	Address    string `json:"address" gorm:"type : varchar(255)" validate:"required"`
	Image      string `json:"image" gorm:"type : varchar(255)" validate:"required"`
}

