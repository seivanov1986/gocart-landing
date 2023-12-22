package cors

type middleware struct{}

func New() *middleware {
	return &middleware{}
}
