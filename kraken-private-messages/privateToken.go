package kraken_private_messages

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// getSha256 creates a sha256 hash for given []byte
func getSha256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// getHMacSha512 creates a hmac hash with sha512
func getHMacSha512(message, secret []byte) []byte {
	mac := hmac.New(sha512.New, secret)
	mac.Write(message)
	return mac.Sum(nil)
}

func createSignature(urlPath string, values url.Values, secret []byte) string {
	// See https://www.kraken.com/help/api#general-usage for more information
	shaSum := getSha256([]byte(values.Get("nonce") + values.Encode()))
	macSum := getHMacSha512(append([]byte(urlPath), shaSum...), secret)
	return base64.StdEncoding.EncodeToString(macSum)
}

func (api *API) getPrivateToken(apiKey string, apiSecret string) (error, token) {
	values := url.Values{}
	var nonce Nonce
	values.Set("nonce", nonce.New())

	req, err := http.NewRequest("POST", REST_URI+REST_URI_PATH, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("API-Key", apiKey)
	secret, _ := base64.StdEncoding.DecodeString(apiSecret)
	req.Header.Add("API-Sign", createSignature(REST_URI_PATH, values, secret))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, ""
	}
	var jsonData KrakenTokenResponse

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return err, ""
	}
	if len(jsonData.Error) > 0 {
		return errors.New(strings.Join(jsonData.Error, "; ")), ""
	}
	return nil, jsonData.Result.Token
}
