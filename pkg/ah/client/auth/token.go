package ahauth

import (
	"encoding/json"
	"io/ioutil"
  "net/http"
  "bytes"
  "fmt"
  "os"
)

type Token struct {
  Access string `json:"access_token"`
  Refresh string `json:"refresh_token"`
  Anonymous bool `json:"anonymous"`
}

func GetToken() Token {
  var t Token
  t = loadToken()

  if t.Access == "" {
    t = requestNewToken()
  }

  return t
}

func saveToken(token Token) {
  f, _ := json.Marshal(token)
  _ = ioutil.WriteFile("token.json", f, 0644)
}

func loadToken() Token {
  f, _ := os.Open("token.json")
  b, _ := ioutil.ReadAll(f)

  var t Token
  json.Unmarshal([]byte(b), &t)
  return t
}

func requestNewToken() Token {
  var url string
  var data []byte
  var anonymous bool

  code, err := GetCode()
  c := &http.Client{}

  if err != nil {
    url = "https://api.ah.nl/mobile-auth/v1/auth/token/anonymous"
    data = []byte(`{"clientId":"appie"}`)
    anonymous = true
  } else {
    url = "https://api.ah.nl/mobile-auth/v1/auth/token"
    data = []byte(`{"clientId":"appie","code":"` + code + `"}`)
    anonymous = false
  }

  body := bytes.NewBuffer(data)

  req, err := http.NewRequest("POST", url, body)

  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Accept", "*/*")
  req.Header.Add("X-Application", "AHWEBSHOP")
  req.Header.Add("Accept-Language", "en-NL;q=1.0, nl-NL;q=0.9")
  req.Header.Add("User-Agent", "nl.ah.Appie/7.34 Model/iPhone iPhoneOS/15.0")

  res, err := c.Do(req)

  if err != nil {
    fmt.Println("Failure : ", err)
  }

  resBody, _ := ioutil.ReadAll(res.Body)

  var t Token
  json.Unmarshal([]byte(resBody), &t)
  t.Anonymous = anonymous

  saveToken(t)

  return t
}
