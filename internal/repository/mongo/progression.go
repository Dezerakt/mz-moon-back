package mongoRepo

import (
	progressionDmn "accord-generator/internal/domain/progression"
	"accord-generator/internal/repository"
	mongoPkg "accord-generator/pkg/mongo"
	"accord-generator/vo"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type progression struct {
	ProgressionCollection *mongo.Collection
}

func NewProgressionRepo(client *mongoPkg.Wrap) repository.Progression {
	return &progression{
		ProgressionCollection: client.AppCollection(vo.Progression),
	}
}

func (obj *progression) GetAll(ctx context.Context) ([]progressionDmn.Progression, error) {
	var result []progressionDmn.Progression

	find, err := obj.ProgressionCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "moodRepo.GetAll: can't execute find mongo")
	}

	for find.Next(ctx) {
		var tempRecord progressionDmn.Progression

		if err = find.Decode(&tempRecord); err != nil {
			return nil, errors.Wrap(err, "moodRepo.GetAll: error while decode data")
		}

		result = append(result, tempRecord)
	}

	return result, nil
}

func (obj *progression) GetByParams(ctx context.Context, params map[string]interface{}) ([]progressionDmn.Progression, error) {
	filter := bson.M{}

	for key, value := range params {
		switch value.(type) {
		case string:
			if value.(string) == "" {
				continue
			}
		case int:
			if value.(int) == 0 {
				continue
			}
		}

		filter[key] = value
	}

	find, err := obj.ProgressionCollection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "moodRepo.GetAll: can't execute find mongo")
	}

	var result []progressionDmn.Progression
	for find.Next(ctx) {
		var tempRecord progressionDmn.Progression

		if err = find.Decode(&tempRecord); err != nil {
			return nil, errors.Wrap(err, "moodRepo.GetAll: error while decode data")
		}

		result = append(result, tempRecord)
	}

	return result, nil
}
