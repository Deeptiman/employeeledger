package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func getFromLedger(stub shim.ChaincodeStubInterface, key string, id string, result interface{}) error {

	fmt.Println("getFromLedger :: " + id)

	resultAsByte, err := stub.GetState(key)
	if err != nil {
		return fmt.Errorf("unable to retrieve the object in the ledger: %v", err)
	}
	if resultAsByte == nil {
		return fmt.Errorf("the object doesn't exist in the ledger")
	}
	err = byteToObject(resultAsByte, result)
	if err != nil {
		return fmt.Errorf("unable to convert the result to object: %v", err)
	}
	return nil
}

// deleteFromLedger delete an object in the ledger
func deleteFromLedger(stub shim.ChaincodeStubInterface, key string, id string) error {

	err := stub.DelState(key)
	if err != nil {
		return fmt.Errorf("unable to delete the object in the ledger: %v", err)
	}
	return nil
}

func objectToByte(object interface{}) ([]byte, error) {
	objectAsByte, err := json.Marshal(object)
	if err != nil {
		return nil, fmt.Errorf("unable convert the object to byte: %v", err)
	}
	return objectAsByte, nil
}

func byteToObject(objectAsByte []byte, result interface{}) error {
	err := json.Unmarshal(objectAsByte, result)
	if err != nil {
		return fmt.Errorf("unable to convert the result to object: %v", err)
	}
	return nil
}
