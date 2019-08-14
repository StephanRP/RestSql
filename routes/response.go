package routes

//Add Address like work laptop version
type Member struct {
	Id      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Lob     string  `json:"lob,omitempty"`
	Pcp     string  `json:"pcp,omitempty"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	Id          int    `json:"id,omitempty"`
	AddressLine string `json:"addressLine,omitempty"`
	City        string `json:"lob,omitempty"`
	State       string `json:"pcp,omitempty"`
	Zip         int    `json:"zip,omitempty"`
}
