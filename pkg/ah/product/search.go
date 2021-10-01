package ahproduct

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client"
	"encoding/json"
	"strconv"
)

func SearchGTIN(gtin uint64) Product {
	ac := ahclient.Client()
	resBody := ac.Get("product/search/v1/gtin/" + strconv.FormatUint(gtin, 10))

	var p Product
  json.Unmarshal([]byte(resBody), &p)

	return p
}
