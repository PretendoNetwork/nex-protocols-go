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
		request := packet.RMCRequest()

		if RankingProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			default:
				go respondNotImplemented(packet, RankingProtocolID)
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewRankingMK8Protocol returns a new RankingMK8Protocol
func NewRankingMK8Protocol(server *nex.Server) *RankingMK8Protocol {
	rankingMK8Protocol := &RankingMK8Protocol{server: server}
	rankingMK8Protocol.RankingProtocol.server = server

	rankingMK8Protocol.Setup()

	return rankingMK8Protocol
}
