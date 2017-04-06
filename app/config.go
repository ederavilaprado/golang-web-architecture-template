package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// debug mode
	Debug bool `mapstructure:"debug"`
	// the path to the error message file. Defaults to "config/errors.yaml"
	ErrorFile string `mapstructure:"error_file"`
	// the server host bind, defaults to localhost
	Host string `mapstructure:"host"`
	// the server port, defaults to 8080
	Port int `mapstructure:"port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod string `mapstructure:"jwt_signing_method"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
}

func (config appConfig) Validate() error {
	// TODO: change this validation lib
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
	)
}

// LoadConfig loads the app configuration and populates it into the Config variable.
// The default configuration file is config/config.yaml
// Environment variables with the prefix "SERVER_" in their names are also read automatically.
func LoadConfig() error {
	v := viper.New()
	// check if the config file was specified: --config
	if ConfigFile != nil {
		v.SetConfigFile(*ConfigFile)
	} else {
		// whit this 3, wee will have the config side by side with the binary: ./config/config.yaml
		// also it's important to keep this two here to use by the tests
		v.AddConfigPath("./config")
		v.AddConfigPath("../config")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
	}
	// this will be the prefix for the env vars, following the 12factor: SERVER_PORT=8081
	v.SetEnvPrefix("server")
	v.AutomaticEnv()
	// defaults if not present inside the config file or with the env variables
	v.SetDefault("debug", "false")
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("host", "localhost")
	v.SetDefault("port", 8080)
	v.SetDefault("jwt_signing_method", "HS256")

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
