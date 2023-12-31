package model

const (
	//  账户类型
	ACCOUNT_TYPE_DEFAULT        = "DEFAULT"        // 默认
	ACCOUNT_TYPE_BALANCE        = "BALANCE"        // 可以使用的余额
	ACCOUNT_TYPE_NET_VALUE      = "NET_VALUE"      // 净值
	ACCOUNT_TYPE_FREEZE         = "FREEZE"         // 冻结金额
	ACCOUNT_TYPE_TOTAL_RECHARGE = "TOTAL_RECHARGE" // 累计充值
	ACCOUNT_TYPE_TOTAL_WITHDRAW = "TOTAL_WITHDRAW" // 累计提现
	ACCOUNT_TYPE_FLOAT_PROFIT   = "FLOAT_PROFIT"   // 持仓盈亏/浮动盈亏
	ACCOUNT_TYPE_TOTAL_PROFIT   = "TOTAL_PROFIT"   // 总盈亏

	//  交易类型
	TRADE_TYPE_DEFAULT    = "DEFAULT"    // 默认
	TRADE_TYPE_RECHARGE   = "RECHARGE"   // 充值
	TRADE_TYPE_WITHDRAWAL = "WITHDRAWAL" // 提现
	TRADE_TYPE_TRANS_IN   = "TRANS_IN"   // 转入
	TRADE_TYPE_TRANS_OUT  = "TRANS_OUT"  // 转出

)
