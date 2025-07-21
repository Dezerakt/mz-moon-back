package mongoRepo

import (
	moodDmn "accord-generator/internal/domain/mood"
	"accord-generator/internal/repository"
	mongoPkg "accord-generator/pkg/mongo"
	"accord-generator/vo"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type moodRepo struct {
	MoodCollection *mongo.Collection
}

func NewMoodRepo(client *mongoPkg.Wrap) repository.Mood {
	return &moodRepo{
		MoodCollection: client.AppCollection(vo.Mood),
	}
}

func (obj *moodRepo) GetAll(ctx context.Context) ([]moodDmn.Mood, error) {
	var result []moodDmn.Mood

	find, err := obj.MoodCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "moodRepo.GetAll: can't execute find mongo")
	}

	for find.Next(ctx) {
		var tempRecord moodDmn.Mood

		if err = find.Decode(&tempRecord); err != nil {
			return nil, errors.Wrap(err, "moodRepo.GetAll: error while decode data")
		}

		result = append(result, tempRecord)
	}

	return result, nil
}
