package types

type AppConfig struct {
	MinioEndpoint        string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKeyId     string `mapstructure:"MINIO_ACCESS_KEY_ID"`
	MinioSecretAccessKey string `mapstructure:"MINIO_SECRET_ACCESS_KEY"`
	DatabaseDSN          string `mapstructure:"DSN"`
	SnowFlakeMachineId   int    `mapstructure:"SNOWFLAKE_MACHINE_ID"`
	LogPath              string `mapstructure:"LOG_PATH"`
}
