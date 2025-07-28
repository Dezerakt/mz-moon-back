package artist

type Request struct {
	ArtistName string `json:"artistName,omitempty" validate:"required,max=50"`
}

func (obj *Request) ToEntity() *Entity {
	return &Entity{
		ArtistName: obj.ArtistName,
	}
}
