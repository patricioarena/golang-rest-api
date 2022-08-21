package services

import "patricioa.e.arena/rest-api/models"

type GameService interface {
    GetAll() ([]*models.Game, error)
}
