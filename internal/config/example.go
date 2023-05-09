package config

type ExampleConfig struct {
	Env    string `mapstructure:"example_env"`
	Domain string `mapstructure:"example_domain"`
}
