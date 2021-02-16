package models

type Customer struct {
	Model
	FirstName string `gorm:"size:250;not null;" sql:"index" json:"first_name"`
	Email     string `gorm:"unique;size:50;" sql:"index" json:"email"`
}

func (c *Customer) Add() (*Customer, error) {
	err := DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (c *Customer) Find(id int) (err error) {
	err = DB.First(&c, id).Error
	if err != nil {
		return
	}
	return
}

func (c *Customer) Update() (*Customer, error) {
	var cc Customer
	err := DB.First(&cc, c.ID).Error
	if err != nil {
		return nil, err
	}
	err = DB.Save(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (c *Customer) Remove() (err error) {
	err = DB.First(&c, c.ID).Error
	if err != nil {
		return
	}
	err = DB.Delete(&c).Error
	if err != nil {
		return
	}
	return
}
