package requestlist

import (
	"testing"
)

// sum of all the request in classRL
var totalReq = 60

// TestClass_TotalReq_Pass test if TotalReq func return proper value
func TestClass_TotalReq_Pass(t *testing.T) {
	// check if function return the actual value
	if TClassRL.TotalReq() != totalReq {
		t.Errorf("> Error: totalReq=%v\ti=%v\n", totalReq, TClassRL.TotalReq())
	}

	// check if function return the actual value
	if TClassRL.TotalReq() > totalReq || TClassRL.TotalReq() < totalReq {
		t.Errorf("> Error: totalReq=%v\ti=%v\n", totalReq, TClassRL.TotalReq())
	}
}
