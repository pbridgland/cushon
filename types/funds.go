package types

//Funds is a slice of Fund
type Funds []Fund

//Fund is the what users can choose to invest in
type Fund struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
