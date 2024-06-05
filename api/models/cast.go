package models

type Cast struct {
	Id        int32  `json:"id" example:"123"`
	Character string `json:"character" example:"Captain Jack Sparrow"`
	Name      string `json:"name" example:"Johnny Depp"`
	Gender    int    `json:"gender" example:"1"`
	Order     int    `json:"order" example:"9"`
}

type CastDTO struct {
	Id        int32  `json:"id" example:"123"`
	Character string `json:"character" example:"Captain Jack Sparrow"`
	Name      string `json:"name" example:"Johnny Depp"`
	Gender    string `json:"gender" example:"male"`
}

func (c *Cast) ToDTO() CastDTO {
	var gender string
	if c.Gender == 1 {
		gender = "Female"
	} else {
		gender = "Male"
	}

	return CastDTO{
		Id:        c.Id,
		Character: c.Character,
		Name:      c.Name,
		Gender:    gender,
	}
}

func CastArrayToDto(cast *[]Cast) *[]CastDTO {
	dtos := make([]CastDTO, len(*cast))
	for i := range *cast {
		dtos[i] = (*cast)[i].ToDTO()
	}
	return &dtos
}
