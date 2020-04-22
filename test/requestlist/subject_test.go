package requestlist

/*
var tRls = rl.Subject{
	values.SubjectIDs[0].Bytes(): rl.Class{
		values.ClassIDs[0].Bytes(): 6,
		values.ClassIDs[1].Bytes(): 6,
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[1].Bytes(): rl.Class{
		values.ClassIDs[1].Bytes(): 6,
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[2].Bytes(): rl.Class{
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[3].Bytes(): rl.Class{
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[4].Bytes(): rl.Class{
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[5].Bytes(): rl.Class{
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[6].Bytes(): rl.Class{
		values.ClassIDs[6].Bytes(): 3,
		values.ClassIDs[7].Bytes(): 3,
		values.ClassIDs[8].Bytes(): 3,
		values.ClassIDs[9].Bytes(): 3,
	}, values.SubjectIDs[7].Bytes(): rl.Class{
		values.ClassIDs[7].Bytes(): 3,
		values.ClassIDs[8].Bytes(): 3,
		values.ClassIDs[9].Bytes(): 3,
	}, values.SubjectIDs[8].Bytes(): rl.Class{
		values.ClassIDs[8].Bytes(): 3,
		values.ClassIDs[9].Bytes(): 3,
	}, values.SubjectIDs[9].Bytes(): rl.Class{
		values.ClassIDs[9].Bytes(): 3,
	},
}

var tSIDFail = [6]byte{1, 2, 3, 4, 5, 6}
var tFailSID, tFailCID = 8, 6
var tFailSID2, tFailCID2 = 6, 6

var tRlsR = rl.Subject{
	values.SubjectIDs[0].Bytes(): rl.Class{
		values.ClassIDs[0].Bytes(): 6,
		values.ClassIDs[1].Bytes(): 6,
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[1].Bytes(): rl.Class{
		values.ClassIDs[1].Bytes(): 6,
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[2].Bytes(): rl.Class{
		values.ClassIDs[2].Bytes(): 6,
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[3].Bytes(): rl.Class{
		values.ClassIDs[3].Bytes(): 6,
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[4].Bytes(): rl.Class{
		values.ClassIDs[4].Bytes(): 6,
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[5].Bytes(): rl.Class{
		values.ClassIDs[5].Bytes(): 6,
		values.ClassIDs[6].Bytes(): 6,
		values.ClassIDs[7].Bytes(): 6,
		values.ClassIDs[8].Bytes(): 6,
		values.ClassIDs[9].Bytes(): 6,
	}, values.SubjectIDs[tFailSID2].Bytes(): rl.Class{
		values.ClassIDs[tFailCID].Bytes(): 6,
		values.ClassIDs[7].Bytes():        3,
		values.ClassIDs[8].Bytes():        3,
		values.ClassIDs[9].Bytes():        3,
	}, values.SubjectIDs[tFailSID].Bytes(): rl.Class{
		values.ClassIDs[tFailCID].Bytes(): 3,
		values.ClassIDs[8].Bytes():        3,
		values.ClassIDs[9].Bytes():        3,
	}, values.SubjectIDs[8].Bytes(): rl.Class{
		values.ClassIDs[8].Bytes(): 3,
		values.ClassIDs[9].Bytes(): 3,
	}, values.SubjectIDs[9].Bytes(): rl.Class{
		values.ClassIDs[9].Bytes(): 3,
	}, tSIDFail: rl.Class{
		values.ClassIDs[9].Bytes(): 3,
	},
}

func TestInit_Pass(t *testing.T) {
	rls := make(rl.Subject)
	rls.Init(&values.Classes)

	// check if the value match the test
	for tSID, tCC := range tRls {
		cc, ok := rls[tSID] // get the class map
		// check class map exist with the given key
		if !ok {
			t.Errorf("> Error: tSID=%v not found.\n", tSID)
			continue
		}

		// loop through all the classIDs
		for tCID, tReq := range tCC {
			// get the request
			req, ok := cc[tCID]
			// check if the given classID exist in the map
			if !ok {
				t.Errorf("> Error: tCID=%v not found.\n", tCID)
				continue
			}

			// check if the request matchs
			if req != tReq {
				t.Errorf("> Error: tReq=%v req=%v.\n", tReq, req)
				continue
			}
		}
	}
}

func TestInit_Fail(t *testing.T) {
	rls := make(rl.Subject)
	rls.Init(&values.Classes)

	// check if the value match the test
	for tSID, tCC := range tRlsR {
		cc, ok := rls[tSID] // get the class map
		if tSID == tSIDFail {
			// check class map exist with the given key
			if ok {
				t.Errorf("> Error: tSID=%v unknown found.\n", tSID)
			}
			continue
		}

		// check class map exist with the given key
		if !ok {
			t.Errorf("> Error: tSID=%v not found.\n", tSID)
			continue
		}

		// loop through all the classIDs
		for tCID, tReq := range tCC {
			req, ok := cc[tCID] // get the request
			if tSID == values.SubjectIDs[tFailSID].Bytes() && tCID == values.ClassIDs[tFailCID].Bytes() {
				// check class map exist with the given key
				if ok {
					t.Errorf("> Error: unknown tCID=%v at tSID=%v.\n", tCID, tSID)
				}
				continue
			}

			// check if the given classID exist in the map
			if !ok {
				t.Errorf("> Error: tCID=%v not found.\n", tCID)
				continue
			}

			// check if the value are of error units
			if tSID == values.SubjectIDs[tFailSID2].Bytes() && tCID == values.ClassIDs[tFailCID2].Bytes() {
				// the error unit should not be equal
				if req == tReq {
					t.Errorf("> Error: equal request tReq=%v req=%v at tCID=%v & tSID=%v\n", tReq, req, tCID, tSID)
				}
				continue
			}
			// check if the request matchs
			if req != tReq {
				t.Errorf("> Error: unequal request tReq=%v req=%v at tCID=%v & tSID=%v\n", tReq, req, tCID, tSID)
				continue
			}
		}
	}
}
*/
