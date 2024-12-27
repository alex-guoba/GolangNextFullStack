package setting

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	Conf = &Configuration{}
)

// type EmailSettingS struct {
// 	Host     string
// 	Port     int
// 	UserName string
// 	Password string
// 	IsSSL    bool
// 	From     string
// 	To       []string
// }

// type JWTSettingS struct {
// 	Secret string
// 	Issuer string
// 	Expire time.Duration
// }

type ServerSettingS struct {
	RunMode      string        `mapstructure:"RunMode"`
	HTTPPort     string        `mapstructure:"HTTPPort"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

type AppSettingS struct {
	DefaultPageSize       int           `mapstructure:"DefaultPageSize"`
	MaxPageSize           int           `mapstructure:"MaxPageSize"`
	ServerShutdownTimeout time.Duration `mapstructure:"ServerShutdownTimeout"`

	// DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"`
	// UploadSavePath        string        `mapstructure:"UploadSavePath"`
	// UploadServerURL       string        `mapstructure:"UploadServerURL"`
	// UploadImageMaxSize    int           `mapstructure:"UploadImageMaxSize"`
}

type LogSettingS struct {
	LogSavePath string `mapstructure:"LogSavePath"`
	LogFileName string `mapstructure:"LogFileName"`
	MaxSize     int    `mapstructure:"MaxSize"`
	MaxBackups  int    `mapstructure:"MaxBackups"`
	Compress    bool   `mapstructure:"Compress"`
	Level       string `mapstructure:"Level"`
}

type DatabaseSettingS struct {
	DBType       string `mapstructure:"DBType"`
	UserName     string `mapstructure:"UserName"`
	Password     string `mapstructure:"Password"`
	Host         string `mapstructure:"Host"`
	DBName       string `mapstructure:"DBName"`
	Charset      string `mapstructure:"Charset"`
	ParseTime    bool   `mapstructure:"ParseTime"`
	MaxIdleConns int    `mapstructure:"MaxIdleConns"`
	MaxOpenConns int    `mapstructure:"MaxOpenConns"`
	MigrationURL string `mapstructure:"MigrationURL"`
}

type RatelimitSettingS struct {
	Enable          bool    `mapstructure:"Enable"`
	ConfigFile      string  `mapstructure:"ConfigFile"`
	CPULoadThresh   float64 `mapstructure:"CPULoadThresh"`
	CPULoadStrategy int     `mapstructure:"CPULoadStrategy"`
}

// UnmarshalKey / Sub only uses read config, neglects environment variables
// https://github.com/spf13/viper/issues/1012
type Configuration struct {
	Server    ServerSettingS    `mapstructure:"Server"`
	App       AppSettingS       `mapstructure:"App"`
	Log       LogSettingS       `mapstructure:"Log"`
	Database  DatabaseSettingS  `mapstructure:"Database"`
	Ratelimit RatelimitSettingS `mapstructure:"Ratelimit"`
}

func Load(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("configs/")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return err
	}

	fmt.Println("path:", Conf.Log.LogSavePath)

	return nil
}
