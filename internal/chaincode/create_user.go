package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/Deeptiman/employeeledger/chaincode/model"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *HelloWorldServiceChaincode) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var name, email, company, occupation, salary, userType string

	name = args[1]
	email = args[2]
	company = args[3]
	occupation = args[4]
	salary = args[5]
	userType = args[6]

	fmt.Println("Create User :>>>> " + name + " == " + email + " == " + company + " == " + occupation + " == " + salary + " == " + userType)

	_, found, err := cid.GetAttributeValue(stub, "usermode")
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to identify the type of the request owner: %v", err))
	}

	if !found {
		return shim.Error("The type of the request owner is not present")
	}

	userID, err := cid.GetID(stub)

	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to identify the ID of the request owner: %v", err))
	}

	fmt.Println("userID = " + userID)

	userdata := &model.UserData{userID, name, email, company, occupation, salary, userType}

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

	err = stub.SetEvent("addUserInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("Chaincode --- Name = %s " + name + " , Email = %s " + email + " , Company = %s " + company + " , Occupation = %s " + occupation + " , Salary = %s " + salary + " , UserType = %s " + userType)

	return shim.Success(nil)
}
