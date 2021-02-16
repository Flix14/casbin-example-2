package models

type Cashbox struct {
	Model
	Total float32 `gorm:"type:decimal(16,2);default:0;" json:"total"`
}

func (c *Cashbox) Add() (*Cashbox, error) {
	err := DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (c *Cashbox) Find(id int) (err error) {
	err = DB.First(&c, id).Error
	if err != nil {
		return
	}
	return
}

func (c *Cashbox) Update() (*Cashbox, error) {
	var cc Cashbox
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

func (c *Cashbox) Remove() (err error) {
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
