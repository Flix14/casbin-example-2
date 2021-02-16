package models

type Monitor struct {
	Model
	Name string `gorm:"size:150;" json:"name"`
}

func (c *Monitor) Add() (*Monitor, error) {
	err := DB.Create(&c).Error
	if err != nil {
		return nil, err
	}
	return c, err
}

func (c *Monitor) Find(id int) (err error) {
	err = DB.First(&c, id).Error
	if err != nil {
		return
	}
	return
}

func (c *Monitor) Update() (*Monitor, error) {
	var cc Monitor
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

func (c *Monitor) Remove() (err error) {
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
