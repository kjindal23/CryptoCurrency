package cryptoerver

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Service provides some "date capabilities" to your application
type Service interface {
	Status(ctx context.Context) (string, error)
	GetCurrency(ctx context.Context, symbol string) (string, error)
}

type cryptoService struct{}

// NewService makes a new Service.
func NewService() Service {
	return cryptoService{}
}

// Status only tell us that our service is ok!
func (cryptoService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

// Get will return today's date
func (cryptoService) GetCurrency(ctx context.Context, symbol string) (string, error) {
	var url string
	if symbol == "all" {
		url = "https://api.demo.hitbtc.com/api/2/public/ticker"
	} else if isSymbolValid(symbol) {
		url = fmt.Sprintf("https://api.hitbtc.com/api/2/public/ticker/" + symbol)
	} else {
		return string("Invalid Symbol "), nil
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	responsedata, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	return string(responsedata), nil
}

func isSymbolValid(symbol string) bool {
	validcurrency := []string{"BTCUSD", "ETHBTC"}
	for _, valid := range validcurrency {
		if valid == symbol {
			return true
		}
	}
	return false
}
