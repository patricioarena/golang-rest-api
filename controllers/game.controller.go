package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "patricioa.e.arena/rest-api/services"
)

type GameController struct {
    GameService services.GameService
}

func NewGameController(gameservice services.GameService) GameController {
    return GameController{
        GameService: gameservice,
    }
}

func (gc *GameController) GetAll(ctx *gin.Context) {
    games, err := gc.GameService.GetAll()
    if err != nil {
            ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
            return
    }
    ctx.JSON(http.StatusOK, games)
}


func (gc *GameController) RegisterGameRoutes (rg *gin.RouterGroup){
    gamerouter := rg.Group("/game")
    gamerouter.GET("/getall", gc.GetAll)
}
