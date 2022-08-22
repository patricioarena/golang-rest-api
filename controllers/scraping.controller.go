package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"patricioa.e.arena/rest-api/services"
)

type ScrapingController struct {
	ScrapingService services.ScrapingService
}

func NewScrapingController(scrapingService services.ScrapingService) ScrapingController {
	return ScrapingController{
		ScrapingService: scrapingService,
	}
}

// @Summary      List games
// @Description  Get a list of all games
// @Tags         Game
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Game
// @Router /game/all [get]
func (gc *ScrapingController) GetAll(ctx *gin.Context) {
	games, err := gc.ScrapingService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, games)
}

// @Summary      List games order by discount
// @Description  Get a list of all games order by discount
// @Tags         Game
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Game
// @Router /game/all/order-by-discount [get]
func (gc *ScrapingController) GetAllPerDiscount(ctx *gin.Context) {
	games, err := gc.ScrapingService.GetAllPerDiscount()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, games)
}

// @Summary      List games with pagination
// @Description  Get a list of all games with pagination
// @Tags         Game
// @Accept       json
// @Produce      json
// @Param        number    query      int     true  "Page number"
// @Param        limit    query      int     true  "Limit results"
// @Success      200 {array} models.Game
// @Router /game/page/ [get]
func (gc *ScrapingController) GetPage(ctx *gin.Context) {
    _number := ctx.Query("number")
    _limit := ctx.Query("limit")

    limit, err := strconv.Atoi(_limit)
	number, err := strconv.Atoi(_number)
	if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
    }

    games, err := gc.ScrapingService.GetPage(number, limit)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, games)
}

// func (gc *ScrapingController) RegisterGameRoutes(rg *gin.RouterGroup) {
// 	gamerouter := rg.Group("/game")
// 	gamerouter.GET("/all", gc.GetAll)
// 	gamerouter.GET("/all/order-by-discount", gc.GetAllPerDiscount)
// 	gamerouter.GET("/page/:number", gc.GetPage)
// }
