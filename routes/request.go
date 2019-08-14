package routes

//Add Address like work laptop version
type MemberRequest struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Lob         string `json:"lob,omitempty"`
	Pcp         string `json:"pcp,omitempty"`
	UpdateField string `json:"update,omitempty"`
	NewValue    string `json:"new,omitempty"`
}
