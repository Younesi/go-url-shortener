package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	_urlController "github.com/younesi/go-url-shortener/controllers"
	_urlRepository "github.com/younesi/go-url-shortener/repositories/mysql"
	"github.com/younesi/go-url-shortener/routes"
	_urlService "github.com/younesi/go-url-shortener/services"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// We have decided to use MYSQL as our database.
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tehran")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := gin.Default()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	urlRepository := _urlRepository.NewMysqlUrlRepository(dbConn)

	urlService := _urlService.NewUrlservice(urlRepository, timeoutContext)
	urlController := _urlController.NewUrlController(urlService)

	routes.LoadApiRoutes(r, urlController)

	serverPort := viper.GetString(`server.port`)
	runErr := r.Run(serverPort)
	if runErr != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", runErr))
	}

}
