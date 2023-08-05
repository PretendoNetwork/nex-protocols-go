// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinOrCreateMatchmakeSession sets the JoinOrCreateMatchmakeSession handler function
func (protocol *Protocol) JoinOrCreateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte) uint32) {
	protocol.joinOrCreateMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleJoinOrCreateMatchmakeSession(packet nex.PacketInterface) {
	if protocol.joinOrCreateMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.joinOrCreateMatchmakeSessionHandler(nil, client, callID, packet.Payload())
}
