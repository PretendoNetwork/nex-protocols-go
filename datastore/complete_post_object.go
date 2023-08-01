// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObject sets the CompletePostObject handler function
func (protocol *Protocol) CompletePostObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompletePostParam)) {
	protocol.completePostObjectHandler = handler
}

func (protocol *Protocol) handleCompletePostObject(packet nex.PacketInterface) {
	if protocol.completePostObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		go protocol.completePostObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.completePostObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreCompletePostParam))
}