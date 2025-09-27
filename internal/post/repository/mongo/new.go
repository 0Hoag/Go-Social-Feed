package mongo

import (
	"time"

	"github.com/hoag/go-social-feed/internal/post/repository"
	"github.com/hoag/go-social-feed/pkg/log"
	"github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/util"
)

type impleRepository struct {
	l     log.Logger
	db    mongo.Database
	clock func() time.Time
}

func New(
	l log.Logger,
	db mongo.Database,
) repository.Repository {
	now := util.Now
	return &impleRepository{
		l:     l,
		db:    db,
		clock: now,
	}
}
