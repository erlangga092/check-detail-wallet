package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

/**
{
	"status": 200,
	"message": "Welcome to the Cek Rekening API",
	"info": {
		"endpoint": "https://cekrek.heirro.dev/api/check",
		"parameters": {
		"accountBank": [
			"bca",
			"royal",
			"bni",
			"bri",
			"cimb",
			"permata",
			"danamon",
			"dki",
			"tabungan_pensiunan_nasional",
			"nationalnobu",
			"artos",
			"hana",
			"linkaja",
			"gopay",
			"ovo",
			"dana"
		],
		"accountNumber": "081234567890"
		},
		"descriptions": {
			"bca": "BCA",
			"royal": "Blu By BCA",
			"bni": "BNI",
			"bri": "BRI",
			"cimb": "CIMB Niaga",
			"permata": "Permata",
			"danamon": "Danamon",
			"dki": "Bank DKI",
			"tabungan_pensiunan_nasional": "BTPN/Jenius",
			"nationalnobu": "Bank NOBU",
			"artos": "Bank Jago",
			"hana": "Line Bank",
			"linkaja": "LinkAja!",
			"gopay": "GoPay",
			"ovo": "OVO",
			"dana": "DANA"
		},
		"method": "POST"
	}
}
**/

var URL = "https://cekrek.heirro.dev/api/check"

type WalletDetailResult struct {
	AccountNumber string `json:"accountNumber"`
	AccountName   string `json:"accountName"`
	AccountBank   string `json:"accountBank"`
}

type WalletResult struct {
	Data    []WalletDetailResult `json:"data"`
	Status  int                  `json:"status"`
	Message string               `json:"message"`
}

func main() {
	param := url.Values{}
	param.Set("accountBank", "gopay")
	param.Set("accountNumber", "085601927930")
	payload := bytes.NewBufferString(param.Encode())

	client := &http.Client{}
	request, err := http.NewRequest("POST", URL, payload)
	if err != nil {
		fmt.Println("<<< ERROR >>>")
		fmt.Println(err.Error())
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("<<< ERROR >>>")
		fmt.Println(err.Error())
	}

	defer response.Body.Close()
	walletResult := WalletResult{}

	fmt.Println("<<< RESULT >>>")
	err = json.NewDecoder(response.Body).Decode(&walletResult)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range walletResult.Data {
		fmt.Println(v.AccountNumber)
		fmt.Println(v.AccountName)
		fmt.Println(v.AccountBank)
	}
}
