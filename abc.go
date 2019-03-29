package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	err := stub.PutState("counter", []byte("0"))
	if err != nil {
		return shim.Error("Failed to set counter")
	}
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "incrementCounter" {
		return t.incrementCounter(stub, args)
	} else if fn == "decrementCounter" {
		return t.decrementCounter(stub, args)
	} else if fn == "getCounter" {
		return t.getCounter(stub)
	}

	resp := shim.Error("Invalid function name : " + fn)
	resp.Status = 404
	return resp
}

// incrementCounter increments the counter value by 1
func (t *SimpleAsset) incrementCounter(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	value, err := stub.GetState("counter")
	counter, _ := strconv.Atoi(string(value))
	err = stub.PutState("counter", []byte(string(counter+1)))
	if err != nil {
		return shim.Error("Failed to set counter")
	}
	return shim.Success([]byte("Updated counter value: " + string(counter+1)))
}

// decrementCounter reduces the counter value by 1
func (t *SimpleAsset) decrementCounter(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	value, err := stub.GetState("counter")
	counter, _ := strconv.Atoi(string(value))
	err = stub.PutState("counter", []byte(string(counter-1)))
	if err != nil {
		return shim.Error("Failed to set counter")
	}
	return shim.Success([]byte("Updated counter value: " + string(counter-1)))
}

// getCounter returns the value of the counter
func (t *SimpleAsset) getCounter(stub shim.ChaincodeStubInterface) peer.Response {
	value, err := stub.GetState("counter")
	if err != nil {
		return shim.Error("Failed to get counter")
	}
	return shim.Success(value)
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
