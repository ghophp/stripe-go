package stripe

import "encoding/json"

type SKUParams struct {
	Params
	ID        string
	Active    *bool
	Desc      string
	Name      string
	Attrs     map[string]string
	Price     int64
	Currency  string
	Inventory Inventory
	Product   string
}

type Inventory struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
	Value    string `json:"value"`
}

type SKU struct {
	ID        string            `json:"id"`
	Created   int64             `json:"created"`
	Updated   int64             `json:"updated"`
	Live      bool              `json:"livemode"`
	Active    bool              `json:"active"`
	Name      string            `json:"name"`
	Desc      string            `json:"description"`
	Attrs     map[string]string `json:"attributes"`
	Price     int64             `json:"price"`
	Currency  string            `json:"currency"`
	Inventory Inventory         `json:"inventory"`
	Product   Product           `json:"product"`
	Meta      map[string]string `json:"metadata"`
}

type SKUList struct {
	ListMeta
	Values []*SKU `json:"data"`
}

type SKUListParams struct {
	ListParams
	Created int64
}

func (s *SKU) UnmarshalJSON(data []byte) error {
	type sku SKU
	var sk sku
	err := json.Unmarshal(data, &sk)
	if err == nil {
		*s = SKU(sk)
	} else {
		sk.ID = string(data[1 : len(data)-1])
	}

	return nil
}
