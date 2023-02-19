package dataservice

import (
	"context"
)

type Repository interface {
	PostAdd(ctx context.Context) (string, bool, error)
	PostDelete(ctx context.Context) (string, bool, error)

	ArticleAdd(ctx context.Context) (string, bool, error)
	ArticleDelete(ctx context.Context) (string, bool, error)

	FeedList(ctx context.Context)

	TagSubscribe(ctx context.Context)
	TagUnsubscribe(ctx context.Context)
}
