package types

type PostPerson struct {
	Name     string   `json:"name" validate:"lte=100,required"`
	Nickname string   `json:"nickname" validate:"lte=32,required"`
	Birthday string   `json:"birthday" time_format:"2006-01-02" validate:"lte=10,required"`
	Stack    []string `json:"stack"`
}

type GetPerson struct {
	ID string `uri:"id" binding:"required,uuid"`
}
