package config

type Service struct {
	Conn Conn
	HTTP HTTP
}

type Conn struct {
	Network string
	Address string
}

type HTTP struct {
	Address string
}
