package noteDmn

type Note struct {
	Id   int     `json:"id" bson:"id"`
	Name string  `json:"name" bson:"name"`
	Alt  *string `json:"alt" bson:"alt"`
}
