package healthclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func RetrieveHealthInfo(protocol string, host string, port int, path string) (string, error) {
	url := fmt.Sprintf("%s://%s:%d%s", protocol, host, port, path)
	response, err := netClient.Get(url)
	if err != nil {
		return "", err
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
