package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eurofurence/reg-attendee-transferclient/internal/repository/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func performGet(url string) (string, error) {
	response, err := netClient.Get(url)
	if err != nil {
		return "", err
	}
	status := response.StatusCode
	if status != http.StatusOK {
		return "", fmt.Errorf("got unexpected http status %v", status)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = response.Body.Close()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

type AttendeeMaxIdDto struct {
	MaxId uint `json:"max_id"`
}

func retrieveMaxId(url string) (uint, error) {
	body, err := performGet(url)
	if err != nil {
		log.Printf("ERROR failed to retrieve max id: %v", err)
		return 0, err
	}
	dto := &AttendeeMaxIdDto{}
	err = json.Unmarshal([]byte(body), dto)
	if err != nil {
		log.Printf("ERROR failed to parse max id response: %v", err)
		return 0, err
	}
	return dto.MaxId, nil
}

func RetrieveRegsysMaxId() (uint, error) {
	log.Printf("INFO  retrieving max id from classic regsys...")
	return retrieveMaxId(config.RegsysBaseUrl() + "/service/max-regnum-api")
}

func RetrieveAttendeeServiceMaxId() (uint, error) {
	log.Printf("INFO  retrieving max id from attendee service...")
	return retrieveMaxId(config.AttendeeServiceBaseUrl() + "/api/rest/v1/attendees/max-id")
}

// both error and success case dto
type TransferDto struct {
	Ok        bool     `json:"ok"`
	Message   string   `json:"message"`   // only if error
	Details   []string `json:"details"`   // only if error
	Timestamp string   `json:"timestamp"` // only if error
}

func PerformTransfer(id uint) error {
	log.Printf("INFO  performing transfer for id %v ...", id)
	urlWithTokenAndIdParam := fmt.Sprintf("%v/service/transfer-api?id=%v&token=%v",
		config.RegsysBaseUrl(),
		id,
		config.RegsysTransferToken())
	body, err := performGet(urlWithTokenAndIdParam)
	if err != nil {
		log.Printf("ERROR failed to perform request: %v", err)
		return err
	}
	dto := &TransferDto{}
	err = json.Unmarshal([]byte(body), dto)
	if err != nil {
		log.Printf("ERROR failed to parse response: %v", err)
		return err
	}
	if !dto.Ok {
		log.Printf("ERROR failed to initiate transfer - received not OK: %v, details: %v", dto.Message, dto.Details)
		return errors.New("failed to transfer attendee for reason " + dto.Message)
	}
	log.Printf("INFO  success")
	return nil
}
