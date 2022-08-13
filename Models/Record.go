package Models

type Record struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Func string `json:"func"`
	Time string `json:"time"`
}
