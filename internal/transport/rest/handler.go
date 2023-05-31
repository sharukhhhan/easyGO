package rest

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
	}

	users := router.Group("/users")
	{
		users.GET("/", h.getAllUsers)
		users.GET("/:id", h.getUserById)
		users.PUT("/:id", h.updateUser)
		users.DELETE("/:id", h.deleteUser)
	}

	cars := router.Group("/cars")
	{
		cars.POST("/", h.createCar)
		cars.GET("/", h.getAllCars)
		cars.GET("/:id", h.getCarById)
		cars.PUT("/:id", h.updateCar)
		cars.DELETE("/:id", h.deleteCar)
	}

	trips := router.Group("/trips")
	{
		trips.POST("/", h.createTrip)
		trips.GET("/", h.getAllTrips)
		trips.GET("/:id", h.getTripById)
		trips.PUT("/:id", h.updateTrip)
		trips.DELETE("/:id", h.deleteTrip)
	}

	return router
}
