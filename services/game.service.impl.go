package services

import (
    "context"
    "errors"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "patricioa.e.arena/rest-api/models"
)

type GameServiceImpl struct {
    gamecollection *mongo.Collection
    ctx            context.Context
}

func NewGameService(gamecollection *mongo.Collection, ctx context.Context) GameService {
    return &GameServiceImpl{
        gamecollection: gamecollection,
        ctx:            ctx,
    }
}

func (g *GameServiceImpl) GetAll() ([]*models.Game, error) {
    var games []*models.Game
    cursor, err := g.gamecollection.Find(g.ctx, bson.D{{}})

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
