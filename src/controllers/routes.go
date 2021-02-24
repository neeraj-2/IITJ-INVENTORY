package controllers

//InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	r := s.Router

	api := r.Group("/api")
	api.Use(SetMiddlewareAuthentication())
	{
		api.GET("/getHome", s.GetHome)
		api.POST("/addUser", s.AddUser)
	}

	auth := r.Group("/auth")
	{
		auth.GET("/login", s.Login)
		auth.GET("/callback", s.Callback)
		auth.GET("error", s.Error)
		auth.POST("/admin/login", s.AdminLogin)
	}
}
