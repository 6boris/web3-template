package metrics

import (
	"gorm.io/gorm"
)

type GormMetricsPlugin struct {
	InstName        string
	ServiceID       string
	ServiceInstance string
}

func (p *GormMetricsPlugin) Name() string {
	return ""
}

func (p *GormMetricsPlugin) Initialize(db *gorm.DB) error {
	_ = db.Callback().Query().After("gorm:query").Register("after:query", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})
	_ = db.Callback().Create().After("gorm:create").Register("after:create", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})
	_ = db.Callback().Update().After("gorm:update").Register("after:update", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})
	_ = db.Callback().Delete().After("gorm:delete").Register("after:delete", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})
	_ = db.Callback().Row().After("gorm:row").Register("after:row", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})

	_ = db.Callback().Raw().After("gorm:raw").Register("after:raw", func(db *gorm.DB) {
		cmd := ""
		if len(db.Statement.BuildClauses) > 0 {
			cmd = db.Statement.BuildClauses[0]
		}
		MySQLMetricRequests.WithLabelValues(p.InstName, p.ServiceID, p.ServiceInstance, db.Statement.Table, cmd).Inc()
	})

	return nil
}
