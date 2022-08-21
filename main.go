package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "patricioa.e.arena/rest-api/controllers"
    "patricioa.e.arena/rest-api/services"
)

var (
    server         *gin.Engine
    gameService    services.GameService
    gameController controllers.GameController
    ctx            context.Context
    scrapedData    *mongo.Database
    gameCollection *mongo.Collection
    mongoclient    *mongo.Client
    err            error
)

func init() {
    godotenv.Load()
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
    gameCollection = scrapedData.Collection("2022-08-13T06:49:18.829845")
    gameService = services.NewGameService(gameCollection, ctx)
    gameController = controllers.NewGameController(gameService)
    server = gin.Default()
}

func main() {
    godotenv.Load()

    defer mongoclient.Disconnect(ctx)

    basepath := server.Group("/v1")

    gameController.RegisterGameRoutes(basepath)
    log.Fatal(server.Run(":" + os.Getenv("PORT")))

    fmt.Println("mongo connection")
}
