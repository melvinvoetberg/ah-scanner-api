package ahproduct

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client"
	"encoding/json"
	"strconv"
)

func SearchGTIN(gtin int) Product {
	ac := ahclient.Client()
	resBody := ac.Get("product/search/v1/gtin/" + strconv.Itoa(gtin))

	var p Product
  json.Unmarshal([]byte(resBody), &p)

	return p
}
