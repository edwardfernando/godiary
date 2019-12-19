package config

import "fmt"

import "time"

// ApplicationConfig manages all configurations
type ApplicationConfig struct {
	AppDBUser                    string
	AppDBHost                    string
	AppDBName                    string
	AppDBPass                    string
	AppDBMaxIdleConn             int
	AppDBMaxOpenConn             int
	AppDBConnMaxLifetimeDuration time.Duration
	AppPort                      int
}

// LoadApplicationConfig load all configuration from ENV
func LoadApplicationConfig() *ApplicationConfig {
	v := NewViper()
	v.SetDefault("LOG_LEVEL", "INFO")

	return &ApplicationConfig{
		AppDBUser:                    v.GetString("APP_DB_USER"),
		AppDBHost:                    v.GetString("APP_DB_HOST"),
		AppDBName:                    v.GetString("APP_DB_NAME"),
		AppDBPass:                    v.GetString("APP_DB_PASS"),
		AppDBMaxIdleConn:             v.GetInt("APP_DB_MAX_IDLE_CONN"),
		AppDBMaxOpenConn:             v.GetInt("APP_DB_MAX_OPEN_CONN"),
		AppDBConnMaxLifetimeDuration: v.GetDuration("APP_DB_CONN_MAX_LIFETIME_DURATION"),
		AppPort:                      v.GetInt("APP_PORT"),
	}
}

// AppDBConnectionURL returns URL to connect to database instance
func (a *ApplicationConfig) AppDBConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", a.AppDBUser, a.AppDBPass, a.AppDBHost, a.AppDBName)
}
