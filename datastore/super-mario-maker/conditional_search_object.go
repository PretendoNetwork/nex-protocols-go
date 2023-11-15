// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleConditionalSearchObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ConditionalSearchObject == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ConditionalSearchObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	condition, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.ConditionalSearchObject(fmt.Errorf("Failed to read condition from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		_, errorCode = protocol.ConditionalSearchObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		_, errorCode = protocol.ConditionalSearchObject(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, 0, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ConditionalSearchObject(nil, packet, callID, condition, param.(*datastore_types.DataStoreSearchParam), extraData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
