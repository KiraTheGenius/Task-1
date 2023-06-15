package flightcontroller

import (
	"net/http"
	"strconv"
	"ticket/services/flightService"
	"time"

	"github.com/labstack/echo"
)

type FlightController struct {
	FlightService flightService.FlightService
}

func (f *FlightController) GetFlightByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	result, err := f.FlightService.GetFlight(int64(id))
	if err != nil {
		return c.String(http.StatusNotFound, "Flight Not Found!!")

	}
	return c.JSON(http.StatusOK, result)
}

func (f *FlightController) GetFlightByDate(c echo.Context) error {

	origin := c.Param("origin")
	destination := c.Param("destination")
	dateStr := c.Param("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid date format")
	}

	result, err := f.FlightService.GetFlightByDate(origin, destination, date)
	if err != nil {
		return c.String(http.StatusNotFound, "Flight Not Found!!")

	}

	return c.JSON(http.StatusOK, result)
}

func (f *FlightController) GetPlanesList(c echo.Context) error {
	result, err := f.FlightService.GetPlanesList()
	if err != nil {
		return c.String(http.StatusNotFound, "No Planes Found!!")

	}
	return c.JSON(http.StatusOK, result)
}

func (f *FlightController) GetCitiesList(c echo.Context) error {
	result, err := f.FlightService.GetCitiesList()
	if err != nil {
		return c.String(http.StatusNotFound, "No City Found!!")

	}
	return c.JSON(http.StatusOK, result)
}

func (f *FlightController) GetDaysList(c echo.Context) error {
	result, err := f.FlightService.GetDaysList()
	if err != nil {
		return c.String(http.StatusNotFound, "No Day Found!!")

	}
	return c.JSON(http.StatusOK, result)
}
