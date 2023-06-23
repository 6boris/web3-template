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

	"web3/app/service/data/storage/mysql/model"
)

func newChainInfo(db *gorm.DB, opts ...gen.DOOption) chainInfo {
	_chainInfo := chainInfo{}

	_chainInfo.chainInfoDo.UseDB(db, opts...)
	_chainInfo.chainInfoDo.UseModel(&model.ChainInfo{})

	tableName := _chainInfo.chainInfoDo.TableName()
	_chainInfo.ALL = field.NewAsterisk(tableName)
	_chainInfo.UniqueID = field.NewInt64(tableName, "unique_id")
	_chainInfo.OrderSeq = field.NewInt64(tableName, "order_seq")
	_chainInfo.ChainID = field.NewInt64(tableName, "chain_id")
	_chainInfo.ChainName = field.NewString(tableName, "chain_name")
	_chainInfo.ChainEnv = field.NewString(tableName, "chain_env")
	_chainInfo.LogoURL = field.NewString(tableName, "logo_url")
	_chainInfo.WebsiteURL = field.NewString(tableName, "website_url")
	_chainInfo.ExplorerURL = field.NewString(tableName, "explorer_url")
	_chainInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_chainInfo.UpdatedAt = field.NewTime(tableName, "updated_at")

	_chainInfo.fillFieldMap()

	return _chainInfo
}

type chainInfo struct {
	chainInfoDo chainInfoDo

	ALL         field.Asterisk
	UniqueID    field.Int64
	OrderSeq    field.Int64
	ChainID     field.Int64
	ChainName   field.String
	ChainEnv    field.String
	LogoURL     field.String
	WebsiteURL  field.String
	ExplorerURL field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time

	fieldMap map[string]field.Expr
}

func (c chainInfo) Table(newTableName string) *chainInfo {
	c.chainInfoDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c chainInfo) As(alias string) *chainInfo {
	c.chainInfoDo.DO = *(c.chainInfoDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *chainInfo) updateTableName(table string) *chainInfo {
	c.ALL = field.NewAsterisk(table)
	c.UniqueID = field.NewInt64(table, "unique_id")
	c.OrderSeq = field.NewInt64(table, "order_seq")
	c.ChainID = field.NewInt64(table, "chain_id")
	c.ChainName = field.NewString(table, "chain_name")
	c.ChainEnv = field.NewString(table, "chain_env")
	c.LogoURL = field.NewString(table, "logo_url")
	c.WebsiteURL = field.NewString(table, "website_url")
	c.ExplorerURL = field.NewString(table, "explorer_url")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *chainInfo) WithContext(ctx context.Context) *chainInfoDo {
	return c.chainInfoDo.WithContext(ctx)
}

func (c chainInfo) TableName() string { return c.chainInfoDo.TableName() }

func (c chainInfo) Alias() string { return c.chainInfoDo.Alias() }

func (c *chainInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *chainInfo) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 10)
	c.fieldMap["unique_id"] = c.UniqueID
	c.fieldMap["order_seq"] = c.OrderSeq
	c.fieldMap["chain_id"] = c.ChainID
	c.fieldMap["chain_name"] = c.ChainName
	c.fieldMap["chain_env"] = c.ChainEnv
	c.fieldMap["logo_url"] = c.LogoURL
	c.fieldMap["website_url"] = c.WebsiteURL
	c.fieldMap["explorer_url"] = c.ExplorerURL
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c chainInfo) clone(db *gorm.DB) chainInfo {
	c.chainInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c chainInfo) replaceDB(db *gorm.DB) chainInfo {
	c.chainInfoDo.ReplaceDB(db)
	return c
}

type chainInfoDo struct{ gen.DO }

func (c chainInfoDo) Debug() *chainInfoDo {
	return c.withDO(c.DO.Debug())
}

func (c chainInfoDo) WithContext(ctx context.Context) *chainInfoDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c chainInfoDo) ReadDB() *chainInfoDo {
	return c.Clauses(dbresolver.Read)
}

func (c chainInfoDo) WriteDB() *chainInfoDo {
	return c.Clauses(dbresolver.Write)
}

func (c chainInfoDo) Session(config *gorm.Session) *chainInfoDo {
	return c.withDO(c.DO.Session(config))
}

func (c chainInfoDo) Clauses(conds ...clause.Expression) *chainInfoDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c chainInfoDo) Returning(value interface{}, columns ...string) *chainInfoDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c chainInfoDo) Not(conds ...gen.Condition) *chainInfoDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c chainInfoDo) Or(conds ...gen.Condition) *chainInfoDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c chainInfoDo) Select(conds ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c chainInfoDo) Where(conds ...gen.Condition) *chainInfoDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c chainInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *chainInfoDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c chainInfoDo) Order(conds ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c chainInfoDo) Distinct(cols ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c chainInfoDo) Omit(cols ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c chainInfoDo) Join(table schema.Tabler, on ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c chainInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c chainInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c chainInfoDo) Group(cols ...field.Expr) *chainInfoDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c chainInfoDo) Having(conds ...gen.Condition) *chainInfoDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c chainInfoDo) Limit(limit int) *chainInfoDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c chainInfoDo) Offset(offset int) *chainInfoDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c chainInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *chainInfoDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c chainInfoDo) Unscoped() *chainInfoDo {
	return c.withDO(c.DO.Unscoped())
}

func (c chainInfoDo) Create(values ...*model.ChainInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c chainInfoDo) CreateInBatches(values []*model.ChainInfo, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c chainInfoDo) Save(values ...*model.ChainInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c chainInfoDo) First() (*model.ChainInfo, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainInfo), nil
	}
}

func (c chainInfoDo) Take() (*model.ChainInfo, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainInfo), nil
	}
}

func (c chainInfoDo) Last() (*model.ChainInfo, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainInfo), nil
	}
}

func (c chainInfoDo) Find() ([]*model.ChainInfo, error) {
	result, err := c.DO.Find()
	return result.([]*model.ChainInfo), err
}

func (c chainInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ChainInfo, err error) {
	buf := make([]*model.ChainInfo, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c chainInfoDo) FindInBatches(result *[]*model.ChainInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c chainInfoDo) Attrs(attrs ...field.AssignExpr) *chainInfoDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c chainInfoDo) Assign(attrs ...field.AssignExpr) *chainInfoDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c chainInfoDo) Joins(fields ...field.RelationField) *chainInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c chainInfoDo) Preload(fields ...field.RelationField) *chainInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c chainInfoDo) FirstOrInit() (*model.ChainInfo, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainInfo), nil
	}
}

func (c chainInfoDo) FirstOrCreate() (*model.ChainInfo, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ChainInfo), nil
	}
}

func (c chainInfoDo) FindByPage(offset int, limit int) (result []*model.ChainInfo, count int64, err error) {
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

func (c chainInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c chainInfoDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c chainInfoDo) Delete(models ...*model.ChainInfo) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *chainInfoDo) withDO(do gen.Dao) *chainInfoDo {
	c.DO = *do.(*gen.DO)
	return c
}
