package requestlist

import (
	"testing"

	rl "github.com/yumyum-pi/go-schoolScheduler/internal/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/test/values"
)

var classRL rl.Class = rl.Class{
	values.ClassIDs[0].Bytes(): 9,
	values.ClassIDs[1].Bytes(): 10,
	values.ClassIDs[2].Bytes(): 4,
	values.ClassIDs[3].Bytes(): 3,
	values.ClassIDs[4].Bytes(): 8,
	values.ClassIDs[5].Bytes(): 12,
	values.ClassIDs[6].Bytes(): 2,
	values.ClassIDs[7].Bytes(): 5,
	values.ClassIDs[8].Bytes(): 1,
	values.ClassIDs[9].Bytes(): 6,
}

// sum of all the request in classRL
var totalReq = 60

// TestClass_TotalReq_Pass test if TotalReq func return proper value
func TestClass_TotalReq_Pass(t *testing.T) {
	// check if function return the actual value
	if classRL.TotalReq() != totalReq {
		t.Errorf("> Error: totalReq=%v\ti=%v\n", totalReq, classRL.TotalReq())
	}

	// check if function return the actual value
	if classRL.TotalReq() > totalReq || classRL.TotalReq() < totalReq {
		t.Errorf("> Error: totalReq=%v\ti=%v\n", totalReq, classRL.TotalReq())
	}
}
