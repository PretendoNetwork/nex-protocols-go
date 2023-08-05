// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompleteUpdateObject sets the CompleteUpdateObject handler function
func (protocol *Protocol) CompleteUpdateObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreCompleteUpdateParam) uint32) {
	protocol.completeUpdateObjectHandler = handler
}

func (protocol *Protocol) handleCompleteUpdateObject(packet nex.PacketInterface) {
	if protocol.completeUpdateObjectHandler == nil {
		globals.Logger.Warning("DataStore::CompleteUpdateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreCompleteUpdateParam())
	if err != nil {
		go protocol.completeUpdateObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.completeUpdateObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreCompleteUpdateParam))
}
