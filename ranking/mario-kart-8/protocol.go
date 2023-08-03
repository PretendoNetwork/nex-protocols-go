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
	Server *nex.Server
	rankingProtocol
	getCompetitionRankingScoreHandler    func(err error, client *nex.Client, callID uint32, packetPayload []byte)
	uploadCompetitionRankingScoreHandler func(err error, client *nex.Client, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingUploadScoreParam)
	getCompetitionInfoHandler            func(err error, client *nex.Client, callID uint32, param *ranking_mario_kart8_types.CompetitionRankingInfoGetParam)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.rankingProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetCompetitionRankingScore:
		go protocol.handleGetCompetitionRankingScore(packet)
	case MethodUploadCompetitionRankingScore:
		go protocol.handleUploadCompetitionRankingScore(packet)
	case MethodGetCompetitionInfo:
		go protocol.handleGetCompetitionInfo(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Ranking (Mario Kart 8) method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new RankingMarioKart8 protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.rankingProtocol.Server = server

	protocol.Setup()

	return protocol
}
