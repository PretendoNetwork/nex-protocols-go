// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareUpdateObject sets the PrepareUpdateObject handler function
func (protocol *Protocol) PrepareUpdateObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareUpdateParam) uint32) {
	protocol.prepareUpdateObjectHandler = handler
}

func (protocol *Protocol) handlePrepareUpdateObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.prepareUpdateObjectHandler == nil {
		globals.Logger.Warning("DataStore::PrepareUpdateObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareUpdateParam())
	if err != nil {
		errorCode = protocol.prepareUpdateObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.prepareUpdateObjectHandler(nil, client, callID, param.(*datastore_types.DataStorePrepareUpdateParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
