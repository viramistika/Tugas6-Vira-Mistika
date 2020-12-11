package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cm "pnp-master/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) TripHandler(ctx context.Context, req cm.TripRequest) (res cm.TripResponse) {
	defer panicRecovery()

	//Request to server
	tripRequest := &cm.TripRequest{
		Provinsi:       req.Provinsi,
		DepartureDate1: req.DepartureDate1,
		DepartureDate2: req.DepartureDate2,
	}

	reqBody, err := json.Marshal(tripRequest)
	if err != nil {
		panic(err.Error())
	}

	resp, err := http.Post("http://35.186.147.192/travel/GetTripsSample.php", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var response cm.TripResponse
	json.Unmarshal(body, &response)

	res.Message = response.Message
	res.Status = response.Status
	res.TripDetail = response.TripDetail

	//Insert into database
	var db *sql.DB
	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)
	db, err = sql.Open("mysql", mySQL)
	if err != nil {
		panic(err.Error())
	}

	for _, data := range response.TripDetail {
		AirlineName := data.AirlineName
		CityName := data.CityName
		Duration := data.Duration
		HotelName := data.HotelName

		fmt.Println("AirlineName : ", AirlineName)
		fmt.Println("CityName : ", CityName)
		fmt.Println("Duration : ", Duration)
		fmt.Println("HotelName : ", HotelName)

		sql := "INSERT INTO trip (AirlineName, CityName, Duration, HotelName) values (?,?,?,?)"
		stmt, err := db.Prepare(sql)
		if err != nil {
			panic(err.Error())
		}

		_, err = stmt.Exec(AirlineName, CityName, Duration, HotelName)

	}

	return
}
