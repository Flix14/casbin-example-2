package models

import "gorm.io/gorm"

type User struct {
	ID   int    `gorm:"primarykey;" json:"id"`
	Name string `gorm:"size:250;not null;" sql:"index" json:"name"`
	Role string `gorm:"size:250;not null;" json:"role"`
}

func CreateUsers(db *gorm.DB) {
	db.FirstOrCreate(&User{ID: 1, Name: "Felix", Role: "admin"})
	db.FirstOrCreate(&User{ID: 2, Name: "Maribel", Role: "cashier"})
	db.FirstOrCreate(&User{ID: 3, Name: "Josue", Role: "waiter"})
	db.FirstOrCreate(&User{ID: 4, Name: "Luis", Role: "chef"})
}

func (u *User) VerifyCredentials(name string) (pu *User, err error) {
	err = DB.Where("name=?", name).Find(&u).Error
	if err != nil {
		return
	}
	pu = u
	return
}
