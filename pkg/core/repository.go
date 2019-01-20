package core

import (
	"context"
	"net/http"
)

// ID is the common type interface for database IDs
type ID interface {
	String() string
}

// Repository defines the API repository implementation should follow.
type Repository interface {
	Find(ctx context.Context, id ID, result interface{}) error
	FindAll(ctx context.Context, selector interface{}, result interface{}) error
	Delete(ctx context.Context, id ID) error
	Update(ctx context.Context, id ID, item interface{}) error
	Create(ctx context.Context, items ...interface{}) error
	Count(ctx context.Context) (int, error)
	NewContext(ctx context.Context, r *http.Request) (context.Context, error)
}
