package user

import (
	"github.com/doffy007/banking-api/libs/constant"
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/libs/utils"
	"github.com/doffy007/banking-api/pkg/model"
)

func (r *UserRepository) LogoutUser(query *types.Query, email string) error {
	model := new(model.User)

	sql := r.db.Model(model)

	if query != nil {
		if len(query.SelectColumns) > 0 {
			sql = sql.Select(utils.UniqueSlice(query.SelectColumns, constant.DEFAULT_SELECT_COLUMNS))
		}
	}

	if email != "" {
		sql = sql.Model(model).Where("status = ?", true).Update("status = ?", false)
	}

	if err := sql.First(model).Error; err != nil {
		return err
	}

	return nil
}
