package base

type PaginationResponse[T any] struct {
	Result       *T       `json:"result"`
	PageIndex    int      `json:"pageIndex"`
	PageSize     int      `json:"pageSize"`
	TotalRecords int64    `json:"totalRecords"`
	StatusCode   int      `json:"statusCode"`
	Errors       []string `json:"errors"`
}

func NewPaginationResponse[T any](result *T, pageIndex int, pageSize int,
	totalRecords int64, statusCode int, err []string,
) PaginationResponse[T] {
	return PaginationResponse[T]{
		Result:       result,
		PageIndex:    pageIndex,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		StatusCode:   statusCode,
		Errors:       err,
	}
}
