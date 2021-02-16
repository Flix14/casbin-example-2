package controllers

import (
	"strconv"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

func AllMonitors(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var monitors []m.Monitor
	var count int64

	db.Scopes(m.Pagination(page, limit)).Find(&monitors)
	db.Model(m.Monitor{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
		Records:     monitors,
	}
	c.JSON(200, paginator)
}

func FindMonitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var monitor m.Monitor
	err := monitor.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, monitor)
}

func AddMonitor(c *gin.Context) {
	var monitor m.Monitor
	err := c.BindJSON(&monitor)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	newmonitor, err := monitor.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newmonitor)
}

func UpdMonitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var monitor m.Monitor
	err := c.BindJSON(&monitor)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	monitor.ID = id
	newmonitor, err := monitor.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newmonitor)
}

func DelMonitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var monitor m.Monitor
	monitor.ID = id
	err := monitor.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
