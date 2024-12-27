package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"ayush": {
		AuthToken: "123ABC",
		Username: "ayushsharma",
	},
	"varsha": {
		AuthToken: "123ABC",
		Username: "varshaayushkibiwi",
	},
	"farhan": {
		AuthToken: "453ABC",
		Username: "farhankh1234",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"ayush" : {
		Coins: 100,
		Username: "ayushsharma",
	},
	"varsha" : {
		Coins: 500,
		Username: "varshaayushkibiwi",
	},
	"farhan" :{
		Coins: 0,
		Username: "farhankh1234",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil 
	}
	return &clientData
}