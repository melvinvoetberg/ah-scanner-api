package ahproduct

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client"
	"encoding/json"
	"strconv"
)

type Detail struct {
	Product Product `json:"productCard"`
}

type Product struct {
	FIR uint64 `json:"webshopId"`
	Title string `json:"title"`
	AvailableOnline bool `json:"availableOnline`
	Orderable bool `json:"isOrderable"`
}

func GetProduct(fir uint64) Product {
	ac := ahclient.Client()
	resBody := ac.Get("product/detail/v4/fir/" + strconv.FormatUint(fir, 10))

	var d Detail
  json.Unmarshal([]byte(resBody), &d)

	return d.Product
}
