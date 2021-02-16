package models

type Table struct {
	Model
	Name string `gorm:"size:250;not null;" sql:"index" json:"name"`
	Zone string `gorm:"size:250;" sql:"index" json:"zone"`
}

func (c *Table) Add() (*Table, error) {
	err := DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (c *Table) Find(id int) (err error) {
	err = DB.First(&c, id).Error
	if err != nil {
		return
	}
	return
}

func (c *Table) Update() (*Table, error) {
	var cc Table
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

func (c *Table) Remove() (err error) {
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
