package common

const (
	baseStoragePath = "/tmp/project"
)

type service struct {
}

func New() *service {
	return &service{}
}
