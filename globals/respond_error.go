// Package globals implements variables and functions used by all protocol packages
package globals

import "github.com/PretendoNetwork/nex-go"

// RespondError sends the client a given error code
func RespondError(packet nex.PacketInterface, protocolID uint16, errorCode uint32) {
	client := packet.Sender()
	request := packet.RMCMessage()

	var responsePacket nex.PacketInterface
	var rmcResponseBytes []byte

	// TODO - Add HPP support back once nex-go supports it again
	if packet, ok := packet.(nex.PRUDPPacketInterface); ok {
		rmcResponse := nex.NewRMCError(errorCode)
		rmcResponse.ProtocolID = request.ProtocolID
		rmcResponse.CallID = request.CallID

		rmcResponseBytes = rmcResponse.Bytes()

		// * Go won't type assert responsePacket in the version check below,
		// * so to avoid a bunch of assertions just create a temp variable
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
		prudpPacket.SetSourceStreamType(packet.DestinationStreamType())
		prudpPacket.SetSourcePort(packet.DestinationPort())
		prudpPacket.SetDestinationStreamType(packet.SourceStreamType())
		prudpPacket.SetDestinationPort(packet.SourcePort())

		responsePacket = prudpPacket
	}

	responsePacket.SetPayload(rmcResponseBytes)

	client.Server().Send(responsePacket)
}
