package common

type middleware struct {
	ServiceBasePath string
}

func New(ServiceBasePath string) *middleware {
	return &middleware{ServiceBasePath: ServiceBasePath}
}
