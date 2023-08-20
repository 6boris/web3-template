// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"web3/app/service/dex-finance/storage/mysql/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Account{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Account{}) fail: %s", err)
	}
}

func Test_accountQuery(t *testing.T) {
	account := newAccount(db)
	account = *account.As(account.TableName())
	_do := account.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(account.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <account> fail:", err)
		return
	}

	_, ok := account.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from account success")
	}

	err = _do.Create(&model.Account{})
	if err != nil {
		t.Error("create item in table <account> fail:", err)
	}

	err = _do.Save(&model.Account{})
	if err != nil {
		t.Error("create item in table <account> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Account{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <account> fail:", err)
	}

	_, err = _do.Select(account.ALL).Take()
	if err != nil {
		t.Error("Take() on table <account> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <account> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <account> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <account> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Account{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <account> fail:", err)
	}

	_, err = _do.Select(account.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <account> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <account> fail:", err)
	}

	_, err = _do.Select(account.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <account> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <account> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <account> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <account> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Account{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <account> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <account> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <account> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <account> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <account> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <account> fail:", err)
	}
}
