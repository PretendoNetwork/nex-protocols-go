// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetPersistenceInfo == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfo not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	ownerID, err := parametersStream.ReadPID()
	if err != nil {
		_, errorCode = protocol.GetPersistenceInfo(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	persistenceSlotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		_, errorCode = protocol.GetPersistenceInfo(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPersistenceInfo(nil, packet, callID, ownerID, persistenceSlotID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
