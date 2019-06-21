package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (t *HelloWorldServiceChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("ChainCode - DeleteUser === " + args[1])

	indexName := "email"
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{args[1]})

	err = deleteFromLedger(stub, userNameIndexKey, args[1])
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to delete the user in the ledger: %v", err))
	}

	err = stub.SetEvent("deleteUserInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
