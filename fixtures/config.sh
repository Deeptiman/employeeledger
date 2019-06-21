export FABRIC_CFG_PATH=$PWD

./bin/cryptogen generate --config=./crypto-config.yaml

./bin/configtxgen -profile EmployeeLedger -outputBlock ./artifacts/orderer.genesis.block

./bin/configtxgen -profile EmployeeLedger -outputCreateChannelTx ./artifacts/employeeledger.channel.tx -channelID employeeledger

./bin/configtxgen -profile EmployeeLedger -outputAnchorPeersUpdate ./artifacts/org1.employeeledger.anchors.tx -channelID employeeledger -asOrg EmployeeLedgerOrganization1
