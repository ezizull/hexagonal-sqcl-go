package postgres

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type infoDatabasPostgres struct {
	Read struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
	Write struct {
		Hostname   string
		Name       string
		Username   string
		Password   string
		Port       string
		Parameter  string
		Timezone   string
		DriverConn string
	}
}

// Database cradential
var (
	hostname = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	username = os.Getenv("POSTGRES_USERNAME")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_POSGRES")
)

// getPostgresConn is a function to get a setup of sql database connection
func (infoDB *infoDatabasPostgres) getPostgresConn(nameMap string) (err error) {
	fmt.Println("check ", username, password, hostname, port, dbname)

	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = mapstructure.Decode(viper.GetStringMap(nameMap), infoDB)
	if err != nil {
		return
	}

	if hostname != "" {
		infoDB.Read.Hostname = hostname
		infoDB.Write.Hostname = hostname
	}

	if port != "" {
		infoDB.Read.Port = port
		infoDB.Write.Port = port
	}
	if username != "" {
		infoDB.Read.Username = username
		infoDB.Write.Username = username
	}
	if password != "" {
		infoDB.Read.Password = password
		infoDB.Write.Password = password
	}

	if dbname != "" {
		infoDB.Read.Name = dbname
		infoDB.Write.Name = dbname
	}

	infoDB.Read.DriverConn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		infoDB.Read.Hostname, infoDB.Read.Port, infoDB.Read.Username, infoDB.Read.Password, infoDB.Read.Name)
	infoDB.Write.DriverConn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		infoDB.Write.Hostname, infoDB.Write.Port, infoDB.Write.Username, infoDB.Write.Password, infoDB.Write.Name)
	return
}

// initPostgresDB is a function that returns a sql database connection
func initPostgresDB(infoPg infoDatabasPostgres) (*sql.DB, error) {
	db, err := sql.Open("postgres", infoPg.Write.DriverConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
