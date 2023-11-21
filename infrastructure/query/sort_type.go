package query

type SortType string

const (
	DESC SortType = "DESC"
	ASC  SortType = "ASC"
)

func (st SortType) String() string {
	switch st {
	case DESC:
		return string(DESC)
	case ASC:
		return string(ASC)
	default:
		return string(st)
	}
}
