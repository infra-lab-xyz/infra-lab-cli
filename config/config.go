package config

import (
	"fmt"
	"infra-lab-cli/src/utils"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var GlobalSettings ILCConfig
var GlobalSettingsConfig *viper.Viper
var ProjectPath = "~/.infra-lab"

func GetConfig() *ILCConfig {
	return &GlobalSettings
}

func setDefaultsAndBindEnvs(cfg any, path, envPrefix string, viperCfg *viper.Viper) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic in setDefaultsAndBindEnvs: %v\n", r)
		}
	}()

	cfgType := reflect.TypeOf(cfg)
	cfgValue := reflect.ValueOf(cfg)

	for i := 0; i < cfgType.NumField(); i++ {
		field := cfgType.Field(i)
		tag := field.Tag.Get("mapstructure")

		if field.Type.Kind() == reflect.Struct {
			var newPath string
			if path == "" {
				newPath = tag
			} else {
				newPath = path + "." + tag
			}
			setDefaultsAndBindEnvs(cfgValue.Field(i).Interface(), newPath, envPrefix, viperCfg)
		} else {
			var cfgPath string
			defaultValue := field.Tag.Get("default")
			if path == "" {
				cfgPath = tag
			} else {
				cfgPath = path + "." + tag
			}
			envPath := strings.ToUpper(strings.ReplaceAll(cfgPath, ".", "__"))
			err := viperCfg.BindEnv(cfgPath, envPrefix+"_"+envPath)
			if err != nil {
				fmt.Printf("BindEnv error: %v\n", err)
			}
			viperCfg.SetDefault(cfgPath, defaultValue)

		}
	}
}

// TODO: add test to be sure all the configs has default value or not if this is expected

func LoadConfig() (err error) {
	envPrefix := "ILC_"

	GlobalSettingsConfig.SetEnvPrefix(envPrefix)
	GlobalSettingsConfig.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	GlobalSettingsConfig.AutomaticEnv()

	setDefaultsAndBindEnvs(GlobalSettings, "", envPrefix, GlobalSettingsConfig)
	// TODO: Since IDK how to gather this value from config before config loaded and not ready to just resolve env var will keep projectDir static
	configPath := utils.ExpandPath(fmt.Sprintf("%s/infra-lab-cli.yaml", ProjectPath))

	GlobalSettingsConfig.SetConfigFile(configPath)
	_ = GlobalSettingsConfig.ReadInConfig()
	if err = GlobalSettingsConfig.Unmarshal(&GlobalSettings); err != nil {
		return err
	}

	return nil
}

func init() {
	GlobalSettingsConfig = viper.New()
	_ = LoadConfig()
}
