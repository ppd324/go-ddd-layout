package conf

import (
	"flag"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type AppConfig struct {
	Server *ServerConfig `yaml:"server"`
	Mysql  *MysqlConfig  `yaml:"mysql"`
}

var ProviderSet = wire.NewSet(
	NewAppConfig,
	wire.FieldsOf(new(*AppConfig), "Mysql"),
)

type ServerConfig struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

type MysqlConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	TableName string `yaml:"table_name"`
	User      string `yaml:"user"`
	Passwd    string `yaml:"passwd"`
	PoolSize  int    `yaml:"pool_size"`
	Charset   string `yaml:"charset"`
}

func GetEnv() string {
	// 1. 从命令行参数读取（例如 --env=prod）
	var env string
	flag.StringVar(&env, "env", "", "运行环境（例如 dev, prod）")
	flag.Parse()

	// 2. 如果命令行未指定，从环境变量读取
	if env == "" {
		env = os.Getenv("APP_ENV")
	}

	// 3. 默认值（例如 dev）
	if env == "" {
		env = "dev"
	}

	return strings.ToLower(env)
}

func GetConfigFile() string {
	env := GetEnv()
	var configFile string
	flag.StringVar(&configFile, "config", "", "yaml config file path")
	flag.Parse()
	if configFile == "" {
		configFile = os.Getenv("APP_CONFIG_FILE")
	}
	if configFile == "" {
		configFile = "../configs/config-" + env + ".yaml"
	}
	return configFile
}

func NewAppConfig() (*AppConfig, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(GetConfigFile())
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed, err:%v\n", err)
	}
	var conf AppConfig
	err = v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf, nil
}
