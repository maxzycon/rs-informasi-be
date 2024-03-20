package dto

type DoctorRow struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	SpecializationName string `json:"specialization_name"`
}

type DoctorByIdRow struct {
	ID               uint                  `json:"id"`
	Name             string                `json:"name"`
	SpecializationID uint                  `json:"specialization_id"`
	Photo            *string               `json:"photo"`
	Skills           []*DoctorSkillsRow    `json:"skills"`
	Educations       []*DoctorEducationRow `json:"educations"`
	Slots            []*DoctorSlotRow      `json:"slots"`
}

type DoctorEducationRow struct {
	ID    uint   `json:"id"`
	Grade string `json:"grade"`
	Major string `json:"major"`
	Name  string `json:"location"`
}

type DoctorSlotRow struct {
	ID        uint   `json:"id"`
	Day       uint   `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type DoctorSkillsRow struct {
	ID   uint   `json:"id"`
	Name string `json:"description"`
}

type PayloadDoctor struct {
	Name             string                    `json:"name"`
	SpecializationID uint                      `json:"specialization_id"`
	Photo            *string                   `json:"photo"`
	Skills           []*PayloadDoctorSkills    `json:"skills"`
	Educations       []*PayloadDoctorEducation `json:"educations"`
	Slots            []*PayloadDoctorSlot      `json:"slots"`
}

type PayloadDoctorEducation struct {
	Grade string `json:"grade"`
	Major string `json:"major"`
	Name  string `json:"location"`
}

type PayloadDoctorSlot struct {
	Day       uint   `json:"day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type PayloadDoctorSkills struct {
	Name string `json:"description"`
}
