package repository

import (
	"github.com/k-ueki/tmanager/server/domain/entity"
)

type (
	FollowerQueryRepository interface {
		FindByUserID(userID int) (*entity.User, error)
		FindFollowerByUserID(userID int) ([]*entity.User, error)
		FindFollowByUserID(userID int) ([]*entity.User, error)
		FindUnrequitedUser(userID int) ([]*entity.User, error)
		FindFollowerTwitterIDsByUserID(userID int) ([]*entity.TwitterID, error)
	}

	FollowerCommandRepository interface {
		Reset(inp []interface{}) error
	}
)
