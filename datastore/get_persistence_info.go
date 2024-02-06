// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	if protocol.GetPersistenceInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::GetPersistenceInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	ownerID := types.NewPID(0)
	persistenceSlotID := types.NewPrimitiveU16(0)

	var err error

	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPersistenceInfo(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = persistenceSlotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPersistenceInfo(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPersistenceInfo(nil, packet, callID, ownerID, persistenceSlotID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
