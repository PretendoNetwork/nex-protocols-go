// Package protocol implements the RankingSplatoon protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking "github.com/PretendoNetwork/nex-protocols-go/ranking"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Ranking (Splatoon) protocol
	ProtocolID = 0x70

	// MethodGetCompetitionRankingScore is the method ID for the GetCompetitionRankingScore method
	MethodGetCompetitionRankingScore = 0x10

	// MethodGetcompetitionRankingScoreByPeriodList is the method ID for the GetcompetitionRankingScoreByPeriodList method
	MethodGetcompetitionRankingScoreByPeriodList = 0x11

	// MethodUploadCompetitionRankingScore is the method ID for the UploadCompetitionRankingScore method
	MethodUploadCompetitionRankingScore = 0x12

	// MethodDeleteCompetitionRankingScore is the method ID for the DeleteCompetitionRankingScore method
	MethodDeleteCompetitionRankingScore = 0x13
)

var patchedMethods = []uint32{
	MethodGetCompetitionRankingScore,
	MethodGetcompetitionRankingScoreByPeriodList,
	MethodUploadCompetitionRankingScore,
	MethodDeleteCompetitionRankingScore,
}

type rankingProtocol = ranking.Protocol

// Protocol stores all the RMC method handlers for the Ranking (Splatoon) protocol and listens for requests
// Embeds the Ranking protocol
type Protocol struct {
	Server nex.ServerInterface
	rankingProtocol
	GetCompetitionRankingScore             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetcompetitionRankingScoreByPeriodList func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	UploadCompetitionRankingScore          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DeleteCompetitionRankingScore          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
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
	case MethodGetcompetitionRankingScoreByPeriodList:
		protocol.handleGetcompetitionRankingScoreByPeriodList(packet)
	case MethodUploadCompetitionRankingScore:
		protocol.handleUploadCompetitionRankingScore(packet)
	case MethodDeleteCompetitionRankingScore:
		protocol.handleDeleteCompetitionRankingScore(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Ranking (Splatoon) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new RankingSplatoon protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.rankingProtocol.Server = server

	protocol.Setup()

	return protocol
}
