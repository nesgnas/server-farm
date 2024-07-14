package dataStruct

type Inventory struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Slot  int    `json:"slot" bson:"slot"`
	Empty int    `json:"empty" bson:"empty"`
	Item  []Item `json:"item" bson:"item"`
}

type Item struct {
	Id       string `json:"id" bson:"_id,omitempty"`
	Name     string `json:"name" bson:"name"`
	Rank     int    `json:"rank" bson:"rank"`
	Quantity int    `json:"quantity" bson:"quantity"`
}
