package dao

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3/app/service/dex-finance/storage/mysql/model"
)

func TestDao_Biz_User_Unite(t *testing.T) {
	t.Run("BizGetAccount", func(t *testing.T) {
		user, err := testDao.BizGetAccount(testCtx, 1)
		assert.Nil(t, err)
		spew.Dump(user)
	})
	t.Run("BizCreateUser", func(t *testing.T) {
		data := &model.User{
			NickName: uuid.NewString(),
		}
		err := testDao.BizCreateUser(testCtx, data)
		assert.Nil(t, err)
	})
}
