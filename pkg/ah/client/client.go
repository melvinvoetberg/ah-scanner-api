package ahclient

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client/auth"
	"io/ioutil"
	"net/http"
	"bytes"
	"fmt"
)

const BaseUrl = "https://api.ah.nl/mobile-services/"

type AhTransport struct {
	Transport http.RoundTripper
}

type AhClient struct {
	Client *http.Client
}

func (tp *AhTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t := ahauth.GetToken()

	req.Header.Add("Authorization", "Bearer " + t.Access)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-Application", "AHWEBSHOP")
	req.Header.Add("Accept-Language", "en-NL;q=1.0, nl-NL;q=0.9")
	req.Header.Add("User-Agent", "nl.ah.Appie/7.34 Model/iPhone iPhoneOS/15.0")

	return tp.transport().RoundTrip(req)
}

func (tp *AhTransport) Client() *http.Client {
	return &http.Client{Transport: tp}
}

func (tp *AhTransport) transport() http.RoundTripper {
	if tp.Transport != nil {
		return tp.Transport
	}
	return http.DefaultTransport
}

func Client() AhClient {
	var ac AhClient
	tp := AhTransport{}
	ac.Client = tp.Client()

	return ac
}

func (ac *AhClient) Get(path string) []byte {
	req, _ := http.NewRequest("GET", BaseUrl + path, nil)

	return makeRequest(ac, req)
}

func (ac *AhClient) Post(path string, json string) []byte {
	body := bytes.NewBuffer([]byte(json))
	req, _ := http.NewRequest("POST", BaseUrl + path, body)

	return makeRequest(ac, req)
}

func (ac *AhClient) Patch(path string, json string) []byte {
	body := bytes.NewBuffer([]byte(json))
	req, _ := http.NewRequest("PATCH", BaseUrl + path, body)

	return makeRequest(ac, req)
}

func makeRequest(ac *AhClient, req *http.Request) []byte {
	c := ac.Client
	res, err := c.Do(req)

	if err != nil {
		fmt.Println("Failure: ", err)
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	return resBody
}
