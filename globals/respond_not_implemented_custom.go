// Package globals implements variables and functions used by all protocol packages
package globals

import "github.com/PretendoNetwork/nex-go"

// RespondNotImplementedCustom sends the client the Core::NotImplemented error for custom protocols
func RespondNotImplementedCustom(packet nex.PacketInterface, customID uint16) {
	client := packet.Sender()
	request := packet.RMCRequest()

	var responsePacket nex.PacketInterface
	var rmcResponseBytes []byte
	switch packet := packet.(type) {
	case *nex.HPPPacket:
		rmcResponse := nex.NewRMCResponse(0, request.CallID())
		rmcResponse.SetError(nex.Errors.Core.NotImplemented)

		rmcResponseBytes = rmcResponse.Bytes()

		responsePacket, _ = nex.NewHPPPacket(client, nil)
	default:
		rmcResponse := nex.NewRMCResponse(0x7F, request.CallID())
		rmcResponse.SetCustomID(customID)
		rmcResponse.SetError(nex.Errors.Core.NotImplemented)

		rmcResponseBytes = rmcResponse.Bytes()

		if packet.Version() == 1 {
			responsePacket, _ = nex.NewPacketV1(client, nil)
		} else {
			responsePacket, _ = nex.NewPacketV0(client, nil)
		}

		responsePacket.SetVersion(packet.Version())
		responsePacket.SetSource(packet.Destination())
		responsePacket.SetDestination(packet.Source())
		responsePacket.SetType(nex.DataPacket)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)
	}

	responsePacket.SetPayload(rmcResponseBytes)

	client.Server().Send(responsePacket)
}
