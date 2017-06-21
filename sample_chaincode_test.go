package main

import (
	//"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"testing"
)

func TestCreateAppointment(t *testing.T) {

	fmt.Println("Inside test method")
	//attribute := make(map[string][]byte)

	stub := shim.NewMockStub("mockstub", new(SampleChaincode))

	if stub == nil {
		t.Fatalf("mockStub creation failed")
	}

	stub.MockTransactionStart("t1")

	_, err := CreateAppointment(stub, []string{})

	if err == nil {
		t.Fatalf("Expected create Appointment to return error")
	}

	stub.MockTransactionEnd("t1")

}
