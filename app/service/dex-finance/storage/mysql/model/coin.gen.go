// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCoin = "coin"

// Coin mapped from table <coin>
type Coin struct {
	ID           int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:币种ID" json:"id"`                                     // 币种ID
	Name         string     `gorm:"column:name;type:varchar(100);not null;uniqueIndex:idx_uq_name,priority:1;comment:币种名称" json:"name"`             // 币种名称
	Title        string     `gorm:"column:title;type:varchar(100);not null;comment:币种标题" json:"title"`                                              // 币种标题
	Img          string     `gorm:"column:img;type:varchar(255);not null;comment:币种LOGO" json:"img"`                                                // 币种LOGO
	Type         string     `gorm:"column:type;type:varchar(50);not null;comment:币种类型：100000001:BTC 200000001:ETH 200000002:ETH ERC20" json:"type"` // 币种类型：100000001:BTC 200000001:ETH 200000002:ETH ERC20
	Round        int64      `gorm:"column:round;type:bigint;not null;comment:小数位数" json:"round"`                                                    // 小数位数
	BaseAmount   *float64   `gorm:"column:base_amount;type:decimal(20,8);default:0.00000000;comment:最小提现单位" json:"base_amount"`                     // 最小提现单位
	MinAmount    *float64   `gorm:"column:min_amount;type:decimal(20,8);default:0.00000000;comment:单笔最小提现数量" json:"min_amount"`                     // 单笔最小提现数量
	MaxAmount    *float64   `gorm:"column:max_amount;type:decimal(20,8);default:0.00000000;comment:单笔最大提现数量" json:"max_amount"`                     // 单笔最大提现数量
	DayMaxAmount float64    `gorm:"column:day_max_amount;type:decimal(20,8);comment:当日最大提现数量" json:"day_max_amount"`                                // 当日最大提现数量
	Status       *bool      `gorm:"column:status;type:tinyint(1);not null;default:1;comment:状态: 0:启用 1:禁用" json:"status"`                           // 状态: 0:启用 1:禁用
	Rate         float64    `gorm:"column:rate;type:double;comment:手续费率" json:"rate"`                                                               // 手续费率
	MinFeeNum    float64    `gorm:"column:min_fee_num;type:decimal(20,8);comment:最低收取手续费个数" json:"min_fee_num"`                                     // 最低收取手续费个数
	WithdrawFlag int32      `gorm:"column:withdraw_flag;type:tinyint;comment:提现开关 0:关闭 1:打开" json:"withdraw_flag"`                                  // 提现开关 0:关闭 1:打开
	RechargeFlag int32      `gorm:"column:recharge_flag;type:tinyint;comment:充值开关 0:关闭 1:打开" json:"recharge_flag"`                                  // 充值开关 0:关闭 1:打开
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`              // 创建时间
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`              // 更新时间
}

// TableName Coin's table name
func (*Coin) TableName() string {
	return TableNameCoin
}
