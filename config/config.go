package config

import (
	"fmt"
	"time"
	"strings"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("pokescraper")
}

func readViperConfig(appName string) *viper.Viper {
	var err error
	/*home, err := os.UserHomeDir()
	if err != nil { 
		log.Fatal("Warning no home dir found")
	}
	projectFolder := fmt.Sprintf("%s/.%s/",home, appName)
	*/
	projectFolder := "/home/marvin/Projects/Privat/pokescraper/"
	configFile := "config.yml"
	v := viper.New()
	v.SetConfigName(configFile) // name of config file (without extension)
	v.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath(projectFolder)  // call multiple times to add many search paths
	err = v.ReadInConfig() // Find and read the config file
	if err != nil { 
		fmt.Println("Warning no config file found")
	}
	v.SetEnvPrefix(strings.ToUpper(appName))
	v.AutomaticEnv()
	

	// global defaults
	
	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")
	v.SetDefault("api_url", "https://pokeapi.co/api/v2/pokedex/kanto/")
	v.SetDefault("root", projectFolder)
	v.SetDefault("sqlite_db", projectFolder + ".db")
	v.SetDefault("config_folder", projectFolder)
	v.SetDefault("config_file", configFile)
	v.SetDefault("config_path", projectFolder+configFile)
	v.SetDefault("asset_folder", projectFolder+"assets/")
	v.SetDefault("configured", false)
	return v
}
