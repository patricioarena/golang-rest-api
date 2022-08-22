package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"patricioa.e.arena/rest-api/models"
)

type ScrapingServiceImpl struct {
	gameCollection *mongo.Collection
	ctx            context.Context
}

func NewScrapingService(gameCollection *mongo.Collection, ctx context.Context) ScrapingService {
	return &ScrapingServiceImpl{
		gameCollection: gameCollection,
		ctx:            ctx,
	}
}

func (g *ScrapingServiceImpl) GetAll() ([]*models.Game, error) {
	var games []*models.Game
	filter := bson.D{}
	cursor, err := g.gameCollection.Find(g.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(g.ctx) {
		var game models.Game
		err = cursor.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(g.ctx)

	if len(games) == 0 {
		return nil, errors.New("documents not found")
	}

	return games, nil
}

func (g *ScrapingServiceImpl) GetAllPerDiscount() ([]*models.Game, error) {
	var games []*models.Game
	filter := bson.D{}
	options := options.Find().SetSort(bson.D{{Key: "discount", Value: -1}})
	cursor, err := g.gameCollection.Find(g.ctx, filter, options)

	if err != nil {
		return nil, err
	}

	for cursor.Next(g.ctx) {
		var game models.Game
		err = cursor.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(g.ctx)

	if len(games) == 0 {
		return nil, errors.New("documents not found")
	}

	return games, nil
}

func (g *ScrapingServiceImpl) GetPage(number int, limit int) ([]*models.Game, error) {
    skip := (number - 1) * limit
    var games []*models.Game
	filter := bson.D{}
	options := options.Find().SetLimit(int64(limit)).SetSkip(int64(skip))

	cursor, err := g.gameCollection.Find(g.ctx, filter, options)

	if err != nil {
		return nil, err
	}

	for cursor.Next(g.ctx) {
		var game models.Game
		err = cursor.Decode(&game)
		if err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(g.ctx)

	if len(games) == 0 {
		return nil, errors.New("documents not found")
	}

	return games, nil
}
