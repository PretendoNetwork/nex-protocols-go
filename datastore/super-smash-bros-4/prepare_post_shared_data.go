// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_smash_bros_4_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-smash-bros-4/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostSharedData sets the PreparePostSharedData handler function
func (protocol *Protocol) PreparePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *datastore_super_smash_bros_4_types.DataStorePreparePostSharedDataParam)) {
	protocol.PreparePostSharedDataHandler = handler
}

func (protocol *Protocol) handlePreparePostSharedData(packet nex.PacketInterface) {
	if protocol.PreparePostSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::PreparePostSharedData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_super_smash_bros_4_types.NewDataStorePreparePostSharedDataParam())
	if err != nil {
		go protocol.PreparePostSharedDataHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PreparePostSharedDataHandler(nil, client, callID, param.(*datastore_super_smash_bros_4_types.DataStorePreparePostSharedDataParam))
}
