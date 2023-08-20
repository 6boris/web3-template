// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"web3/app/service/dex-member/storage/mysql/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.User{}) fail: %s", err)
	}
}

func Test_userQuery(t *testing.T) {
	user := newUser(db)
	user = *user.As(user.TableName())
	_do := user.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(user.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <user> fail:", err)
		return
	}

	_, ok := user.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from user success")
	}

	err = _do.Create(&model.User{})
	if err != nil {
		t.Error("create item in table <user> fail:", err)
	}

	err = _do.Save(&model.User{})
	if err != nil {
		t.Error("create item in table <user> fail:", err)
	}

	err = _do.CreateInBatches([]*model.User{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <user> fail:", err)
	}

	_, err = _do.Select(user.ALL).Take()
	if err != nil {
		t.Error("Take() on table <user> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <user> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <user> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <user> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.User{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <user> fail:", err)
	}

	_, err = _do.Select(user.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <user> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <user> fail:", err)
	}

	_, err = _do.Select(user.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <user> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <user> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <user> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <user> fail:", err)
	}

	_, err = _do.ScanByPage(&model.User{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <user> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <user> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <user> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <user> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <user> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <user> fail:", err)
	}
}