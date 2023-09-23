package blockchain

import (
	"fmt"
)

func (user *FabricUser) addUserToLedger(nameValue, emailValue, companyValue, occupationValue, salaryValue, userType string) error {

	fmt.Println("BlockChain :: Add User Input nameValue = " + nameValue + " , emailValue = " + emailValue + " , companyValue = " + companyValue + " , occupationValue = " + occupationValue + " , salaryValue = " + salaryValue + " , userType = " + userType)

	var args []string
	args = append(args, "invoke")
	args = append(args, "createUser")
	args = append(args, nameValue)
	args = append(args, emailValue)
	args = append(args, companyValue)
	args = append(args, occupationValue)
	args = append(args, salaryValue)
	args = append(args, userType)

	eventID := "addUserInvoke"

	_, err := user.ExecuteTranctionEvent(eventID, args[0],
		[][]byte{
			[]byte(args[1]), []byte(args[2]), []byte(args[3]),
			[]byte(args[4]), []byte(args[5]), []byte(args[6]),
			[]byte(args[7]),
		})

	if err != nil {
		return fmt.Errorf("Error - addUserToLedger : %s", err.Error())
	}

	return nil
}
