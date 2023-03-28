package acctstatement

import (
	"batch-acctstatement/config"
	"batch-acctstatement/helpers"
	"bytes"
	"crypto/sha256"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	BaseURL      = os.Getenv("OB_URL")
	ApiKey       = os.Getenv("OB_API_KEY")
	ApiSecret    = os.Getenv("OB_API_SECRET")
	SecondaryKey = os.Getenv("OB_SECONDARY_KEY")
	ApiToken     = os.Getenv("OB_API_TOKEN")
	PartnerID    = os.Getenv("OB_PARTNER_ID")
)

const (
	GET = iota
	PUT
	POST
	DELETE
)

type Request struct {
	UserReferenceNumber string        `json:"UserReferenceNumber"`
	RequestTime         string        `json:"RequestTime"`
	AccountNumber       string        `json:"AccountNumber"`
	StartDate           string        `json:"StartDate"`
	EndDate             string        `json:"EndDate"`
	PagingControl       PagingControl `json:"PagingControl"`
}

type PagingControl struct {
	MaximumRecord string `json:"MaximumRecord"`
	NextRecord    string `json:"NextRecord"`
	MatchedRecord string `json:"MatchedRecord"`
}

type Response struct {
	ResponseTime      string         `json:"ResponseTime"`
	CodeStatus        string         `json:"CodeStatus"`
	DescriptionStatus string         `json:"DescriptionStatus"`
	AvailableBalance  string         `json:"AvailableBalance"`
	Transactions      []Transactions `json:"Transactions"`
	PagingControl     PagingControl  `json:"PagingControl"`
}
type Transactions struct {
	PostingDate       string `json:"PostingDate"`
	EffectiveDate     string `json:"EffectiveDate"`
	TransactionAmount string `json:"TransactionAmount"`
	DebitCredit       string `json:"DebitCredit"`
	ReferenceNumber   string `json:"ReferenceNumber"`
	Description       string `json:"Description"`
	TransactionType   string `json:"TransactionType"`
	BranchCode        string `json:"BranchCode"`
}

func AccountStatement(body *Request, token string) (*Response, error) {
	log.Println("Begin request account statement..")
	defer func() {
		log.Println("End request account statement..")
	}()
	var path = "/v2/api/financialinfo/casa/accountstatement"
	var timeNow = helpers.GetLocalDate("UTC")
	var bdiTimestamp = timeNow.Format("2006-01-02T15:04:05.999Z")

	body.RequestTime = helpers.GetLocalDate(config.TIMEZONE).Format("20060102150405")
	var genNumber = helpers.RandCharacter(34)
	body.UserReferenceNumber = PartnerID + genNumber

	var signature = GenerateSignature(POST, path, bdiTimestamp, body)

	var bodyBytes, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 30 * time.Second,
	}

	log.Printf("Sending request to %s%s", BaseURL, path)
	log.Println("With Body: ", string(bodyBytes))
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", BaseURL, path), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	var refereneceNumber = helpers.RandCharacter(16)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("BDI-Key", ApiKey)
	req.Header.Add("BDI-Timestamp", bdiTimestamp)
	req.Header.Add("BDI-Signature", signature)
	req.Header.Add("ReferenceNumber", refereneceNumber)
	log.Println("Authorization:", "Bearer "+token)
	log.Println("BDI-Key:", ApiKey)
	log.Println("BDI-Timestamp:", bdiTimestamp)
	log.Println("BDI-Signature:", signature)
	log.Println("ReferenceNumber", refereneceNumber)
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
	if resp.StatusCode != 200 {
		log.Printf("Status not ok %d", resp.StatusCode)
		log.Printf("Response: %+v", string(respBytes))
		return nil, errors.New("status not ok")
	}

	log.Println("Response:", string(respBytes))
	if err := json.Unmarshal(respBytes, &response); err != nil {
		log.Printf("Error unmarshaling response, cause: %+v\n", err)
		return nil, err
	}

	return &response, nil
}

func GenerateSignature(requestType int, path, timestamp string, body interface{}) string {
	var result string
	switch requestType {
	case GET, DELETE:
		var payload = path + timestamp + ApiKey + ApiSecret
		var sha = sha256.New()
		sha.Write([]byte(payload))
		result = fmt.Sprintf("%x", sha.Sum(nil))
	case PUT, POST:
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			log.Printf("Error marshaling body, cause: %+v\n", err)
			return result
		}
		var bodyString = strings.ReplaceAll(string(bodyBytes), "\r", "")
		bodyString = strings.ReplaceAll(bodyString, "\n", "")
		bodyString = strings.ReplaceAll(bodyString, "\t", "")
		bodyString = strings.ReplaceAll(bodyString, " ", "")
		var payload = path + timestamp + ApiKey + ApiSecret + bodyString
		log.Println("payload:", payload)
		var sha = sha256.New()
		sha.Write([]byte(payload))
		result = fmt.Sprintf("%x", sha.Sum(nil))
	}

	return result
}
