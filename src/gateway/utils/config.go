package utils

type KafkaConfig struct {
	Endpoints []string `json:"endpoints"`
	Topics    []string `json:"topics"`
}

type EndpointConfig struct {
	Statistics       string `json:"statistics"`
	IdentityProvider string `json:"identity-provider"`
	Reservations     string `json:"reservations"`
	Payments         string `json:"payments"`
	Loyalties        string `json:"loyalties"`
}

type Configuration struct {
	LogFile   string         `json:"log_file"`
	Port      uint16         `json:"port"`
	RawJWKS   string         `json:"raw-jwks"`
	Kafka     KafkaConfig    `json:"kafka"`
	Endpoints EndpointConfig `json:"endpoints"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		"logs/server.log",
		8080,
		`{"keys":[{"kty":"RSA","alg":"RS256","kid":"XVh9VRM57Bic_gSk2s4owqSKYVQYhZrd7ONvInnszyQ","use":"sig","e":"AQAB","n":"r25i2X_caK8RpM5r4Gugi0N01TGL-rR_3f7vNgkXpL0RlvgJSTWjt8o_NqreLE2b9YLktYI8o7e1Asmz-U2wGA0cepppU5g-7J7B83KyWc8a71Wzj5fSBHr3_SQx2L_sPSQ2lp27fdVjIeL-c2htV69889MGz3ut3snJiGbMNdxfyEbcL8OzjUp1BYkM69A-NBc8xwCwZNWkll5sIHxIb7D5S4m_SnyZ3VTdTTvbhz-8ao3j8ofjWfuC5ys4sTLVXVrWUPKv6yrLBqcLeMezHDaUF-Ocx62dpJTC3-ZhtctmFWhHJdsK0U0VAkHaDX6qGsnpigL0ukskt04rhkr0lQ"}]}`,
		KafkaConfig{Endpoints: []string{"kafka:9092"}, Topics: []string{"statistics"}},
		EndpointConfig{
			Statistics:       "http://statistics:8030",
			IdentityProvider: "http://identity-provider:8040",
			Loyalties:        "http://loyalties:8050",
			Payments:         "http://payments:8060",
			Reservations:     "http://reservations:8070",
		},
	}
}
