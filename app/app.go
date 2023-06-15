package app

import (
	"ticket/controllers/flightcontroller"
	"ticket/controllers/usercontroller"
	"ticket/repository/flightRepository"
	"ticket/repository/userRepository"
	"ticket/services/flightService"
	"ticket/services/userService"

	"github.com/labstack/echo"
	_ "gorm.io/driver/mysql"
)

type App struct {
	E *echo.Echo
}

func NewApp() *App {
	e := echo.New()
	routing(e)
	return &App{
		E: e,
	}
}

func (a *App) Start(addr string) error {
	a.E.Logger.Fatal(a.E.Start(addr))
	return nil
}

func routing(e *echo.Echo) {
	userRepo := userRepository.NewGormUserRepository()
	flightRepo := flightRepository.NewGormFlightRepository()
	UserService := userService.NewUserService(userRepo)
	FlightService := flightService.NewFlightService(flightRepo)
	UserController := usercontroller.UserController{UserService: UserService}
	FlightController := flightcontroller.FlightController{FlightService: FlightService}

	// Public routes
	e.GET("/flights/:id", FlightController.GetFlightByID)
	e.GET("/flights/:origin/:destination/:date", FlightController.GetFlightByDate)
	e.GET("/flights/planes", FlightController.GetPlanesList)
	e.GET("/flights/cities", FlightController.GetCitiesList)
	e.GET("/flights/days", FlightController.GetDaysList)

	// Signup and login routes
	e.POST("/signup", UserController.Signup)
	e.POST("/login", UserController.Login)
}
