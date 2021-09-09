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
	FIR int `json:"webshopId"`
	Title string `json:"title"`
	AvailableOnline bool `json:"availableOnline`
	Orderable bool `json:"isOrderable"`
}

func GetProduct(fir int) Product {
	ac := ahclient.Client()
	resBody := ac.Get("product/detail/v4/fir/" + strconv.Itoa(fir))

	var d Detail
  json.Unmarshal([]byte(resBody), &d)

	return d.Product
}
