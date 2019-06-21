package blockchain

import (
	"fmt"
)

func (user *FabricUser) DeleteUserFromLedger(emailValue string) error {

	fmt.Println("BlockChain :: deleteUserFromLedger : " + emailValue)

	var args []string
	args = append(args, "invoke")
	args = append(args, "deleteUser")
	args = append(args, emailValue)

	eventID := "deleteUserInvoke"

	_, err := user.ExecuteTranctionEvent(eventID, args[0],
		[][]byte{
			[]byte(args[1]), []byte(args[2]),
		})

	if err != nil {
		return fmt.Errorf("Error - DeleteUserFromLedger : %s", err.Error())
	}

	return nil
}
