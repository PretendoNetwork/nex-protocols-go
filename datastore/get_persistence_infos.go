// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPersistenceInfos(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetPersistenceInfos == nil {
		globals.Logger.Warning("DataStore::GetPersistenceInfos not implemented")
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
		_, errorCode = protocol.GetPersistenceInfos(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	persistenceSlotIDs := types.NewList[*types.PrimitiveU16]()
	persistenceSlotIDs.Type = types.NewPrimitiveU16(0)
	err = persistenceSlotIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetPersistenceInfos(fmt.Errorf("Failed to read persistenceSlotIDs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPersistenceInfos(nil, packet, callID, ownerID, persistenceSlotIDs)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
