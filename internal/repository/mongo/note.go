package mongoRepo

import (
	noteDmn "accord-generator/internal/domain/note"
	"accord-generator/internal/repository"
	mongoPkg "accord-generator/pkg/mongo"
	"accord-generator/vo"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type noteRepo struct {
	NoteCollection *mongo.Collection
}

func NewNoteRepo(client *mongoPkg.Wrap) repository.Note {
	return &noteRepo{
		NoteCollection: client.AppCollection(vo.Note),
	}
}

func (obj *noteRepo) GetAll(ctx context.Context) ([]noteDmn.Note, error) {
	var result []noteDmn.Note

	find, err := obj.NoteCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "moodRepo.GetAll: can't execute find mongo")
	}

	for find.Next(ctx) {
		var tempRecord noteDmn.Note

		if err = find.Decode(&tempRecord); err != nil {
			return nil, errors.Wrap(err, "moodRepo.GetAll: error while decode data")
		}

		result = append(result, tempRecord)
	}

	return result, nil
}
