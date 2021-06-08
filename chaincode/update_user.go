package main

import (
	"encoding/json"
	"fmt"

	"employeeledger/chaincode/model"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *HelloWorldServiceChaincode) updateUserData(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var userId, name, email, company, occupation, salary, userType string

	userId = args[1]
	name = args[2]
	email = args[3]
	company = args[4]
	occupation = args[5]
	salary = args[6]
	userType = args[7]

	fmt.Println("Chaincode Update User Data :>>>> " + userId + " == " + name + " == " +
		email + " == " + company + " == " + occupation + " == " + salary + " == " + userType)

	userdata := &model.UserData{userId, name, email, company, occupation, salary, userType}

	userDataJSONasBytes, err := json.Marshal(userdata)

	if err != nil {
		return shim.Error(err.Error())
	}

	indexName := "email"
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{userdata.Email})

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(userNameIndexKey, userDataJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.SetEvent("updateUserDataInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
