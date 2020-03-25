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

// SubjectIDs is a slice to hold subject ids
type SubjectIDs []SubjectID

// SubTypeE is a struct to hold subject type and its frequency
type SubTypeE struct {
	Type [4]byte
	freq int
}

// SubType is a slice made up of SubTypeE
type SubType []SubTypeE

// Find return the index no. of the given subjectID
func (s *SubType) Find(sub SubjectID) int {
	// loop to check each item in the slice
	for i, subTE := range *s {
		// check if ids are same
		if subTE.Type == sub.Type {
			return i // return if found a match
		}
	}
	return -1
}

// Add func add the subject id
func (s *SubType) Add(sub SubjectID) {
	// check if the given id is in the slice
	i := (*s).Find(sub)

	// check if match is found
	if i != -1 {
		// match as found
		(*s)[i].freq++ // add to the frequency
	} else {
		// match not found
		(*s) = append((*s), SubTypeE{sub.Type, 1}) // add new element to the slice
	}

}

// Types return an slice of types of subject in the slice
func (ids *SubjectIDs) Types() SubType {
	var types SubType
	for _, id := range *ids {
		types.Add(id)
	}

	return types
}

// Subject is a struct to store subject data
type Subject struct {
	ID         SubjectID `json:"id"`         // unique identifier for a subject
	ReqClasses int       `json:"reqClasses"` // required classes per week
}
