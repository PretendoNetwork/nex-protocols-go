// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePicture(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdatePicture == nil {
		globals.Logger.Warning("Friends3DS::UpdatePicture not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	unknown := types.NewPrimitiveU32(0)
	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePicture(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	picture := types.NewBuffer(nil)
	err = picture.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdatePicture(fmt.Errorf("Failed to read picture from parameters. %s", err.Error()), packet, callID, 0, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdatePicture(nil, packet, callID, unknown, picture)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
