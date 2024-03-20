package resolver

import (
	"context"
	b64 "encoding/base64"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/graph/model"
	"github.com/jcxldn/fosscat/backend/structs"
)

// ID is the resolver for the id field.
func (r *fileResolver) ID(ctx context.Context, obj *structs.File) (string, error) {
	return obj.ID.String(), nil
}

// Data is the resolver for the data field.
func (r *fileResolver) Data(ctx context.Context, obj *structs.File) (string, error) {
	return b64.StdEncoding.EncodeToString(obj.Data), nil
}

// Type is the resolver for the type field.
func (r *fileResolver) Type(ctx context.Context, obj *structs.File) (model.FileTypes, error) {
	return model.FileTypes(obj.Type), nil
}

// File returns graph.FileResolver implementation.
func (r *Resolver) File() graph.FileResolver { return &fileResolver{r} }

type fileResolver struct{ *Resolver }
