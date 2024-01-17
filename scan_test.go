package tronscan

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var _apikey = os.Getenv("API_KEY")
var _address = os.Getenv("TEST_ADDRESS")

func TestScan_GetAccount(t *testing.T) {
	response, err := getScan().GetAccount(context.Background(), _address)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestScan_GetAccountAuthSecurity(t *testing.T) {
	response, err := getScan().GetAccountAuthSecurity(context.Background(), _address)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#+v", response)
}

func TestScan_ListTrc20TransfersWithStatus(t *testing.T) {
	req := &ListTrc20TransfersWithStatusRequest{
		Start:     0,
		Limit:     10,
		Trc20Id:   "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		Address:   _address,
		DbVersion: 0,
	}
	response, err := getScan().ListTrc20TransfersWithStatus(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#+v", response)
}

func TestScan_ListAccountTokens(t *testing.T) {
	req := &ListAccountTokensRequest{
		Address: _address,
		Start:   0,
		Limit:   10,
		Token:   "USDT",
	}
	response, err := getScan().ListAccountTokens(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#+v", response)
}

func getScan() (scan *Scan) {
	client := NewHttpClient("https://apilist.tronscanapi.com")
	scan = NewScan(client)
	scan.SetKeys([]string{_apikey})

	return
}
