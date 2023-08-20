package dao

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDao_Biz_User_Unite(t *testing.T) {
	t.Run("BizGetCoinInfo", func(t *testing.T) {
		user, err := testDao.BizGetCoinInfo(testCtx, 1001784310992068612)
		assert.Nil(t, err)
		spew.Dump(user)
	})
}
