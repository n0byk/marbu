package repository

import "context"

type Repository interface {
	ServiceRepository
	BlogRepository
	ProfileRepository
	FeedRepository
}

type ServiceRepository interface {
	CheckConnection()
	CloseConnection()
}

type BlogRepository interface {
	BlogList(ctx context.Context, id string) error
	BlogCreate(ctx context.Context, id string) error
	BlogGet(ctx context.Context, id string) error
	BlogUpdate(ctx context.Context, id string) error
	BlogDelete(ctx context.Context, id string) error
}

type ProfileRepository interface {
	ProfileList(ctx context.Context, id string) error
	ProfileCreate(ctx context.Context, id string) error
	ProfileGet(ctx context.Context, id string) error
	ProfileUpdate(ctx context.Context, id string) error
	ProfileDelete(ctx context.Context, id string) error
}

type FeedRepository interface {
	FeedList(ctx context.Context, id string) (string, error)
}
