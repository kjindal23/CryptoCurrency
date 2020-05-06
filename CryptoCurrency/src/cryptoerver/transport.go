package cryptoerver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type getRequest struct {
	symbol string
}

type resp struct {
	symbol      string `json:"symbol"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volumeQuote"`
	Timestamp   string `json:"timestamp"`
}

type getResponse struct {
	Currency string `json:"currency"`
	Err      string `json:"err,omitempty"`
}

type statusRequest struct{}

type statusResponse struct {
	Status string `json:"status"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	symbol := mux.Vars(r)["symbol"]
	req.symbol = symbol
	return req, nil
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusRequest
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
