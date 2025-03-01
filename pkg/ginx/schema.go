package ginx

import "clean-hex/pkg/framwork/errors"

type Direction string

const (
	ResBodyKey           = "res-body"
	ASC        Direction = "ASC"
	DESC       Direction = "DESC"
)

type ResponseResult struct {
	Success bool          `json:"success"`
	Data    interface{}   `json:"data,omitempty"`
	Total   int64         `json:"total,omitempty"`
	Page    int64         `json:"page,omitempty"`
	Pages   int64         `json:"pages,omitempty"`
	Error   *errors.Error `json:"error,omitempty"`
}

type PaginationResult struct {
	Total   int64         `json:"total"`
	Skip    int64         `json:"skip"`
	Limit   int64         `json:"limit"`
	OrderBy OrderByParams `json:"order_by"`
}

type OrderByParam struct {
	Field     string
	Direction Direction
}
type OrderByParams []OrderByParam

func (a OrderByParams) ToSQL() string {
	if len(a) == 0 {
		return ""
	}

	var sql string
	for _, v := range a {
		sql += v.Field + " " + string(v.Direction) + ","
	}
	return sql[:len(sql)-1]
}
