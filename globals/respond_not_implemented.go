package globals

import "github.com/PretendoNetwork/nex-go"

func RespondNotImplemented(packet nex.PacketInterface, protocolID uint8) {
	switch packet := packet.(type) {
	case *nex.HPPPacket:
		client := packet.Sender()
		request := packet.RMCRequest()

		rmcResponse := nex.NewRMCResponse(0, request.CallID())
		rmcResponse.SetError(nex.Errors.Core.NotImplemented)

		rmcResponseBytes := rmcResponse.Bytes()

		responsePacket, _ := nex.NewHPPPacket(client, nil)

		responsePacket.SetPayload(rmcResponseBytes)

		client.Server().Send(responsePacket)
	default:
		client := packet.Sender()
		request := packet.RMCRequest()

		rmcResponse := nex.NewRMCResponse(protocolID, request.CallID())
		rmcResponse.SetError(nex.Errors.Core.NotImplemented)

		rmcResponseBytes := rmcResponse.Bytes()

		var responsePacket nex.PacketInterface
		if packet.Version() == 1 {
			responsePacket, _ = nex.NewPacketV1(client, nil)
		} else {
			responsePacket, _ = nex.NewPacketV0(client, nil)
		}

		responsePacket.SetVersion(packet.Version())
		responsePacket.SetSource(packet.Destination())
		responsePacket.SetDestination(packet.Source())
		responsePacket.SetType(nex.DataPacket)
		responsePacket.SetPayload(rmcResponseBytes)

		responsePacket.AddFlag(nex.FlagNeedsAck)
		responsePacket.AddFlag(nex.FlagReliable)

		client.Server().Send(responsePacket)
	}
}
