package models

// Period is a single cell in the timetable
type Period struct {
	ID       [2]int8
	Students []Student
	Subject  Subject
	Teacher  Teacher
}

// Day is a collection of 8 classes
type Day [8]Period

// Length return the number of Period per day
func (d Day) Length() int {
	return len(d)
}

// TODO make GetPeriod Function

// Week is a collection of 6 days
type Week [6]Day

// Length return the number of Period per week
func (w Week) Length() int {
	return len(w) * w[0].Length()
}

// TODO make GetDay Function
// TODO make GetPeriod Function
