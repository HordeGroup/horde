package model

type User struct {
	Model
	Id        uint32 `gorm:"column:id;primary_key;auto_increment"`
	Name      string `gorm:"column:name;type:char(16);not null"`
	Password  string `gorm:"column:password;type:varchar(256);not null"`
	Email     string `gorm:"column:email;type:varchar(255)"`
	Telephone string `gorm:"column:telephone;type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}

func (User) UniqueIndexes() map[string][]string {
	return map[string][]string{
		"name": {"name", "deleted"},
	}
}

func (User) Indexes() map[string][]string {
	return map[string][]string{
		"email":     {"email"},
		"telephone": {"telephone"},
	}
}
