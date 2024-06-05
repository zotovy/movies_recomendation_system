package models

type Crew struct {
	Id         int
	Name       string
	Job        string
	Gender     int
	Department string
}

type CrewDTO struct {
	Id         int    `json:"id" example:"123"`
	Name       string `json:"name" example:"Johnny Depp"`
	Job        string `json:"job" example:"Camera man"`
	Gender     string `json:"gender" example:"Male"`
	Department string `json:"department" example:"Videography"`
}

func (c *Crew) ToDTO() CrewDTO {
	var gender string
	if c.Gender == 1 {
		gender = "Female"
	} else {
		gender = "Male"
	}

	return CrewDTO{
		Id:         c.Id,
		Name:       c.Name,
		Job:        c.Job,
		Gender:     gender,
		Department: c.Department,
	}
}

func CrewArrayToDto(cast *[]Crew) *[]CrewDTO {
	dtos := make([]CrewDTO, len(*cast))
	for i := range *cast {
		dtos[i] = (*cast)[i].ToDTO()
	}
	return &dtos
}
