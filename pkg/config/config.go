package config

type Configuration struct {
	Server   server
	Logger   logger
}

type server struct {
	Port string
	// path to the x509 certificate for https
	CertFile *string
	// path to the x509 private key matching `CertFile`
	KeyFile *string
}

type logger struct {
	LogLevel *string
}
