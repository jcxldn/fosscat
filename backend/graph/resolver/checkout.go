package resolver

import (
	"context"

	"github.com/jcxldn/fosscat/backend/graph"
	"github.com/jcxldn/fosscat/backend/structs"
)

// ID is the resolver for the id field.
func (r *checkoutResolver) ID(ctx context.Context, obj *structs.Checkout) (string, error) {
	return obj.ID.String(), nil
}

// Checkout returns graph.CheckoutResolver implementation.
func (r *Resolver) Checkout() graph.CheckoutResolver { return &checkoutResolver{r} }

type checkoutResolver struct{ *Resolver }
