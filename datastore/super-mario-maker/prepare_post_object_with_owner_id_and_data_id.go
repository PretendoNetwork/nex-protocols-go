// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handlePreparePostObjectWithOwnerIDAndDataID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PreparePostObjectWithOwnerIDAndDataID == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::PreparePostObjectWithOwnerIDAndDataID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := nex.StreamReadStructure(parametersStream, datastore_types.NewDataStorePreparePostParam())
	if err != nil {
		_, errorCode = protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PreparePostObjectWithOwnerIDAndDataID(nil, packet, callID, ownerID, dataID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
