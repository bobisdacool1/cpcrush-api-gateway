package healthcheck

type Service interface {
	Ping() map[string]string
}
