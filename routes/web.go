package routes

import (
	"github.com/Flix14/casbin-example-2/controllers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	auth := r.Group("/auth")
	{
		auth.Any("/signin", controllers.Signin)
		auth.Any("/logout", controllers.Logout)
	}

	customer := r.Group("/customer").Use(VerifyJWT())
	{
		customer.GET("", controllers.AllCustomers)
		customer.GET("/:id", controllers.FindCustomer)
		customer.POST("", controllers.AddCustomer)
		customer.PUT("/:id", controllers.UpdCustomer)
		customer.DELETE("/:id", controllers.DelCustomer)
	}

	cashbox := r.Group("/cashbox").Use(VerifyJWT())
	{
		cashbox.GET("", controllers.AllCashboxes)
		cashbox.GET("/:id", controllers.FindCashbox)
		cashbox.POST("", controllers.AddCashbox)
		cashbox.PUT("/:id", controllers.UpdCashbox)
		cashbox.DELETE("/:id", controllers.DelCashbox)
	}

	table := r.Group("/table").Use(VerifyJWT())
	{
		table.GET("", controllers.AllTables)
		table.GET("/:id", controllers.FindTable)
		table.POST("", controllers.AddTable)
		table.PUT("/:id", controllers.UpdTable)
		table.DELETE("/:id", controllers.DelTable)
	}

	monitor := r.Group("/monitor").Use(VerifyJWT())
	{
		monitor.GET("", controllers.AllMonitors)
		monitor.GET("/:id", controllers.FindMonitor)
		monitor.POST("", controllers.AddMonitor)
		monitor.PUT("/:id", controllers.UpdMonitor)
		monitor.DELETE("/:id", controllers.DelMonitor)
	}
	return r
}
