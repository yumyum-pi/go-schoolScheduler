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

// Bytes return the byte array of period data
func (p *Period) Bytes() (b [PeriodByteSize]byte) {
	cID := (*p).ClassID.Bytes()
	sID := (*p).SubjectID.Bytes()
	tID := (*p).TeacherID.Bytes()
	cs := ClassIDBS + SubjectIDBS
	for i := 0; i < PeriodByteSize; i++ {
		if i < ClassIDBS {
			b[i] = cID[i]
		} else if i >= ClassIDBS && i < cs {
			j := i - ClassIDBS
			b[i] = sID[j]
		} else {
			j := i - cs
			b[i] = tID[j]
		}
	}

	return
}

// Init return the byte array of period data
// TODO write test
func (p *Period) Init(b [PeriodByteSize]byte) {
	var cID [ClassIDBS]byte
	var sID [SubjectIDBS]byte
	var tID [TeacherIDBS]byte
	cs := ClassIDBS + SubjectIDBS
	for i := 0; i < PeriodByteSize; {
		if i < ClassIDBS {
			copy(cID[:], b[:ClassIDBS])
			(*p).ClassID.Init(cID)
			i += ClassIDBS
		} else if i >= ClassIDBS && i < cs {
			copy(sID[:], b[ClassIDBS:cs])
			(*p).SubjectID.Init(sID)
			i = cs
		} else {
			copy(tID[:], b[cs:])
			(*p).TeacherID.Init(tID)
			i = PeriodByteSize
		}
	}

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
