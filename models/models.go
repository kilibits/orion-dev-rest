package models

//A Profile of member of parliament with basic details
type Profile struct {
	Name                string
	Group               string
	Area                string
	Phone               string
	Email               string
	Image               string
	MemberType          string `db:"member_type"`
	Address             string
	BirthDate           string `db:"birth_date"`
	Id                  int
	Term                int
	Source              string
	PoliticalExperience []PoliticalCareerHistory
	EducationHistory    []EducationHistory
	EmploymentHistory   []EmploymentHistory
}

//EducationHistory (Education History) of member of parliament
type EducationHistory struct {
	Institution string
	Level       string
	Award       string
	From        int
	To          int
}

//EmploymentHistory (Employment History) of member of parliament
type EmploymentHistory struct {
	Institution string
	Position    string
	From        int
	To          int
}

//PoliticalCareerHistory (Political Experience History) of member of parliament
type PoliticalCareerHistory struct {
	Institution string
	Position    string
	From        int
	To          int
}
