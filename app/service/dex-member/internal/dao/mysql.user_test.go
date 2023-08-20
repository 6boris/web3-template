package dao

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3/app/service/dex-finance/storage/mysql/model"
)

func TestDao_MySQL_User_Unite_CRUD(t *testing.T) {
	t.Run("BizCreateUser", func(t *testing.T) {

		createData := &model.User{
			NickName: uuid.NewString(),
		}
		err := testDao.BizCreateUser(testCtx, createData)
		assert.Nil(t, err)
		spew.Dump(createData)
		user, err := testDao.BizGetAccount(testCtx, createData.ID)
		assert.Nil(t, err)
		spew.Dump(user)
	})
}
