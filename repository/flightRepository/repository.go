package flightRepository

import (
	"errors"
	"fmt"
	"ticket/models"
	"time"

	"gorm.io/gorm"
)

type FlightRepository interface {
	GetFlightsByCityAndDate(origin string, destination string, day time.Time) ([]*models.Flight, error)
	GetFlightByID(ID int64) (*models.Flight, error)
	GetPlanesList() ([]string, error)
	GetCitiesList() ([]string, error)
	GetDaysList() ([]string, error)
}

type flightGormRepository struct {
	db *gorm.DB
}

func NewGormFlightRepository() FlightRepository {
	return &flightGormRepository{
		db: getDbConnection(),
	}
}

func (fl *flightGormRepository) GetFlightsByCityAndDate(origin string, destination string, day time.Time) ([]*models.Flight, error) {
	var flights []*models.Flight
	result := fl.db.Where("origin = ? and destination = ? and date(day) = ?", origin, destination, day).Find(&flights)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no flights not found")
		}
		return nil, err
	}
	return flights, nil
}

func (fl *flightGormRepository) GetFlightByID(ID int64) (*models.Flight, error) {
	var flight models.Flight
	result := fl.db.First(&flight, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("flight not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &flight, nil
}

func (fl *flightGormRepository) GetPlanesList() ([]string, error) {
	var planes []string
	result := fl.db.Model(&models.Flight{}).Distinct("aircraft").Pluck("aircraft", &planes) // pluck Retrieves only aircraft column
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("planes not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return planes, nil
}

func (fl *flightGormRepository) GetCitiesList() ([]string, error) {
	var cities []string
	result := fl.db.Model(&models.Flight{}).Distinct("origin", "destination").Pluck("origin", &cities)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("no cities found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	result = fl.db.Model(&models.Flight{}).Distinct("origin", "destination").Pluck("destination", &cities)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return cities, nil
}

func (fl *flightGormRepository) GetDaysList() ([]string, error) {
	var days []string
	result := fl.db.Model(&models.Flight{}).Distinct("date(startTime)").Pluck("date(startTime)", &days)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("flights not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return days, nil

}
