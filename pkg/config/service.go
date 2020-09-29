package config

type Service struct {
	UDP  UDP
	HTTP HTTP
}

type UDP struct {
	Address string
}

type HTTP struct {
	Address string
}
