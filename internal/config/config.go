package config

type HTTPServer struct {
	Address string
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage-path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() {

}
