package repository

import "context"

type Repository interface {
	BlogRepository
	ProfileRepository
	FeedRepository
}

type BlogRepository interface {
	List(ctx context.Context, id string) error
	Create(ctx context.Context, id string) error
	Get(ctx context.Context, id string) error
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type ProfileRepository interface {
	List(ctx context.Context, id string) error
	Create(ctx context.Context, id string) error
	Get(ctx context.Context, id string) error
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type FeedRepository interface {
	List(ctx context.Context, id string) error
}
