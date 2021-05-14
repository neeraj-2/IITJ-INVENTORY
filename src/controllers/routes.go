package controllers

// import (
// 	"github.com/gin-gonic/gin"
// )

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		// if c.Request.Method == "OPTIONS" {
// 		// 	c.AbortWithStatus(204)
// 		// 	return
// 		// }

// 		c.Next()
// 	}
// }

//InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	r := s.Router

	// r.Use(CORSMiddleware())

	api := r.Group("/api")
	api.GET("/societies", s.GetSociety)
	api.Use(SetMiddlewareAuthentication())
	{
		api.GET("/getUsers", s.GetUsers)
		api.POST("/addUser", s.AddUser)
		api.POST("/addItem", s.AddItem)
		api.GET("/getItems", s.GetItems)
		api.GET("/searchItems", s.SearchItems)

	}

	student := r.Group("/student")
	{
		student.POST("/login", s.StudentLogin)
		protected := student.Group("/")
		protected.Use(SetMiddlewareAuthenticationStudent(s.DB))
		{
			protected.GET("/profile", s.GetProfile)
			protected.POST("/requestItem", s.RequestItem)
		}
	}

	society := r.Group("/society")
	{
		society.POST("/login", s.SocietyLogin)
		protected := society.Group("/")
		protected.Use(SetMiddlewareAuthenticationSociety(s.DB))
		{

		}
	}

	admin := r.Group("/admin")
	{
		admin.POST("/login", s.AdminLogin)
		protected := admin.Group("/")
		protected.Use(SetMiddlewareAuthenticationAdmin(s.DB))
		{
			protected.POST("/createSociety", s.CreateSociety)
		}
	}
}
