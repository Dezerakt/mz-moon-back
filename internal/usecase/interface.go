package usecase

import (
	"bytes"
	"context"
	songDmn "spotifykiller/internal/domain/song"
)

type (
	Song interface {
		Upload(ctx context.Context, fileMeta *songDmn.Request, buffer *bytes.Buffer)
	}
)
