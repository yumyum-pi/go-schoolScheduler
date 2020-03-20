package models

// SubjectID is a unique identifier for a subject
//TODO write test
type SubjectID struct {
	Standerd [2]byte `json:"stnID"` // ID for the standerd of the subject
	Type     [4]byte `json:"type"`  // type of subject. Example English, Hindi
}

// Init adds value to the SubjectID
func (id *SubjectID) Init(seq [6]byte) {
	id.Standerd = [2]byte{
		seq[0],
		seq[1],
	}
	id.Type = [4]byte{
		seq[2],
		seq[3],
		seq[4],
		seq[5],
	}
}

// Byte func updates the byte value from Standerd and Type values
func (id *SubjectID) Byte() [6]byte {
	return [6]byte{id.Standerd[0], id.Standerd[1], id.Type[0], id.Type[1], id.Type[2], id.Type[3]}
}

// Subject is a struct to store subject data
type Subject struct {
	ID         SubjectID `json:"id"`         // unique identifier for a subject
	ReqClasses int       `json:"reqClasses"` // required classes per week
}
