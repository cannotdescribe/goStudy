package myError

type IndexOverFlowError struct{}

func (e IndexOverFlowError) Error() string {
	return "数组发生的越界"
}
