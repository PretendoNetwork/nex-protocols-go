// Package match_making implements the Match Making NEX protocol
package match_making

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetInvitationsReceived sets the GetInvitationsReceived handler function
func (protocol *MatchMakingProtocol) GetInvitationsReceived(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getInvitationsReceivedHandler = handler
}

func (protocol *MatchMakingProtocol) handleGetInvitationsReceived(packet nex.PacketInterface) {
	if protocol.getInvitationsReceivedHandler == nil {
		globals.Logger.Warning("MatchMaking::GetInvitationsReceived not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getInvitationsReceivedHandler(nil, client, callID)
}
