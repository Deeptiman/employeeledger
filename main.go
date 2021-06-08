package main

import (
	"fmt"
	"employeeledger/blockchain"
	"employeeledger/web"
    "employeeledger/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: 		"orderer.employee.ledger.com",

		// Channel parameters
		ChannelID:      	"employeeledger",
		ChannelConfig:  	os.Getenv("GOPATH") + "/src/employeeledger/fixtures/artifacts/employeeledger.channel.tx",

		// Chaincode parameters
		ChaincodeID:    	"employeeledger",
		ChaincodeGoPath: 	os.Getenv("GOPATH"),
		ChaincodePath:   	"employeeledger/chaincode/",
		OrgAdmin:        	"Admin",
		OrgName:         	"org1",
		ConfigFile:      	"config.yaml",

		// CA parameters
		CaID:                	"ca.org1.employee.ledger.com",	
		
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	

	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}



