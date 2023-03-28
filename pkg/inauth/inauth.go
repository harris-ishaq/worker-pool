package inauth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	BaseURL  = os.Getenv("OB_URL")
	ApiToken = os.Getenv("OB_API_TOKEN")
)

type Request struct {
	GrantType string `json:"grant_type"`
}

type Response struct {
	AccessToken  string  `json:"access_token,omitempty"`
	TokenType    string  `json:"token_type,omitempty"`
	ExpiresIn    float64 `json:"expires_in,omitempty"`
	Scope        string  `json:"scope,omitempty"`
	ErrorMessage struct {
		Indonesian string `json:"Indonesian"`
		English    string `json:"English"`
	} `json:"errMessage,omitempty"`
	ErrorCode        string `json:"ErrorCode,omitempty"`
	ErrorDescription string `json:"ErrorDescription,omitempty"`
}

func GetToken(body *Request) (*Response, error) {
	log.Println("Begin request get token..")
	defer func() {
		log.Println("End request get token..")
	}()
	var formBody = url.Values{}
	formBody.Add("grant_type", body.GrantType)
	log.Printf("Body: %s", formBody.Encode())

	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
	}

	log.Println("URL:", fmt.Sprintf("%s%s", BaseURL, "/api/oauth/token"))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", BaseURL, "/api/oauth/token"), strings.NewReader(formBody.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", ApiToken)
	log.Println("Authorization: ", ApiToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body, cause: %+v\n", err)
		return nil, err
	}
	log.Println("Response:", string(respBytes))
	if err := json.Unmarshal(respBytes, &response); err != nil {
		log.Printf("Error unmarshaling response, cause: %+v\n", err)
		return nil, err
	}

	return &response, nil
}
