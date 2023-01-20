package user

import (
	"github.com/doffy007/banking-api/libs/constant"
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/libs/utils"
	"github.com/doffy007/banking-api/pkg/model"
)

func (r *UserRepository) LoginUser(query *types.Query, payload *model.User) (*model.User, error) {
	model := new(model.User)

	sql := r.db.Model(model)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(utils.UniqueSlice(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if payload != nil {
		if payload.Username != "" {
			sql = sql.Where("username = ?", payload.Username)
		}

		if payload.Password != "" {
			sql = sql.Where("password = ?", payload.Password)
		}
	}

	if err := sql.First(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
