package controllers

//InitializeRoutes ...
func (s *Server) InitializeRoutes() {
	r := s.Router

	api := r.Group("/api")
	{
		api.GET("/getHome", s.GetHome)
		api.POST("/checkEmailExists", s.CheckIfEmailExists)
	}
}
