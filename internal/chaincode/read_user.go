package chaincode

import (
	"fmt"

	"github.com/Deeptiman/employeeledger/chaincode/model"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (t *HelloWorldServiceChaincode) readUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var userData model.UserData

	indexName := "email"
	userNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{args[1]})

	err = getFromLedger(stub, userNameIndexKey, args[1], &userData)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve userData in the ledger: %v", err))
	}
	userAsByte, err := objectToByte(userData)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable convert the userData to byte: %v", err))
	}

	err = stub.SetEvent("getUserInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(userAsByte)
}

func (t *HelloWorldServiceChaincode) readAllUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("##### Read All User #####")

	indexName := "email"
	iterator, err := stub.GetStateByPartialCompositeKey(indexName, []string{})
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to retrieve the list of resource in the ledger: %v", err))
	}

	allUsersData := make([]model.UserData, 0)

	for iterator.HasNext() {
		keyValueState, errIt := iterator.Next()
		if errIt != nil {
			return shim.Error(fmt.Sprintf("Unable to retrieve a user in the ledger: %v", errIt))
		}
		var userdata model.UserData
		err = byteToObject(keyValueState.Value, &userdata)
		if err != nil {
			return shim.Error(fmt.Sprintf("Unable to convert a user: %v", err))
		}

		allUsersData = append(allUsersData, userdata)
	}

	allUsersAsByte, err := objectToByte(allUsersData)
	if err != nil {
		return shim.Error(fmt.Sprintf("Unable to convert the users list to byte: %v", err))
	}

	err = stub.SetEvent("getAllUsersInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(allUsersAsByte)
}
