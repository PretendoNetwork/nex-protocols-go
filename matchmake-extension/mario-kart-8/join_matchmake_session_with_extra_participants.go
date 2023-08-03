// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSessionWithExtraParticipants sets the JoinMatchmakeSessionWithExtraParticipants handler function
func (protocol *Protocol) JoinMatchmakeSessionWithExtraParticipants(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.joinMatchmakeSessionWithExtraParticipantsHandler = handler
}

func (protocol *Protocol) handleJoinMatchmakeSessionWithExtraParticipants(packet nex.PacketInterface) {
	if protocol.joinMatchmakeSessionWithExtraParticipantsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionMarioKart8::JoinMatchmakeSessionWithExtraParticipants STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.joinMatchmakeSessionWithExtraParticipantsHandler(nil, client, callID, packet.Payload())
}
