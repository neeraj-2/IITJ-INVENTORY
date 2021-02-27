package controllers

//InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	r := s.Router

	api := r.Group("/api")
	api.Use(SetMiddlewareAuthentication())
	{
		api.GET("/getUsers", s.GetUsers)
		api.POST("/addUser", s.AddUser)
		api.POST("/addItem", s.AddItem)
		api.GET("/getItems", s.GetItems)

	}

	auth := r.Group("/auth")
	{
		auth.GET("/login", s.Login)
		auth.GET("/callback", s.Callback)
		auth.GET("error", s.Error)
		auth.POST("/admin/login", s.AdminLogin)
	}

	admin := r.Group("/admin")
	{
		admin.POST("/createSociety", s.CreateSociety)
	}
}
