package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// RankingMK8ProtocolID is the protocol ID for the Ranking (Mario Kart 8) protocol. ID is the same as the Ranking protocol
	RankingMK8ProtocolID = 0x70
)

// RankingMK8Protocol handles the Ranking (Mario Kart 8) nex protocol. Embeds RankingProtocol
type RankingMK8Protocol struct {
	server *nex.Server
	RankingProtocol
}

// Setup initializes the protocol
func (rankingMK8Protocol *RankingMK8Protocol) Setup() {
	nexServer := rankingMK8Protocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if RankingProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			default:
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (rankingMK8Protocol *RankingMK8Protocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(RankingMK8ProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	rankingMK8Protocol.server.Send(responsePacket)
}

// NewRankingMK8Protocol returns a new RankingMK8Protocol
func NewRankingMK8Protocol(server *nex.Server) *RankingMK8Protocol {
	rankingMK8Protocol := &RankingMK8Protocol{server: server}
	rankingMK8Protocol.RankingProtocol.server = server

	rankingMK8Protocol.Setup()

	return rankingMK8Protocol
}
