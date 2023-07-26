// Package matchmake_referee implements the Matchmake Referee NEX protocol
package matchmake_referee

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_referee_types "github.com/PretendoNetwork/nex-protocols-go/matchmake-referee/types"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x78

	// MethodStartRound is the method ID for the method StartRound
	MethodStartRound = 0x1

	// MethodGetStartRoundParam is the method ID for the method GetStartRoundParam
	MethodGetStartRoundParam = 0x2

	// MethodEndRound is the method ID for the method EndRound
	MethodEndRound = 0x3

	// MethodEndRoundWithoutReport is the method ID for the method EndRoundWithoutReport
	MethodEndRoundWithoutReport = 0x4

	// MethodGetRoundParticipants is the method ID for the method GetRoundParticipants
	MethodGetRoundParticipants = 0x5

	// MethodGetNotSummarizedRound is the method ID for the method GetNotSummarizedRound
	MethodGetNotSummarizedRound = 0x6

	// MethodGetRound is the method ID for the method GetRound
	MethodGetRound = 0x7

	// MethodGetStatsPrimary is the method ID for the method GetStatsPrimary
	MethodGetStatsPrimary = 0x8

	// MethodGetStatsPrimaries is the method ID for the method GetStatsPrimaries
	MethodGetStatsPrimaries = 0x9

	// MethodGetStatsAll is the method ID for the method GetStatsAll
	MethodGetStatsAll = 0xA

	// MethodCreateStats is the method ID for the method CreateStats
	MethodCreateStats = 0xB

	// MethodGetOrCreateStats is the method ID for the method GetOrCreateStats
	MethodGetOrCreateStats = 0xC

	// MethodResetStats is the method ID for the method ResetStats
	MethodResetStats = 0xD
)

// MatchmakeRefereeProtocol handles the Matchmake Referee NEX protocol
type MatchmakeRefereeProtocol struct {
	Server                       *nex.Server
	StartRoundHandler            func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam)
	GetStartRoundParamHandler    func(err error, client *nex.Client, callID uint32, roundID uint64)
	EndRoundHandler              func(err error, client *nex.Client, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam)
	EndRoundWithoutReportHandler func(err error, client *nex.Client, callID uint32, roundID uint64)
	GetRoundParticipantsHandler  func(err error, client *nex.Client, callID uint32, roundID uint64)
	GetNotSummarizedRoundHandler func(err error, client *nex.Client, callID uint32)
	GetRoundHandler              func(err error, client *nex.Client, callID uint32, roundID uint64)
	GetStatsPrimaryHandler       func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget)
	GetStatsPrimariesHandler     func(err error, client *nex.Client, callID uint32, targets []*matchmake_referee_types.MatchmakeRefereeStatsTarget)
	GetStatsAllHandler           func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget)
	CreateStatsHandler           func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam)
	GetOrCreateStatsHandler      func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam)
	ResetStatsHandler            func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *MatchmakeRefereeProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodStartRound:
				go protocol.handleStartRound(packet)
			case MethodGetStartRoundParam:
				go protocol.handleGetStartRoundParam(packet)
			case MethodEndRound:
				go protocol.handleEndRound(packet)
			case MethodEndRoundWithoutReport:
				go protocol.handleEndRoundWithoutReport(packet)
			case MethodGetRoundParticipants:
				go protocol.handleGetRoundParticipants(packet)
			case MethodGetNotSummarizedRound:
				go protocol.handleGetNotSummarizedRound(packet)
			case MethodGetRound:
				go protocol.handleGetRound(packet)
			case MethodGetStatsPrimary:
				go protocol.handleGetStatsPrimary(packet)
			case MethodGetStatsPrimaries:
				go protocol.handleGetStatsPrimaries(packet)
			case MethodGetStatsAll:
				go protocol.handleGetStatsAll(packet)
			case MethodCreateStats:
				go protocol.handleCreateStats(packet)
			case MethodGetOrCreateStats:
				go protocol.handleGetOrCreateStats(packet)
			case MethodResetStats:
				go protocol.handleResetStats(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported MatchmakeReferee method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewMatchmakeRefereeProtocol returns a new MatchmakeRefereeProtocol
func NewMatchmakeRefereeProtocol(server *nex.Server) *MatchmakeRefereeProtocol {
	protocol := &MatchmakeRefereeProtocol{Server: server}

	protocol.Setup()

	return protocol
}
