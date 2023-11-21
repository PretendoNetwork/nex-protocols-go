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
	Server                nex.ServerInterface
	StartRound            func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, uint32)
	GetStartRoundParam    func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) (*nex.RMCMessage, uint32)
	EndRound              func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, uint32)
	EndRoundWithoutReport func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) (*nex.RMCMessage, uint32)
	GetRoundParticipants  func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) (*nex.RMCMessage, uint32)
	GetNotSummarizedRound func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetRound              func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) (*nex.RMCMessage, uint32)
	GetStatsPrimary       func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)
	GetStatsPrimaries     func(err error, packet nex.PacketInterface, callID uint32, targets []*matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)
	GetStatsAll           func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)
	CreateStats           func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)
	GetOrCreateStats      func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)
	ResetStats            func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodStartRound:
				protocol.handleStartRound(packet)
			case MethodGetStartRoundParam:
				protocol.handleGetStartRoundParam(packet)
			case MethodEndRound:
				protocol.handleEndRound(packet)
			case MethodEndRoundWithoutReport:
				protocol.handleEndRoundWithoutReport(packet)
			case MethodGetRoundParticipants:
				protocol.handleGetRoundParticipants(packet)
			case MethodGetNotSummarizedRound:
				protocol.handleGetNotSummarizedRound(packet)
			case MethodGetRound:
				protocol.handleGetRound(packet)
			case MethodGetStatsPrimary:
				protocol.handleGetStatsPrimary(packet)
			case MethodGetStatsPrimaries:
				protocol.handleGetStatsPrimaries(packet)
			case MethodGetStatsAll:
				protocol.handleGetStatsAll(packet)
			case MethodCreateStats:
				protocol.handleCreateStats(packet)
			case MethodGetOrCreateStats:
				protocol.handleGetOrCreateStats(packet)
			case MethodResetStats:
				protocol.handleResetStats(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported MatchmakeReferee method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Matchmake Referee protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
