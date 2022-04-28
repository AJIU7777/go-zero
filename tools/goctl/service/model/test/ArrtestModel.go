package test

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArrtestModel = (*customArrtestModel)(nil)

type (
	// ArrtestModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArrtestModel.
	ArrtestModel interface {
		arrtestModel
		TestFind(id int64) (*Arrtest, error)
	}

	customArrtestModel struct {
		*defaultArrtestModel
	}
)

// NewArrtestModel returns a model for the database table.
func NewArrtestModel(conn sqlx.SqlConn) ArrtestModel {
	return &customArrtestModel{
		defaultArrtestModel: newArrtestModel(conn),
	}
}

func (m *defaultArrtestModel) TestFind(id int64) (*Arrtest, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", arrtestRows, m.table)
	var resp Arrtest
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
