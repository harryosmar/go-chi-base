package entities

type Account struct {
	Id       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
	FullName string `json:"full_name" gorm:"column:full_name"`
	Password string `json:"password" gorm:"column:password"`
}

func (Account) TableName() string {
	return "account"
}

func (Account) PrimaryKey() string {
	return "id"
}
