// Package ranking_mario_kart_8 implements the Mario Kart 8 Ranking NEX protocol
package ranking_mario_kart_8

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"github.com/PretendoNetwork/nex-protocols-go/ranking"
)

const (
	// ProtocolID is the protocol ID for the Ranking (Mario Kart 8) protocol. ID is the same as the Ranking protocol
	ProtocolID = 0x70
)

// RankingMK8Protocol handles the Ranking (Mario Kart 8) NEX protocol. Embeds RankingProtocol
type RankingMK8Protocol struct {
	Server *nex.Server
	ranking.RankingProtocol
}

// Setup initializes the protocol
func (protocol *RankingMK8Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported Ranking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewRankingMK8Protocol returns a new RankingMK8Protocol
func NewRankingMK8Protocol(server *nex.Server) *RankingMK8Protocol {
	protocol := &RankingMK8Protocol{Server: server}
	protocol.RankingProtocol.Server = server

	protocol.Setup()

	return protocol
}
