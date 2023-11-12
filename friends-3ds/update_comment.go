// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateComment sets the UpdateComment handler function
func (protocol *Protocol) UpdateComment(handler func(err error, packet nex.PacketInterface, callID uint32, comment string) uint32) {
	protocol.updateCommentHandler = handler
}

func (protocol *Protocol) handleUpdateComment(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateCommentHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateComment not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	comment, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.updateCommentHandler(fmt.Errorf("Failed to read comment from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateCommentHandler(nil, packet, callID, comment)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
