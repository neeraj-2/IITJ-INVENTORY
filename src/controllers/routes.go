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

	student := r.Group("/student")
	{
		student.POST("/login", s.StudentLogin)
	}

	society := r.Group("/society")
	{
		society.POST("/login", s.SocietyLogin)
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
