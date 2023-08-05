// Package protocol implements the Matchmake Referee protocol
package protocol

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

// Protocol stores all the RMC method handlers for the Matchmake Referee protocol and listens for requests
type Protocol struct {
	Server                       *nex.Server
	startRoundHandler            func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) uint32
	getStartRoundParamHandler    func(err error, client *nex.Client, callID uint32, roundID uint64) uint32
	endRoundHandler              func(err error, client *nex.Client, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) uint32
	endRoundWithoutReportHandler func(err error, client *nex.Client, callID uint32, roundID uint64) uint32
	getRoundParticipantsHandler  func(err error, client *nex.Client, callID uint32, roundID uint64) uint32
	getNotSummarizedRoundHandler func(err error, client *nex.Client, callID uint32) uint32
	getRoundHandler              func(err error, client *nex.Client, callID uint32, roundID uint64) uint32
	getStatsPrimaryHandler       func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32
	getStatsPrimariesHandler     func(err error, client *nex.Client, callID uint32, targets []*matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32
	getStatsAllHandler           func(err error, client *nex.Client, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) uint32
	createStatsHandler           func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) uint32
	getOrCreateStatsHandler      func(err error, client *nex.Client, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) uint32
	resetStatsHandler            func(err error, client *nex.Client, callID uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
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
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported MatchmakeReferee method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewProtocol returns a new Matchmake Referee protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
