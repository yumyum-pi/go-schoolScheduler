package models

// Decode timetable data
func (tt *TimeTable) Decode() (b []byte, maxcap int) {
	// calculate the capacity
	l := len((*tt).Period)
	// get the length of the last byte
	lb := len((*tt).Period[l-1])
	cap := (l - 1) * 32
	cap += lb

	b = make([]byte, 0, cap)
	// loop through each package byte
	for _, pkg := range (*tt).Period {
		b = append(b, pkg...)
	}

	// calculate maxcap
	maxcap = int((*tt).NDays * (*tt).NPeriods)
	return
}
