// Package protocol implements the Match Making protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetInvitationsReceived sets the GetInvitationsReceived handler function
func (protocol *Protocol) GetInvitationsReceived(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getInvitationsReceivedHandler = handler
}

func (protocol *Protocol) handleGetInvitationsReceived(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getInvitationsReceivedHandler == nil {
		globals.Logger.Warning("MatchMaking::GetInvitationsReceived not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getInvitationsReceivedHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
