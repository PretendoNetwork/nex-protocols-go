// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetPersistenceInfo == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfo not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	ownerID := types.NewPID(0)
	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetPersistenceInfo(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	persistenceSlotID := types.NewPrimitiveU16(0)
	err = persistenceSlotID.ExtractFrom(parametersStream)
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
