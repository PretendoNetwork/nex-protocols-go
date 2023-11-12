// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TouchObject sets the TouchObject handler function
func (protocol *Protocol) TouchObject(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreTouchObjectParam) uint32) {
	protocol.touchObjectHandler = handler
}

func (protocol *Protocol) handleTouchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.touchObjectHandler == nil {
		globals.Logger.Warning("DataStore::TouchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreTouchObjectParam())
	if err != nil {
		errorCode = protocol.touchObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.touchObjectHandler(nil, packet, callID, param.(*datastore_types.DataStoreTouchObjectParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
