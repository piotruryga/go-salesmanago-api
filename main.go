package main

import (
	"errors"
	"fmt"
	"github.com/carlescere/scheduler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
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
		handleHasContact(authRequest, client)

	}

	scheduler.Every(15).Seconds().Run(methodsScheduler)
	time.Sleep(10 * time.Minute)
	//handleHasContact(authRequest, client)

	deleteContactTt, error := callDeleteContact(authRequest, client)
	if error != nil {
		log.Printf("Cannot call deleteContact %s", error)
	}
	log.Println(deleteContactTt)

}

func handleHasContact(authRequest AuthRequest, client http.Client) {
	hasContactTt, error := callHasContact(authRequest, client)
	if error != nil {
		log.Printf("Cannot call hasContact %s", error)
	} else {
		dbTT.pgDB.Create(&hasContactTt)
	}
	log.Println(hasContactTt)
}

func callDeleteContact(authRequest AuthRequest, client http.Client) (TimeTrack, error) {
	if request, ok := ReturnImplementation("contactDeleteRequest").(*ContactDeleteRequest); ok {
		request.InitContactDeleteRequest("piotrek.uryga@gmail.com")
		return request.CallMethod(authRequest, client), nil
	} else {
		return TimeTrack{}, errors.New("cannot call hasContactRequest")
	}
}

func callHasContact(authRequest AuthRequest, client http.Client) (TimeTrack, error) {
	if request, ok := ReturnImplementation("hasContactRequest").(*HasContactRequest); ok {
		request.InitHasContactRequest("piotrek.uryga@gmail.com")
		return request.CallMethod(authRequest, client), nil
	} else {
		return TimeTrack{}, errors.New("cannot call hasContactRequest")
	}
}
