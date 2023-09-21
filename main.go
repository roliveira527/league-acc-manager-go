package main

import (
	"encoding/json"
	"io/ioutil"
)

type AccountData struct {
	SummonerName, AccountID, AccountPW string
}

func main() {
	data := AccountData{
		SummonerName: "yuumiplayer1",
		AccountID:    "ihateshaco",
		AccountPW:    "hehehecat123",
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("acc_data.json", file, 0644)
}
