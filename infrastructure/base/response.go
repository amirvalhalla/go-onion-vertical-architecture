package base

type Response[T any] struct {
	Result     *T       `json:"result"`
	StatusCode int      `json:"statusCode"`
	Errors     []string `json:"errors"`
}

func NewResponse[T any](result *T, statusCode int, err []string) Response[T] {
	return Response[T]{
		Result:     result,
		StatusCode: statusCode,
		Errors:     err,
	}
}
