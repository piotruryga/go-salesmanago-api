package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"strconv"
	"time"
)

const (
	APP1_ENDPOINT       = "https://www.salesmanago.pl/api"
	HEADER_ACCEPT       = "application/json, application/json"
	HEADER_CONTENT_TYPE = "application/json;charset=UTF-8"
)

func main() {

	connStr := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	connStr = fmt.Sprintf(connStr, "localhost", "5432", "postgres", "sm_api_metrix", "disable")
	//db, err := gorm.Open("postgres", connStr)
	var err error
	dbTT.pgDB, err = gorm.Open("postgres", connStr)
	dbTT.pgDB.DB().SetMaxOpenConns(100)
	dbTT.pgDB.DB().SetMaxIdleConns(10)

	if err != nil {
		panic(err)
	}
	defer dbTT.pgDB.Close()

	var authRequest = AuthRequest{
		ApiKey:      "123abc",
		ClientId:    "h4jsu6pc5txybj04",
		RequestTime: strconv.FormatInt(time.Now().Unix(), 10),
		ApiSecret:   "4f69782e826841f794080cae87648e42",
		Owner:       "admin@vendor.pl",
	}
	authRequest.CreateSha()
	client := http.Client{Timeout: 30 * time.Second}
	InitRequestFactory()

	methodsScheduler := func() {
		go handleHasContact(authRequest, client)
		//go handleDeleteContact(authRequest, client)

	}

	scheduler.Every(5).Seconds().Run(methodsScheduler)
	time.Sleep(10 * time.Minute)

}
