package services

import (
	"patricioa.e.arena/rest-api/models"
)

type ScrapingService interface {
    GetAll() ([]*models.Game, error)
    GetAllPerDiscount() ([]*models.Game, error)
    GetPage(number int, limit int) ([]*models.Game, error)
}
