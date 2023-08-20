// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"web3/app/service/dex-finance/storage/mysql/model"
)

func newCoin(db *gorm.DB, opts ...gen.DOOption) coin {
	_coin := coin{}

	_coin.coinDo.UseDB(db, opts...)
	_coin.coinDo.UseModel(&model.Coin{})

	tableName := _coin.coinDo.TableName()
	_coin.ALL = field.NewAsterisk(tableName)
	_coin.ID = field.NewInt64(tableName, "id")
	_coin.Name = field.NewString(tableName, "name")
	_coin.Title = field.NewString(tableName, "title")
	_coin.Img = field.NewString(tableName, "img")
	_coin.Type = field.NewString(tableName, "type")
	_coin.Round = field.NewInt64(tableName, "round")
	_coin.BaseAmount = field.NewFloat64(tableName, "base_amount")
	_coin.MinAmount = field.NewFloat64(tableName, "min_amount")
	_coin.MaxAmount = field.NewFloat64(tableName, "max_amount")
	_coin.DayMaxAmount = field.NewFloat64(tableName, "day_max_amount")
	_coin.Status = field.NewBool(tableName, "status")
	_coin.Rate = field.NewFloat64(tableName, "rate")
	_coin.MinFeeNum = field.NewFloat64(tableName, "min_fee_num")
	_coin.WithdrawFlag = field.NewInt32(tableName, "withdraw_flag")
	_coin.RechargeFlag = field.NewInt32(tableName, "recharge_flag")
	_coin.CreatedAt = field.NewTime(tableName, "created_at")
	_coin.UpdatedAt = field.NewTime(tableName, "updated_at")

	_coin.fillFieldMap()

	return _coin
}

type coin struct {
	coinDo coinDo

	ALL          field.Asterisk
	ID           field.Int64   // 币种ID
	Name         field.String  // 币种名称
	Title        field.String  // 币种标题
	Img          field.String  // 币种LOGO
	Type         field.String  // 币种类型：100000001:BTC 200000001:ETH 200000002:ETH ERC20
	Round        field.Int64   // 小数位数
	BaseAmount   field.Float64 // 最小提现单位
	MinAmount    field.Float64 // 单笔最小提现数量
	MaxAmount    field.Float64 // 单笔最大提现数量
	DayMaxAmount field.Float64 // 当日最大提现数量
	Status       field.Bool    // 状态: 0:启用 1:禁用
	Rate         field.Float64 // 手续费率
	MinFeeNum    field.Float64 // 最低收取手续费个数
	WithdrawFlag field.Int32   // 提现开关 0:关闭 1:打开
	RechargeFlag field.Int32   // 充值开关 0:关闭 1:打开
	CreatedAt    field.Time    // 创建时间
	UpdatedAt    field.Time    // 更新时间

	fieldMap map[string]field.Expr
}

func (c coin) Table(newTableName string) *coin {
	c.coinDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coin) As(alias string) *coin {
	c.coinDo.DO = *(c.coinDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coin) updateTableName(table string) *coin {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Name = field.NewString(table, "name")
	c.Title = field.NewString(table, "title")
	c.Img = field.NewString(table, "img")
	c.Type = field.NewString(table, "type")
	c.Round = field.NewInt64(table, "round")
	c.BaseAmount = field.NewFloat64(table, "base_amount")
	c.MinAmount = field.NewFloat64(table, "min_amount")
	c.MaxAmount = field.NewFloat64(table, "max_amount")
	c.DayMaxAmount = field.NewFloat64(table, "day_max_amount")
	c.Status = field.NewBool(table, "status")
	c.Rate = field.NewFloat64(table, "rate")
	c.MinFeeNum = field.NewFloat64(table, "min_fee_num")
	c.WithdrawFlag = field.NewInt32(table, "withdraw_flag")
	c.RechargeFlag = field.NewInt32(table, "recharge_flag")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *coin) WithContext(ctx context.Context) *coinDo { return c.coinDo.WithContext(ctx) }

func (c coin) TableName() string { return c.coinDo.TableName() }

func (c coin) Alias() string { return c.coinDo.Alias() }

func (c coin) Columns(cols ...field.Expr) gen.Columns { return c.coinDo.Columns(cols...) }

func (c *coin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coin) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 17)
	c.fieldMap["id"] = c.ID
	c.fieldMap["name"] = c.Name
	c.fieldMap["title"] = c.Title
	c.fieldMap["img"] = c.Img
	c.fieldMap["type"] = c.Type
	c.fieldMap["round"] = c.Round
	c.fieldMap["base_amount"] = c.BaseAmount
	c.fieldMap["min_amount"] = c.MinAmount
	c.fieldMap["max_amount"] = c.MaxAmount
	c.fieldMap["day_max_amount"] = c.DayMaxAmount
	c.fieldMap["status"] = c.Status
	c.fieldMap["rate"] = c.Rate
	c.fieldMap["min_fee_num"] = c.MinFeeNum
	c.fieldMap["withdraw_flag"] = c.WithdrawFlag
	c.fieldMap["recharge_flag"] = c.RechargeFlag
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c coin) clone(db *gorm.DB) coin {
	c.coinDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coin) replaceDB(db *gorm.DB) coin {
	c.coinDo.ReplaceDB(db)
	return c
}

type coinDo struct{ gen.DO }

func (c coinDo) Debug() *coinDo {
	return c.withDO(c.DO.Debug())
}

func (c coinDo) WithContext(ctx context.Context) *coinDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c coinDo) ReadDB() *coinDo {
	return c.Clauses(dbresolver.Read)
}

func (c coinDo) WriteDB() *coinDo {
	return c.Clauses(dbresolver.Write)
}

func (c coinDo) Session(config *gorm.Session) *coinDo {
	return c.withDO(c.DO.Session(config))
}

func (c coinDo) Clauses(conds ...clause.Expression) *coinDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c coinDo) Returning(value interface{}, columns ...string) *coinDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c coinDo) Not(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c coinDo) Or(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c coinDo) Select(conds ...field.Expr) *coinDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c coinDo) Where(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c coinDo) Order(conds ...field.Expr) *coinDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c coinDo) Distinct(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c coinDo) Omit(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c coinDo) Join(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c coinDo) LeftJoin(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c coinDo) RightJoin(table schema.Tabler, on ...field.Expr) *coinDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c coinDo) Group(cols ...field.Expr) *coinDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c coinDo) Having(conds ...gen.Condition) *coinDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c coinDo) Limit(limit int) *coinDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c coinDo) Offset(offset int) *coinDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c coinDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *coinDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c coinDo) Unscoped() *coinDo {
	return c.withDO(c.DO.Unscoped())
}

func (c coinDo) Create(values ...*model.Coin) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c coinDo) CreateInBatches(values []*model.Coin, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c coinDo) Save(values ...*model.Coin) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c coinDo) First() (*model.Coin, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Take() (*model.Coin, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Last() (*model.Coin, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) Find() ([]*model.Coin, error) {
	result, err := c.DO.Find()
	return result.([]*model.Coin), err
}

func (c coinDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Coin, err error) {
	buf := make([]*model.Coin, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c coinDo) FindInBatches(result *[]*model.Coin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c coinDo) Attrs(attrs ...field.AssignExpr) *coinDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c coinDo) Assign(attrs ...field.AssignExpr) *coinDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c coinDo) Joins(fields ...field.RelationField) *coinDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c coinDo) Preload(fields ...field.RelationField) *coinDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c coinDo) FirstOrInit() (*model.Coin, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) FirstOrCreate() (*model.Coin, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coin), nil
	}
}

func (c coinDo) FindByPage(offset int, limit int) (result []*model.Coin, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c coinDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c coinDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c coinDo) Delete(models ...*model.Coin) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *coinDo) withDO(do gen.Dao) *coinDo {
	c.DO = *do.(*gen.DO)
	return c
}