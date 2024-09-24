package models

type Policy struct {
	Id     string `json:"_id" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`
}
