package ahshoppinglist

import (
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/product"
	"github.com/melvinvoetberg/ah-scanner-api/pkg/ah/client"
	"encoding/json"
	"strconv"
)

var productKey string = `json:"product"`

type Shoppinglist struct {
	Items []Item `json:"items"`
}

type Item struct {
	Quantity int `json:"quantity"`
	Product ahproduct.Product `json:"product"`
}

func GetShoppinglist() Shoppinglist {
	ac := ahclient.Client()
	resBody := ac.Get("shoppinglist/v2/shoppinglist")

	var s Shoppinglist
  json.Unmarshal([]byte(resBody), &s)

	return s
}

func AddProduct(fir uint64) Shoppinglist {
	cs := GetShoppinglist()
	i, found := indexInShoppinglist(cs, fir)

	var q int = 1

	if found {
		q = cs.Items[i].Quantity + 1
	}

	b := `{"items":[{"quantity":` + strconv.Itoa(q) + `,"productId":` + strconv.FormatUint(fir, 10) + `,"originCode":"PRD","type":"SHOPPABLE"}]}`

	ac := ahclient.Client()
	ac.Patch("shoppinglist/v2/items", b)

	return GetShoppinglist()
}

func AddProductGTIN(gtin uint64) Shoppinglist {
	p := ahproduct.SearchGTIN(gtin)
	return AddProduct(p.FIR)
}

func indexInShoppinglist(s Shoppinglist, fir uint64) (i int, found bool) {
	for i := range s.Items {
    if s.Items[i].Product.FIR == fir {
    	return i, true
	  }
	}

	return 0, false
}
