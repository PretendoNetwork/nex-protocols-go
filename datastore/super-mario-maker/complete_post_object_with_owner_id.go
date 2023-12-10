// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCompletePostObjectWithOwnerID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CompletePostObjectWithOwnerID == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CompletePostObjectWithOwnerID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	ownerID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.CompletePostObjectWithOwnerID(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	param, err := nex.StreamReadStructure(parametersStream, datastore_types.NewDataStoreCompletePostParam())
	if err != nil {
		_, errorCode = protocol.CompletePostObjectWithOwnerID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CompletePostObjectWithOwnerID(nil, packet, callID, ownerID, param)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
