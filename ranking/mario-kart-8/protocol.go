// Package protocol implements the RankingMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/ranking"
	ranking_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/ranking/mario-kart-8/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Ranking (Mario Kart 8) protocol
	ProtocolID = 0x70

	// MethodGetCompetitionRankingScore is the method ID for the GetCompetitionRankingScore method
	MethodGetCompetitionRankingScore = 0xE

	// MethodUploadCompetitionRankingScore is the method ID for the UploadCompetitionRankingScore method
	MethodUploadCompetitionRankingScore = 0xF

	// MethodGetCompetitionInfo is the method ID for the GetCompetitionInfo method
	MethodGetCompetitionInfo = 0x10
)

var patchedMethods = []uint32{
	MethodGetCompetitionRankingScore,
	MethodUploadCompetitionRankingScore,
	MethodGetCompetitionInfo,
}

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Mario Kart 8) protocol and listens for requests
// Embeds the Ranking protocol
type Protocol struct {
	Server nex.ServerInterface
	rankingProtocol
	GetCompetitionRankingScore    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	UploadCompetitionRankingScore func(err error, packet nex.PacketInterface, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam) (*nex.RMCMessage, uint32)
	GetCompetitionInfo            func(err error, packet nex.PacketInterface, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingInfoGetParam) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.rankingProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodGetCompetitionRankingScore:
		protocol.handleGetCompetitionRankingScore(packet)
	case MethodUploadCompetitionRankingScore:
		protocol.handleUploadCompetitionRankingScore(packet)
	case MethodGetCompetitionInfo:
		protocol.handleGetCompetitionInfo(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Ranking (Mario Kart 8) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new RankingMarioKart8 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.rankingProtocol.Server = server

	protocol.Setup()

	return protocol
}
