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
	err := db.AutoMigrate(&model.AccountTrade{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.AccountTrade{}) fail: %s", err)
	}
}

func Test_accountTradeQuery(t *testing.T) {
	accountTrade := newAccountTrade(db)
	accountTrade = *accountTrade.As(accountTrade.TableName())
	_do := accountTrade.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(accountTrade.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <account_trade> fail:", err)
		return
	}

	_, ok := accountTrade.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from accountTrade success")
	}

	err = _do.Create(&model.AccountTrade{})
	if err != nil {
		t.Error("create item in table <account_trade> fail:", err)
	}

	err = _do.Save(&model.AccountTrade{})
	if err != nil {
		t.Error("create item in table <account_trade> fail:", err)
	}

	err = _do.CreateInBatches([]*model.AccountTrade{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <account_trade> fail:", err)
	}

	_, err = _do.Select(accountTrade.ALL).Take()
	if err != nil {
		t.Error("Take() on table <account_trade> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <account_trade> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <account_trade> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <account_trade> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.AccountTrade{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <account_trade> fail:", err)
	}

	_, err = _do.Select(accountTrade.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <account_trade> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <account_trade> fail:", err)
	}

	_, err = _do.Select(accountTrade.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <account_trade> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <account_trade> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <account_trade> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <account_trade> fail:", err)
	}

	_, err = _do.ScanByPage(&model.AccountTrade{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <account_trade> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <account_trade> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <account_trade> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <account_trade> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <account_trade> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <account_trade> fail:", err)
	}
}