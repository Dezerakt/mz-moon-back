package noteDmn

type Note struct {
	Name string  `json:"name" bson:"name"`
	Alt  *string `json:"alt" bson:"alt"`
}
