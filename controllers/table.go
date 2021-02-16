package controllers

import (
	"strconv"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

func AllTables(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var tables []m.Table
	var count int64

	db.Scopes(m.Pagination(page, limit)).Find(&tables)
	db.Model(m.Table{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
		Records:     tables,
	}
	c.JSON(200, paginator)
}

func FindTable(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var table m.Table
	err := table.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, table)
}

func AddTable(c *gin.Context) {
	var table m.Table
	err := c.BindJSON(&table)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	newtable, err := table.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newtable)
}

func UpdTable(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var table m.Table
	err := c.BindJSON(&table)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	table.ID = id
	newtable, err := table.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newtable)
}

func DelTable(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var table m.Table
	table.ID = id
	err := table.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
