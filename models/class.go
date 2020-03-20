package models

// Class id has should have the following components
// - class: (1 - 12)	= requires 2 byte of data
// - section: (1 - n)	= requires 2 byte of data
// - group: (1 - n)		= requires 2 bye of data

// ClassID is struct to store class id
type ClassID struct {
	Standerd [2]byte `json:"stn"`
	Section  [2]byte `json:"sec"`
	Group    [2]byte `json:"grp"`
}

// Bytes return combined bytes of the ID
func (i *ClassID) Bytes() [6]byte {
	return [6]byte{i.Standerd[0], i.Standerd[1], i.Section[0], i.Section[1], i.Group[0], i.Group[1]}
}

// Create return a new ClassID
func (i ClassID) Create(st, sec, grp [2]byte) ClassID {
	return ClassID{
		Standerd: st,
		Section:  sec,
		Group:    grp,
	}
}

// Class in an struct to hold class data
type Class struct {
	ID       ClassID   `json:"id"`       // Unique identifier for a class
	Student  []Student `json:"student"`  // List of students in this class
	Subjects []Subject `json:"subjects"` // List of subject to be studied in this class
}
