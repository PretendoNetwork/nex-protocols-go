// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendInvitation sets the SendInvitation handler function
func (protocol *Protocol) SendInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID) uint32) {
	protocol.sendInvitationHandler = handler
}

func (protocol *Protocol) handleSendInvitation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.sendInvitationHandler == nil {
		globals.Logger.Warning("Friends3DS::SendInvitation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListPID()
	if err != nil {
		errorCode = protocol.sendInvitationHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.sendInvitationHandler(nil, packet, callID, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
