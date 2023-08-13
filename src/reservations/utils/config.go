package utils

type DatabaseConfiguration struct {
	Type string `json:"type"`
	Name string `json:"name"`

	User     string `json:"user"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Configuration struct {
	DB      DatabaseConfiguration `json:"db"`
	LogFile string                `json:"logFile"`
	Port    uint16                `json:"port"`
	RawJWKS string                `json:"rawJWKS"`
}

var Config Configuration

func InitConfiguration() {
	Config = Configuration{
		DatabaseConfiguration{
			"postgres",
			"reservations",
			"postgres",
			"postgres",
			"localhost",
			"5432",
		},
		"logs/servel.log",
		8070,
		`{"keys":[{"kty":"RSA","alg":"RS256","kid":"zavChDW9u-0gmG6lZhy4rgWspQjKQHVQ4F8hry_7ack","use":"sig","e":"AQAB","n":"kMYZfrUwocgUB-ZuzHu6qmt_Mnd4dgoOEbxLTAP-sfb2C2tBMpQ92Pa-JE7JxzDpJ485hJZh7hdObKU6cEMmLmFSCDuAfXt1dki4lhSFA8iXzRxpO6qNWkiDd48MQwLuRCC8Vog6EYGra-l3bN8j2kQR4FaK7HNDlOUWAU4qXHGhkFCEqr9rU3J74T_BcPAbGZfceyHh2a1wW84GwvGg7WYq0PmgIW5xri-VMHNNJBBIIBy9VHc84AZw0eAeNfY4G2Nf62d9Mxjs8LpSJwsYd093DealpnWapXp8ZioEiZJldEmwBtvkSI5H35upwuCABQrNFasNtno6XlmX-qw60Q"}]}`,
	}
}
