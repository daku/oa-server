package libapi

import (
	"github.com/openaccounting/oa-server/core/model"
	"github.com/openaccounting/oa-server/core/model/types"
)

func CreateUser(user *types.User) error {
	err := model.Instance.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}
