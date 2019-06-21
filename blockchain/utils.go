package blockchain


import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)


func (user *FabricUser) ExecuteTranctionEvent(eventID, fcnName string, args [][]byte) (*channel.Response, error) {

    reg, notifier, err := user.Fabric.event.RegisterChaincodeEvent(user.Fabric.ChaincodeID, eventID)
   
    if err != nil {
	return nil, fmt.Errorf("Blockchain ..... failed to register event: %v", err)
    }
    defer user.Fabric.event.Unregister(reg)

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data to invoke in the ledger")

	response, err := user.ChannelClient.Execute(channel.Request{
		
			ChaincodeID: user.Fabric.ChaincodeID, 
			Fcn: fcnName, 
			Args: args, 
			TransientMap: transientDataMap,
	})

	if err != nil {
		fmt.Printf("failed to execute event: %s", err.Error())
		return nil, nil
	}

	select {
		case ccEvent := <-notifier:
			fmt.Printf("Received CC event: %v\n", ccEvent)
		case <-time.After(time.Second * 60):
			return nil, fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	return &response, nil
}
