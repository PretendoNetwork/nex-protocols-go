// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceTournamentLeafNode sets the ReplaceTournamentLeafNode handler function
func (protocol *Protocol) ReplaceTournamentLeafNode(handler func(err error, client *nex.Client, callID uint32, packetPayload []byte)) {
	protocol.replaceTournamentLeafNodeHandler = handler
}

func (protocol *Protocol) handleReplaceTournamentLeafNode(packet nex.PacketInterface) {
	if protocol.replaceTournamentLeafNodeHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::ReplaceTournamentLeafNode not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::ReplaceTournamentLeafNode STUBBED")

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.replaceTournamentLeafNodeHandler(nil, client, callID, packet.Payload())
}
