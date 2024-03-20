package resolver

import (
	"context"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
)

// ID is the resolver for the id field.
func (r *itemResolver) ID(ctx context.Context, obj *structs.Item) (string, error) {
	return obj.ID.String(), nil
}

// Item returns graph.ItemResolver implementation.
func (r *Resolver) Item() graph.ItemResolver { return &itemResolver{r} }

type itemResolver struct{ *Resolver }
