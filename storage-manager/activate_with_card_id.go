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
	var errorCode uint32

	if protocol.ActivateWithCardID == nil {
		globals.Logger.Warning("StorageManager::ActivateWithCardID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown := types.NewPrimitiveU8(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ActivateWithCardID(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	cardID := types.NewPrimitiveU64(0)
	err = cardID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ActivateWithCardID(fmt.Errorf("Failed to read cardID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ActivateWithCardID(nil, packet, callID, unknown, cardID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
