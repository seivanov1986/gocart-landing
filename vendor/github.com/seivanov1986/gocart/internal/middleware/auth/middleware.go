package auth

type middleware struct {
	sessionClient SessionManager
}

func New(sessionClient SessionManager) *middleware {
	return &middleware{sessionClient: sessionClient}
}
