package models

// Teacher is a struct to store teacher data
type Teacher struct {
	Name             Name    `json:"name"`
	Subject          Subject `json:"subject"`
	Capacity         int     `json:"capacity"`
	FreePeriodPerDay int     `json:"freePeriodPerWeek"`
	Classes          []Class `json:"classes"`
}
