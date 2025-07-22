package progressionDmn

type Progression struct {
	Id       int      `json:"id" bson:"id"`
	Mood     string   `json:"mood" bson:"mood"`
	RootNote string   `json:"rootNote" bson:"rootNote"`
	Notes    []string `json:"notes" bson:"notes"`
}
