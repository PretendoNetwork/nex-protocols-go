package globals

import "github.com/PretendoNetwork/nex-go"

func RespondNotImplemented(packet nex.PacketInterface, protocolID uint8) {
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
		rmcResponse := nex.NewRMCResponse(protocolID, request.CallID())
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