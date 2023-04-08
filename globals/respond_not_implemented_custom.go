package globals

import "github.com/PretendoNetwork/nex-go"

func RespondNotImplementedCustom(packet nex.PacketInterface, customID uint16) {
	client := packet.Sender()
	request := packet.RMCRequest()

	rmcResponse := nex.NewRMCResponse(0x7F, request.CallID())
	rmcResponse.SetCustomID(customID)
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
