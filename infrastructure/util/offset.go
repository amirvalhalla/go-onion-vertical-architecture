package util

func CalcOffset(pageIndex int, pageSize int) int {
	return (pageIndex - 1) * pageSize
}
