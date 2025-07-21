package vo

type Database string

const (
	App Database = "app"
)

func (obj Database) String() string {
	return string(obj)
}

type Collection string

const (
	Mood Collection = "mood"
	Note Collection = "note"
)

func (obj Collection) String() string {
	return string(obj)
}
