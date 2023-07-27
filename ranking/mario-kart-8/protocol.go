// Package protocol implements the Mario Kart 8 Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/ranking"
)

const (
	// ProtocolID is the protocol ID for the Ranking (Mario Kart 8) protocol. ID is the same as the Ranking protocol
	ProtocolID = 0x70
)

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Mario Kart 8) protocol and listens for requests
// Embeds the Ranking Protocol
type Protocol struct {
	Server *nex.Server
	rankingProtocol
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
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

// NewProtocol returns a new Ranking (Mario Kart 8)
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.rankingProtocol.Server = server

	protocol.Setup()

	return protocol
}
