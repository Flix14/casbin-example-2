package controllers

import (
	"strconv"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

func AllCustomers(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var customers []m.Customer
	var count int64

	db.Scopes(m.Pagination(page, limit)).Find(&customers)
	db.Model(m.Customer{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
		Records:     customers,
	}
	c.JSON(200, paginator)
}

func FindCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer m.Customer
	err := customer.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, customer)
}

func AddCustomer(c *gin.Context) {
	var customer m.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	newcustomer, err := customer.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newcustomer)
}

func UpdCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer m.Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	customer.ID = id
	newcustomer, err := customer.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newcustomer)
}

func DelCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var customer m.Customer
	customer.ID = id
	err := customer.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
