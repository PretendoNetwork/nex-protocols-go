// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinOrCreateMatchmakeSession sets the JoinOrCreateMatchmakeSession handler function
func (protocol *Protocol) JoinOrCreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.joinOrCreateMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleJoinOrCreateMatchmakeSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinOrCreateMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinOrCreateMatchmakeSession STUBBED")

	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.joinOrCreateMatchmakeSessionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
