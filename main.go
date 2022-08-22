package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"patricioa.e.arena/rest-api/controllers"
	docs "patricioa.e.arena/rest-api/docs"
	"patricioa.e.arena/rest-api/services"
)

var (
	server         *gin.Engine
	scrapingService    services.ScrapingService
	scrapingController controllers.ScrapingController
	ctx            context.Context
	scrapedData    *mongo.Database
	gameCollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

// @contact.name   LinkedIn
// @contact.url    https://www.linkedin.com/in/patricio-ernesto-antonio-arena-08a0a9133/
// @contact.email  patricio.e.arena@gmail.com
func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	mode := os.Getenv("ENVIRONMENT")
	hostname := os.Getenv("HOSTNAME")

	if err != nil {
		log.Fatal("error trying to get hostname", err)
	}

	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))
	mongoclient, err = mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	scrapedData = mongoclient.Database("ScrapedData")
	collectionNames, err := scrapedData.ListCollectionNames(ctx, bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	gameCollection = scrapedData.Collection(collectionNames[0])
	scrapingService = services.NewScrapingService(gameCollection, ctx)
	scrapingController = controllers.NewScrapingController(scrapingService)

	defer mongoclient.Disconnect(ctx)

	gin.SetMode(strings.ToLower(mode))
	server = gin.New()

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Golang REST API"
	docs.SwaggerInfo.Description = "My first rest api with mongodb and Golang."
	docs.SwaggerInfo.Version = "1.0"

    docs.SwaggerInfo.Host = strings.ToLower(hostname) + ":" + port
    if strings.ToLower(mode) != "debug" {
        docs.SwaggerInfo.Host = strings.ToLower(hostname)
    }

	docs.SwaggerInfo.BasePath = "/api/v1"

	docs.SwaggerInfo.Schemes = []string{"http"}
    if strings.ToLower(mode) != "debug" {
        docs.SwaggerInfo.Schemes = []string{"https"}
    }

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		games := apiRoutes.Group("game")
		{
			games.GET("/all", scrapingController.GetAll)
			games.GET("/all/order-by-discount", scrapingController.GetAllPerDiscount)
			games.GET("/page", scrapingController.GetPage)

		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

    swaggerURL := "http://" + strings.ToLower(hostname) + ":" + port + "/swagger/index.html"
    if strings.ToLower(mode) != "debug" {
        swaggerURL = "https://" + strings.ToLower(hostname) + "/swagger/index.html"
    }

    fmt.Println(swaggerURL)
	log.Fatal(server.Run(":" + port))

}
