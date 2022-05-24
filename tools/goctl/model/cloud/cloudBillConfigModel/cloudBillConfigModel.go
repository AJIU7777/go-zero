package cloudBillConfigModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CloudBillConfigModel = (*customCloudBillConfigModel)(nil)

type (
	// CloudBillConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCloudBillConfigModel.
	CloudBillConfigModel interface {
		cloudBillConfigModel
	}

	customCloudBillConfigModel struct {
		*defaultCloudBillConfigModel
	}
)

// NewCloudBillConfigModel returns a model for the database table.
func NewCloudBillConfigModel(conn sqlx.SqlConn, c cache.CacheConf) CloudBillConfigModel {
	return &customCloudBillConfigModel{
		defaultCloudBillConfigModel: newCloudBillConfigModel(conn, c),
	}
}
