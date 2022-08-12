package Models

type Vehicle struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Model  string `json:"model"`
	Mark   string `json:"mark"`
	Color  string `json:"color"`
	Series string `json:"series"`
}
