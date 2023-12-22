package user

func (s *service) Logout(sessionId string) error {
	_, err := s.sessionManager.Del(sessionId)
	return err
}
