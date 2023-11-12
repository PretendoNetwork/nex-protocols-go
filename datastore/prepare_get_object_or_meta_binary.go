// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObjectOrMetaBinary sets the PrepareGetObjectOrMetaBinary handler function
func (protocol *Protocol) PrepareGetObjectOrMetaBinary(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStorePrepareGetParam) uint32) {
	protocol.prepareGetObjectOrMetaBinaryHandler = handler
}

func (protocol *Protocol) handlePrepareGetObjectOrMetaBinary(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.prepareGetObjectOrMetaBinaryHandler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObjectOrMetaBinary not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParam())
	if err != nil {
		errorCode = protocol.prepareGetObjectOrMetaBinaryHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.prepareGetObjectOrMetaBinaryHandler(nil, packet, callID, param.(*datastore_types.DataStorePrepareGetParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
