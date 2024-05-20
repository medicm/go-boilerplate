package config

import (
	"fmt"
	"go.uber.org/fx"
	"os"
	"reflect"
	"strings"
	"sync"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Env            *string  `validate:"oneof=dev stg prod"`
	LogLevel       *string  `validate:"oneof=debug info warn error fatal panic"`
	Port           *int     `validate:"min=1,max=65535"`
	BasePath       string   `validate:"omitempty"`
	DbUrl          string   `validate:"required,url"`
	TrustedProxies []string `validate:"required"`
}

var (
	cfg     *Config
	onceCfg sync.Once
	errCfg  error
)

func LoadConfig() (*Config, error) {
	onceCfg.Do(func() {
		var path string
		path, errCfg = os.Getwd()
		if errCfg != nil {
			return
		}

		errCfg = godotenv.Load(fmt.Sprintf("%s/.env", path))
		if os.IsNotExist(errCfg) {
			env := os.Getenv("APP_ENV")
			if env == "" {
				env = "prod"
			}

			envFileName := fmt.Sprintf(".env.%s", env)
			errCfg = godotenv.Load(fmt.Sprintf("%s/%s", path, envFileName))
		}
		if errCfg != nil && !os.IsNotExist(errCfg) {
			return
		}

		setDefaultValues()

		var Configs Config

		configValue := reflect.ValueOf(&Configs).Elem()
		for i := 0; i < configValue.NumField(); i++ {
			fieldName := configValue.Type().Field(i).Name
			envVarName := convertToEnvVarName(fieldName)
			if errCfg = viper.BindEnv(fieldName, envVarName); errCfg != nil {
				errCfg = fmt.Errorf("error binding environment variable: %w", errCfg)
				return
			}
		}

		if errCfg = viper.Unmarshal(&Configs); errCfg != nil {
			return
		}

		trustedProxiesStr := viper.GetString("TRUSTED_PROXIES")
		Configs.TrustedProxies = parseTrustedProxies(trustedProxiesStr)

		if errCfg = ValidateConfig(Configs); errCfg != nil {
			return
		}

		cfg = &Configs
	})

	return cfg, errCfg
}

func convertToEnvVarName(str string) string {
	envVarName := make([]rune, 0, len(str)+len(str)/2)
	for i, char := range str {
		if i > 0 && unicode.IsUpper(char) && !unicode.IsUpper(rune(str[i-1])) {
			envVarName = append(envVarName, '_')
		}
		envVarName = append(envVarName, unicode.ToUpper(char))
	}

	envVarName = append([]rune("APP_"), envVarName...)

	return string(envVarName)
}

func ValidateConfig(cfg Config) error {
	validate := validator.New()
	return validate.Struct(cfg)
}

func parseTrustedProxies(proxies string) []string {
	if proxies == "" {
		return nil
	}

	return strings.Split(proxies, ",")
}

func setDefaultValues() {
	viper.SetDefault("Port", 1337)
	viper.SetDefault("BasePath", "/")

	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
}

var Module = fx.Options(
	fx.Provide(LoadConfig),
)
