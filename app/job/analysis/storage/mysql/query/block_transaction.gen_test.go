// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"web3/app/job/analysis/storage/mysql/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.BlockTransaction{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.BlockTransaction{}) fail: %s", err)
	}
}

func Test_blockTransactionQuery(t *testing.T) {
	blockTransaction := newBlockTransaction(db)
	blockTransaction = *blockTransaction.As(blockTransaction.TableName())
	_do := blockTransaction.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(blockTransaction.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <block_transaction> fail:", err)
		return
	}

	_, ok := blockTransaction.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from blockTransaction success")
	}

	err = _do.Create(&model.BlockTransaction{})
	if err != nil {
		t.Error("create item in table <block_transaction> fail:", err)
	}

	err = _do.Save(&model.BlockTransaction{})
	if err != nil {
		t.Error("create item in table <block_transaction> fail:", err)
	}

	err = _do.CreateInBatches([]*model.BlockTransaction{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <block_transaction> fail:", err)
	}

	_, err = _do.Select(blockTransaction.ALL).Take()
	if err != nil {
		t.Error("Take() on table <block_transaction> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <block_transaction> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <block_transaction> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <block_transaction> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.BlockTransaction{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <block_transaction> fail:", err)
	}

	_, err = _do.Select(blockTransaction.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <block_transaction> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <block_transaction> fail:", err)
	}

	_, err = _do.Select(blockTransaction.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <block_transaction> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <block_transaction> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <block_transaction> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <block_transaction> fail:", err)
	}

	_, err = _do.ScanByPage(&model.BlockTransaction{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <block_transaction> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <block_transaction> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <block_transaction> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <block_transaction> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <block_transaction> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <block_transaction> fail:", err)
	}
}
