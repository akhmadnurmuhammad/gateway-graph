package config

import (
	"os"
	"strconv"
)

type Config struct {
	Service  Service  `yaml:"service"`
	Hosts    Hosts    `yaml:"hosts"`
	Services Services `yaml:"service"`
}

type Service struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Address   string `yaml:"address"`
	Port      int    `yaml:"port"`
	JwtSecret string `yaml:"jwt_secret"`
}

type Services struct {
	AuthURL       string `yaml:"auth_url"`
	UserURL       string `yaml:"user_url"`
	MasterDataURL string `yaml:"master_data_url"`
	RolesURL      string `yaml:"roles_url"`
	BucketURL     string `yaml:"bucket_url"`
}

type Hosts struct {
	Cache     Cache     `yaml:"cache"`
	Broker    Broker    `yaml:"broker"`
	Discovery Discovery `yaml:"discovery"`
	Sentry    Sentry    `yaml:"discovery"`
}

type Cache struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	Driver  string `yaml:"driver"`
}

type Broker struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Driver   string `yaml:"driver"`
}

func (b Broker) GetDriver() string {
	return b.Driver
}

func (b Broker) GetAddress() string {
	return b.Address
}

type Discovery struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	Driver   string `yaml:"driver"`
}

type Sentry struct {
	Address string `yaml:"address"`
	Env     string `yaml:"env"`
	Driver  string `yaml:"driver"`
}

func (c *Config) Init() {
	svcPort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	brokerPort, _ := strconv.Atoi(os.Getenv("BROKER_PORT"))
	cachePort, _ := strconv.Atoi(os.Getenv("CACHE_PORT"))
	discoveryPort, _ := strconv.Atoi(os.Getenv("DISCOVERY_PORT"))

	c.Service = Service{
		Name:    os.Getenv("SERVICE_NAME"),
		Version: os.Getenv("SERVICE_VERSION"),
		Address: os.Getenv("SERVICE_ADDRESS"),
		Port:    svcPort,
	}
	c.Hosts.Broker = Broker{
		Username: os.Getenv("BROKER_USER"),
		Password: os.Getenv("BROKER_PASSWORD"),
		Address:  os.Getenv("BROKER_ADDRESS"),
		Port:     brokerPort,
		Driver:   os.Getenv("BROKER_DRIVER"),
	}

	c.Hosts.Cache = Cache{
		Port:    cachePort,
		Driver:  os.Getenv("CACHE_DRIVER"),
		Address: os.Getenv("CACHE_ADDRESS"),
	}

	c.Hosts.Sentry = Sentry{
		Address: os.Getenv("SENTRY_ADDRESS"),
		Env:     os.Getenv("SENTRY_ENV"),
		Driver:  os.Getenv("SENTRY_DRIVER"),
	}

	c.Hosts.Discovery = Discovery{
		Username: os.Getenv("DISCOVERY_USER"),
		Password: os.Getenv("DISCOVERY_PASSWORD"),
		Address:  os.Getenv("DISCOVERY_ADDRESS"),
		Port:     discoveryPort,
		Driver:   os.Getenv("DISCOVERY_DRIVER"),
	}

	c.Services = Services{
		AuthURL:       os.Getenv("AUTH_SERVICE_HOST"),
		UserURL:       os.Getenv("USER_SERVICE_HOST"),
		MasterDataURL: os.Getenv("MASTER_DATA_SERVICE_HOST"),
		RolesURL:      os.Getenv("ROLES_SERVICE_HOST"),
		BucketURL:     os.Getenv("BUCKET_SERVICE_HOST"),
	}
}
