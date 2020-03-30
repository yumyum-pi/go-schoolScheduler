package models

// NDays is the number of days in a week
const NDays = 6

// NPeriods is the number of periods in a day
const NPeriods = 8

// MaxCap is the max number of period in a week
const MaxCap = NDays * NPeriods

// Period is a single cell in the timetable
type Period struct {
	ClassID   ClassID   // assigned class's ID
	SubjectID SubjectID // assigned subject's ID
	TeacherID TeacherID // assigned teacher's ID
}

// IsAssigned return true if period is not assigned
func (p *Period) IsAssigned() bool {
	if p.ClassID == (ClassID{}) || p.SubjectID == (SubjectID{}) || p.TeacherID == (TeacherID{}) {
		return true
	}
	return false
}

// Day is a collection of 8 classes
type Day [NPeriods]Period

// Length return the number of Period per day
func (d Day) Length() int {
	return NPeriods
}

// TODO make GetPeriod Function

// TODO make GetDay Function
// TODO make GetPeriod Function

// TimeTable is struct to store timetable data
type TimeTable [MaxCap]Period
