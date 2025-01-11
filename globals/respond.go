// Package globals implements variables and functions used by all protocol packages
package globals

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/constants"
)

// Respond sends the client a given RMC message
func Respond(packet nex.PacketInterface, message *nex.RMCMessage) {
	sender := packet.Sender()

	var responsePacket nex.PacketInterface

	switch packet := packet.(type) {
	case nex.PRUDPPacketInterface:
		var prudpPacket nex.PRUDPPacketInterface

		connection := sender.(*nex.PRUDPConnection)
		endpoint := connection.Endpoint()
		server := endpoint.(*nex.PRUDPEndPoint).Server
		if packet.Version() == 1 {
			prudpPacket, _ = nex.NewPRUDPPacketV1(server, connection, nil)
		} else {
			prudpPacket, _ = nex.NewPRUDPPacketV0(server, connection, nil)
		}

		prudpPacket.SetType(constants.DataPacket)

		if packet.HasFlag(constants.PacketFlagReliable) {
			prudpPacket.AddFlag(constants.PacketFlagReliable)
		}

		prudpPacket.AddFlag(constants.PacketFlagNeedsAck)
		prudpPacket.SetSourceVirtualPortStreamType(packet.DestinationVirtualPortStreamType())
		prudpPacket.SetSourceVirtualPortStreamID(packet.DestinationVirtualPortStreamID())
		prudpPacket.SetDestinationVirtualPortStreamType(packet.SourceVirtualPortStreamType())
		prudpPacket.SetDestinationVirtualPortStreamID(packet.SourceVirtualPortStreamID())
		prudpPacket.SetSubstreamID(packet.SubstreamID())

		responsePacket = prudpPacket
		responsePacket.SetPayload(message.Bytes())

		sender.Endpoint().Send(responsePacket)

		connection.QueuedOutboundPackets.Clear(func(packet *nex.PRUDPPacketInterface) {
			server.Send(*packet)
		});
	case *nex.HPPPacket:
		// * We reuse the same packet from input and replace
		// * the RMC message so that it can be delivered back
		responsePacket = packet
		responsePacket.SetRMCMessage(message)

		sender.Endpoint().Send(responsePacket)
	}

	if message.OnAfterSend != nil {
		message.OnAfterSend()
	}
}
