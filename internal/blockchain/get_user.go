package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/Deeptiman/employeeledger/chaincode/model"
)

func (user *FabricUser) GetUserFromLedger(emailValue string) (*model.UserData, error) {

	fmt.Println("BlockChain :: getUserFromLedger : " + emailValue)

	var args []string
	args = append(args, "invoke")
	args = append(args, "readUser")
	args = append(args, emailValue)

	eventID := "getUserInvoke"

	response, err := user.ExecuteTranctionEvent(eventID, args[0],
		[][]byte{
			[]byte(args[1]), []byte(args[2]),
		})

	if err != nil {
		return nil, fmt.Errorf("Error - GetUserFromLedger : %s", err.Error())
	}

	fmt.Println("Response Received")

	var userData *model.UserData

	err = json.Unmarshal(response.Payload, &userData)
	if err != nil {
		return nil, fmt.Errorf("unable to convert response to the object given for the query: %v", err)
	}

	return userData, nil
}

func (user *FabricUser) GetAllUsersFromLedger() ([]model.UserData, error) {

	fmt.Println("BlockChain :: getAllUserFromLedger ")

	var args []string
	args = append(args, "invoke")
	args = append(args, "readAllUser")

	eventID := "getAllUsersInvoke"

	response, err := user.ExecuteTranctionEvent(eventID, args[0],
		[][]byte{
			[]byte(args[1]),
		})

	if err != nil {
		return nil, fmt.Errorf("Error - GetAllUsersFromLedger : %s", err.Error())
	}

	fmt.Println("Response Received")

	allUsersData := make([]model.UserData, 0)

	err = json.Unmarshal(response.Payload, &allUsersData)
	if err != nil {
		return nil, fmt.Errorf("unable to convert response to the object given for the query: %v", err)
	}

	return allUsersData, nil
}
