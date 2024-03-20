package resolver

import (
	"context"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
)

// ID is the resolver for the id field.
func (r *entityResolver) ID(ctx context.Context, obj *structs.Entity) (string, error) {
	return obj.ID.String(), nil
}

// Entity returns graph.EntityResolver implementation.
func (r *Resolver) Entity() graph.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
