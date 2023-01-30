package authdto

type SignUpResponse struct {
	Username string `json:"username" gorm:"type : varchar(255)"`
	Message  string `json:"message" gorm:"type : varchar(255)"`
	Password string `json:"password"`
	ListAsRole string `json:"listasrole"`
}

type SignInResponse struct {
	Username string `json:"username" gorm:"type : varchar(255)"`
	Token    string `json:"token" gorm:"type : varchar(255)"`
}
