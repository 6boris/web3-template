package dao

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_Biz_User_Unite(t *testing.T) {
	t.Run("BizGetAccount", func(t *testing.T) {
		user, err := testDao.BizGetAccount(testCtx, 1)
		assert.Nil(t, err)
		spew.Dump(user)
	})
}
