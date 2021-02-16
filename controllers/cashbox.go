package controllers

import (
	"strconv"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

func AllCashboxes(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var cashboxs []m.Cashbox
	var count int64

	db.Scopes(m.Pagination(page, limit)).Find(&cashboxs)
	db.Model(m.Cashbox{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
		Records:     cashboxs,
	}
	c.JSON(200, paginator)
}

func FindCashbox(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cashbox m.Cashbox
	err := cashbox.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, cashbox)
}

func AddCashbox(c *gin.Context) {
	var cashbox m.Cashbox
	err := c.BindJSON(&cashbox)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	newcashbox, err := cashbox.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newcashbox)
}

func UpdCashbox(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cashbox m.Cashbox
	err := c.BindJSON(&cashbox)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	cashbox.ID = id
	newcashbox, err := cashbox.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newcashbox)
}

func DelCashbox(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var cashbox m.Cashbox
	cashbox.ID = id
	err := cashbox.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
