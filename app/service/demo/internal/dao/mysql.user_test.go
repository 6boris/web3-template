package dao

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"web3/app/service/demo/storage/mysql/model"
)

func TestDao_MySQL_User_Unite_CRUD(t *testing.T) {
	t.Run("MySQLCreateUser", func(t *testing.T) {
		createData := &model.User{
			NickName: uuid.NewString(),
		}
		err := testDao.MySQLCreateUser(testCtx, createData)
		assert.Nil(t, err)
		spew.Dump(createData)
	})
	t.Run("MySQLCreateUser", func(t *testing.T) {
		user, err := testDao.MySQLGetUser(testCtx, 1)
		assert.Nil(t, err)
		spew.Dump(user)
	})
}
