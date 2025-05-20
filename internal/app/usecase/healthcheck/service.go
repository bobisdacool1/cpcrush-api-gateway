package healthcheck

type usecase struct{}

func NewService() Service {
	return &usecase{}
}

func (s *usecase) Ping() map[string]string {
	return map[string]string{"status": "ok"}
}
