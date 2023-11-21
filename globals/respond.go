// Package globals implements variables and functions used by all protocol packages
package globals

import "github.com/PretendoNetwork/nex-go"

// Respond sends the client a given RMC message
func Respond(packet nex.PacketInterface, message *nex.RMCMessage) {
	client := packet.Sender()

	var responsePacket nex.PacketInterface

	if packet, ok := packet.(nex.PRUDPPacketInterface); ok {
		var prudpPacket nex.PRUDPPacketInterface

		if packet.Version() == 1 {
			prudpPacket, _ = nex.NewPRUDPPacketV1(client.(*nex.PRUDPClient), nil)
		} else {
			prudpPacket, _ = nex.NewPRUDPPacketV0(client.(*nex.PRUDPClient), nil)
		}

		prudpPacket.SetType(nex.DataPacket)

		if packet.HasFlag(nex.FlagReliable) {
			prudpPacket.AddFlag(nex.FlagReliable)
		}

		prudpPacket.AddFlag(nex.FlagNeedsAck)
		prudpPacket.AddFlag(nex.FlagHasSize)
		prudpPacket.SetSourceStreamType(packet.DestinationStreamType())
		prudpPacket.SetSourcePort(packet.DestinationPort())
		prudpPacket.SetDestinationStreamType(packet.SourceStreamType())
		prudpPacket.SetDestinationPort(packet.SourcePort())
		prudpPacket.SetSubstreamID(packet.SubstreamID())

		responsePacket = prudpPacket
	}

	responsePacket.SetPayload(message.Bytes())

	client.Server().Send(responsePacket)
}
