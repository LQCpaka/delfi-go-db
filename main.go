package main

import (
	"log"

	"delfi-scanner-api/db"
	"delfi-scanner-api/handlers"
	"delfi-scanner-api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var db = make(map[string]string)
//
// func setupRouter() *gin.Engine {
// 	r := gin.Default()
//
// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})
//
// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := db[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})
//
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))
//
// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)
//
// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}
//
// 		if c.Bind(&json) == nil {
// 			db[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})
//
// 	return r
// }

func setupRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db.ConnectDatabase()

	log.Println("Running Migration...")
	err = db.DB.AutoMigrate(&models.User{}, &models.Ticket{}, &models.Product{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// User routes
		users := api.Group("/user")
		{
			users.POST("/signup", handlers.SignUp)
			users.POST("/signin", handlers.SignIn)
		}

		// Ticket routes
		tickets := api.Group("/ticket")
		{
			tickets.POST("/", handlers.CreateTicket) // Tạo ticket mới
			tickets.GET("/", handlers.GetTickets)    // Lấy danh sách tất cả ticket
			// tickets.GET("/:id", handlers.GetTicketByID)      // Lấy một ticket theo ID (kèm sản phẩm)
			tickets.PUT("/:id", handlers.UpdateTicketStatus) // Cập nhật trạng thái ticket
			tickets.DELETE("/:id", handlers.DeleteTicket)
		}

		// Product routes
		products := api.Group("/products")
		{
			// Route này chỉ là ví dụ, bạn có thể tạo thêm các route khác
			products.POST("/", handlers.AddProductToTicket) // Thêm sản phẩm mới vào một ticket
			products.PUT("/:id", handlers.UpdateProduct)    // Cập nhật một sản phẩm
			products.DELETE("/:id", handlers.DeleteProduct) // Xóa một sản phẩm
		}

	}

	// =========================| RUN SERVER |========================
	log.Println("Starting server on port 8080...")
	r.Run(":8080")
}

func main() {
	setupRouter()
}
