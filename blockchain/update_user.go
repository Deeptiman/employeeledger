package blockchain

import (
	"fmt"

	"employeeledger/chaincode/model"
)

func (user *FabricUser) UpdateUserDataFromLedger(userId, name, email, company,
	occupation, salary, userType string) (*model.UserData, error) {

	var args []string
	args = append(args, "invoke")
	args = append(args, "updateUserData")
	args = append(args, userId)
	args = append(args, name)
	args = append(args, email)
	args = append(args, company)
	args = append(args, occupation)
	args = append(args, salary)
	args = append(args, userType)

	eventID := "updateUserDataInvoke"

	_, err := user.ExecuteTranctionEvent(eventID, args[0],
		[][]byte{
			[]byte(args[1]), []byte(args[2]), []byte(args[3]),
			[]byte(args[4]), []byte(args[5]), []byte(args[6]),
			[]byte(args[7]), []byte(args[8]),
		})

	if err != nil {
		return nil, fmt.Errorf("Error - UpdateUserDataFromLedger : %s", err.Error())
	}

	return nil, nil
}
