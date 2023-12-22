package file

const (
	serviceBasePathKey       = "service_base_path"
	adminPrefix              = "/admin/static"
	adminPostfix             = "/service/schemes/admin/index.html"
	serviceBasePathPostfix   = "/service/schemes/admin"
	dynamicPrefix            = "/tmp/project/images/"
	stripPrefix              = "/static"
	fileServerHandlerPostfix = "/service/schemes/public"
)

type handle struct {
}

func New() *handle {
	return &handle{}
}
