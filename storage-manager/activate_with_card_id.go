// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleActivateWithCardID(packet nex.PacketInterface) {
	var err error

	if protocol.ActivateWithCardID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "StorageManager::ActivateWithCardID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown := types.NewPrimitiveU8(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateWithCardID(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	cardID := types.NewPrimitiveU64(0)
	err = cardID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ActivateWithCardID(fmt.Errorf("Failed to read cardID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ActivateWithCardID(nil, packet, callID, unknown, cardID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
