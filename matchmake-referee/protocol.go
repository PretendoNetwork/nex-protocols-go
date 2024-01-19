// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server                nex.ServerInterface
	StartRound            func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, uint32)
	GetStartRoundParam    func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	EndRound              func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, uint32)
	EndRoundWithoutReport func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetRoundParticipants  func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetNotSummarizedRound func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetRound              func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	GetStatsPrimary       func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)
	GetStatsPrimaries     func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, uint32)
	GetStatsAll           func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)
	CreateStats           func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)
	GetOrCreateStats      func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)
	ResetStats            func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Matchmake Referee protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerStartRound(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, uint32))
	SetHandlerGetStartRoundParam(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerEndRound(handler func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, uint32))
	SetHandlerEndRoundWithoutReport(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetRoundParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetNotSummarizedRound(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetRound(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerGetStatsPrimary(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32))
	SetHandlerGetStatsPrimaries(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, uint32))
	SetHandlerGetStatsAll(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32))
	SetHandlerCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32))
	SetHandlerGetOrCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32))
	SetHandlerResetStats(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerStartRound sets the handler for the StartRound method
func (protocol *Protocol) SetHandlerStartRound(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStartRoundParam) (*nex.RMCMessage, uint32)) {
	protocol.StartRound = handler
}

// SetHandlerGetStartRoundParam sets the handler for the GetStartRoundParam method
func (protocol *Protocol) SetHandlerGetStartRoundParam(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetStartRoundParam = handler
}

// SetHandlerEndRound sets the handler for the EndRound method
func (protocol *Protocol) SetHandlerEndRound(handler func(err error, packet nex.PacketInterface, callID uint32, endRoundParam *matchmake_referee_types.MatchmakeRefereeEndRoundParam) (*nex.RMCMessage, uint32)) {
	protocol.EndRound = handler
}

// SetHandlerEndRoundWithoutReport sets the handler for the EndRoundWithoutReport method
func (protocol *Protocol) SetHandlerEndRoundWithoutReport(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.EndRoundWithoutReport = handler
}

// SetHandlerGetRoundParticipants sets the handler for the GetRoundParticipants method
func (protocol *Protocol) SetHandlerGetRoundParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetRoundParticipants = handler
}

// SetHandlerGetNotSummarizedRound sets the handler for the GetNotSummarizedRound method
func (protocol *Protocol) SetHandlerGetNotSummarizedRound(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetNotSummarizedRound = handler
}

// SetHandlerGetRound sets the handler for the GetRound method
func (protocol *Protocol) SetHandlerGetRound(handler func(err error, packet nex.PacketInterface, callID uint32, roundID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.GetRound = handler
}

// SetHandlerGetStatsPrimary sets the handler for the GetStatsPrimary method
func (protocol *Protocol) SetHandlerGetStatsPrimary(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)) {
	protocol.GetStatsPrimary = handler
}

// SetHandlerGetStatsPrimaries sets the handler for the GetStatsPrimaries method
func (protocol *Protocol) SetHandlerGetStatsPrimaries(handler func(err error, packet nex.PacketInterface, callID uint32, targets *types.List[*matchmake_referee_types.MatchmakeRefereeStatsTarget]) (*nex.RMCMessage, uint32)) {
	protocol.GetStatsPrimaries = handler
}

// SetHandlerGetStatsAll sets the handler for the GetStatsAll method
func (protocol *Protocol) SetHandlerGetStatsAll(handler func(err error, packet nex.PacketInterface, callID uint32, target *matchmake_referee_types.MatchmakeRefereeStatsTarget) (*nex.RMCMessage, uint32)) {
	protocol.GetStatsAll = handler
}

// SetHandlerCreateStats sets the handler for the CreateStats method
func (protocol *Protocol) SetHandlerCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)) {
	protocol.CreateStats = handler
}

// SetHandlerGetOrCreateStats sets the handler for the GetOrCreateStats method
func (protocol *Protocol) SetHandlerGetOrCreateStats(handler func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_referee_types.MatchmakeRefereeStatsInitParam) (*nex.RMCMessage, uint32)) {
	protocol.GetOrCreateStats = handler
}

// SetHandlerResetStats sets the handler for the ResetStats method
func (protocol *Protocol) SetHandlerResetStats(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.ResetStats = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
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
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported MatchmakeReferee method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new Matchmake Referee protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
