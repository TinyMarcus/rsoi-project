package utils

type KafkaConfig struct {
	Endpoints []string `json:"endpoints"`
	Topics    []string `json:"topics"`
}

type Configuration struct {
	DB      DBConfiguration `json:"db"`
	LogFile string          `json:"log_file"`
	Port    uint16          `json:"port"`
	Kafka   KafkaConfig     `json:"kafka"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	Config = Configuration{
		DB: DBConfiguration{
			"postgres",
			"statistics",
			"postgres",
			"postgres",
			"postgres",
			"5432",
		},
		LogFile: "logs/server.log",
		Port:    8030,
		Kafka:   KafkaConfig{Endpoints: []string{"kafka:9092"}, Topics: []string{"statistics"}},
	}
}
