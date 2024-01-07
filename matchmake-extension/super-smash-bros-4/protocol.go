// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Matchmake Extension (Super Smash Bros. 4) protocol
	ProtocolID = 0x6D

	// MethodGetTournament is the method ID for the GetTournament method
	MethodGetTournament = 0x24

	// MethodGetTournamentReplayID is the method ID for the GetTournamentReplayID method
	MethodGetTournamentReplayID = 0x25

	// MethodGetTournamentResult is the method ID for the GetTournamentResult method
	MethodGetTournamentResult = 0x26

	// MethodSetTournamentReplayID is the method ID for the SetTournamentReplayID method
	MethodSetTournamentReplayID = 0x27

	// MethodGetTournamentProfiles is the method ID for the GetTournamentProfiles method
	MethodGetTournamentProfiles = 0x28

	// MethodJoinOrCreateMatchmakeSession is the method ID for the JoinOrCreateMatchmakeSession method
	MethodJoinOrCreateMatchmakeSession = 0x29

	// MethodRegisterTournamentPlayerInfo is the method ID for the RegisterTournamentPlayerInfo method
	MethodRegisterTournamentPlayerInfo = 0x2A

	// MethodRegisterTournamentBot is the method ID for the RegisterTournamentBot method
	MethodRegisterTournamentBot = 0x2B

	// MethodReportTournamentBotRoundResult is the method ID for the ReportTournamentBotRoundResult method
	MethodReportTournamentBotRoundResult = 0x2C

	// MethodReplaceTournamentLeafNode is the method ID for the ReplaceTournamentLeafNode method
	MethodReplaceTournamentLeafNode = 0x2D

	// MethodStartTournament is the method ID for the StartTournament method
	MethodStartTournament = 0x2E

	// MethodAutoTournamentMatchmake is the method ID for the AutoTournamentMatchmake method
	MethodAutoTournamentMatchmake = 0x2F

	// MethodSimpleFindByID is the method ID for the SimpleFindByID method
	MethodSimpleFindByID = 0x30

	// MethodGetTournamentCompetitions is the method ID for the GetTournamentCompetitions method
	MethodGetTournamentCompetitions = 0x31

	// MethodGetTournamentCompetition is the method ID for the GetTournamentCompetition method
	MethodGetTournamentCompetition = 0x32

	// MethodGetTournamentReplayIDs is the method ID for the GetTournamentReplayIDs method
	MethodGetTournamentReplayIDs = 0x33

	// MethodRegisterCommunityCompetition is the method ID for the RegisterCommunityCompetition method
	MethodRegisterCommunityCompetition = 0x34

	// MethodUnregisterCommunityCompetition is the method ID for the UnregisterCommunityCompetition method
	MethodUnregisterCommunityCompetition = 0x35

	// MethodUnregisterCommunityCompetitionByID is the method ID for the UnregisterCommunityCompetitionByID method
	MethodUnregisterCommunityCompetitionByID = 0x36

	// MethodGetCommunityCompetitions is the method ID for the GetCommunityCompetitions method
	MethodGetCommunityCompetitions = 0x37

	// MethodGetCommunityCompetitionByID is the method ID for the GetCommunityCompetitionByID method
	MethodGetCommunityCompetitionByID = 0x38

	// MethodFindCommunityCompetitionsByParticipant is the method ID for the FindCommunityCompetitionsByParticipant method
	MethodFindCommunityCompetitionsByParticipant = 0x39

	// MethodFindCommunityCompetitionsByGatheringID is the method ID for the FindCommunityCompetitionsByGatheringID method
	MethodFindCommunityCompetitionsByGatheringID = 0x3A

	// MethodSelectCommunityCompetitionByOwner is the method ID for the SelectCommunityCompetitionByOwner method
	MethodSelectCommunityCompetitionByOwner = 0x3B

	// MethodJoinCommunityCompetition is the method ID for the JoinCommunityCompetition method
	MethodJoinCommunityCompetition = 0x3C

	// MethodJoinCommunityCompetitionByGatheringID is the method ID for the JoinCommunityCompetitionByGatheringID method
	MethodJoinCommunityCompetitionByGatheringID = 0x3D

	// MethodEndCommunityCompetitionParticipation is the method ID for the EndCommunityCompetitionParticipation method
	MethodEndCommunityCompetitionParticipation = 0x3E

	// MethodEndCommunityCompetitionParticipationByGatheringID is the method ID for the EndCommunityCompetitionParticipationByGatheringID method
	MethodEndCommunityCompetitionParticipationByGatheringID = 0x3F

	// MethodSearchCommunityCompetition is the method ID for the SearchCommunityCompetition method
	MethodSearchCommunityCompetition = 0x40

	// MethodPostCommunityCompetitionMatchResult is the method ID for the PostCommunityCompetitionMatchResult method
	MethodPostCommunityCompetitionMatchResult = 0x41

	// MethodGetCommunityCompetitionRanking is the method ID for the GetCommunityCompetitionRanking method
	MethodGetCommunityCompetitionRanking = 0x42

	// MethodDebugRegisterCommunityCompetition is the method ID for the DebugRegisterCommunityCompetition method
	MethodDebugRegisterCommunityCompetition = 0x43

	// MethodDebugUnregisterCommunityCompetition is the method ID for the DebugUnregisterCommunityCompetition method
	MethodDebugUnregisterCommunityCompetition = 0x44

	// MethodDebugJoinCommunityCompetition is the method ID for the DebugJoinCommunityCompetition method
	MethodDebugJoinCommunityCompetition = 0x45

	// MethodDebugEndCommunityCompetitionParticipation is the method ID for the DebugEndCommunityCompetitionParticipation method
	MethodDebugEndCommunityCompetitionParticipation = 0x46

	// MethodDebugPostCommunityCompetitionMatchResult is the method ID for the DebugPostCommunityCompetitionMatchResult method
	MethodDebugPostCommunityCompetitionMatchResult = 0x47
)

var patchedMethods = []uint32{
	MethodGetTournament,
	MethodGetTournamentReplayID,
	MethodGetTournamentResult,
	MethodSetTournamentReplayID,
	MethodGetTournamentProfiles,
	MethodJoinOrCreateMatchmakeSession,
	MethodRegisterTournamentPlayerInfo,
	MethodRegisterTournamentBot,
	MethodReportTournamentBotRoundResult,
	MethodReplaceTournamentLeafNode,
	MethodStartTournament,
	MethodAutoTournamentMatchmake,
	MethodSimpleFindByID,
	MethodGetTournamentCompetitions,
	MethodGetTournamentCompetition,
	MethodGetTournamentReplayIDs,
	MethodRegisterCommunityCompetition,
	MethodUnregisterCommunityCompetition,
	MethodUnregisterCommunityCompetitionByID,
	MethodGetCommunityCompetitions,
	MethodGetCommunityCompetitionByID,
	MethodFindCommunityCompetitionsByParticipant,
	MethodFindCommunityCompetitionsByGatheringID,
	MethodSelectCommunityCompetitionByOwner,
	MethodJoinCommunityCompetition,
	MethodJoinCommunityCompetitionByGatheringID,
	MethodEndCommunityCompetitionParticipation,
	MethodEndCommunityCompetitionParticipationByGatheringID,
	MethodSearchCommunityCompetition,
	MethodPostCommunityCompetitionMatchResult,
	MethodGetCommunityCompetitionRanking,
	MethodDebugRegisterCommunityCompetition,
	MethodDebugUnregisterCommunityCompetition,
	MethodDebugJoinCommunityCompetition,
	MethodDebugEndCommunityCompetitionParticipation,
	MethodDebugPostCommunityCompetitionMatchResult,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Super Smash Bros. 4) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	server nex.ServerInterface
	matchmakeExtensionProtocol
	GetTournament                                     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentReplayID                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentResult                               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	SetTournamentReplayID                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentProfiles                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	JoinOrCreateMatchmakeSession                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	RegisterTournamentPlayerInfo                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	RegisterTournamentBot                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	ReportTournamentBotRoundResult                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	ReplaceTournamentLeafNode                         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	StartTournament                                   func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	AutoTournamentMatchmake                           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	SimpleFindByID                                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentCompetitions                         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentCompetition                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetTournamentReplayIDs                            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	RegisterCommunityCompetition                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	UnregisterCommunityCompetition                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	UnregisterCommunityCompetitionByID                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetCommunityCompetitions                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetCommunityCompetitionByID                       func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	FindCommunityCompetitionsByParticipant            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	FindCommunityCompetitionsByGatheringID            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	SelectCommunityCompetitionByOwner                 func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	JoinCommunityCompetition                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	JoinCommunityCompetitionByGatheringID             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	EndCommunityCompetitionParticipation              func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	EndCommunityCompetitionParticipationByGatheringID func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	SearchCommunityCompetition                        func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	PostCommunityCompetitionMatchResult               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	GetCommunityCompetitionRanking                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DebugRegisterCommunityCompetition                 func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DebugUnregisterCommunityCompetition               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DebugJoinCommunityCompetition                     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DebugEndCommunityCompetitionParticipation         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
	DebugPostCommunityCompetitionMatchResult          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, message.MethodID) {
				protocol.HandlePacket(packet)
			} else {
				protocol.matchmakeExtensionProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodGetTournament:
		protocol.handleGetTournament(packet)
	case MethodGetTournamentReplayID:
		protocol.handleGetTournamentReplayID(packet)
	case MethodGetTournamentResult:
		protocol.handleGetTournamentResult(packet)
	case MethodSetTournamentReplayID:
		protocol.handleSetTournamentReplayID(packet)
	case MethodGetTournamentProfiles:
		protocol.handleGetTournamentProfiles(packet)
	case MethodJoinOrCreateMatchmakeSession:
		protocol.handleJoinOrCreateMatchmakeSession(packet)
	case MethodRegisterTournamentPlayerInfo:
		protocol.handleRegisterTournamentPlayerInfo(packet)
	case MethodRegisterTournamentBot:
		protocol.handleRegisterTournamentBot(packet)
	case MethodReportTournamentBotRoundResult:
		protocol.handleReportTournamentBotRoundResult(packet)
	case MethodReplaceTournamentLeafNode:
		protocol.handleReplaceTournamentLeafNode(packet)
	case MethodStartTournament:
		protocol.handleStartTournament(packet)
	case MethodAutoTournamentMatchmake:
		protocol.handleAutoTournamentMatchmake(packet)
	case MethodSimpleFindByID:
		protocol.handleSimpleFindByID(packet)
	case MethodGetTournamentCompetitions:
		protocol.handleGetTournamentCompetitions(packet)
	case MethodGetTournamentCompetition:
		protocol.handleGetTournamentCompetition(packet)
	case MethodGetTournamentReplayIDs:
		protocol.handleGetTournamentReplayIDs(packet)
	case MethodRegisterCommunityCompetition:
		protocol.handleRegisterCommunityCompetition(packet)
	case MethodUnregisterCommunityCompetition:
		protocol.handleUnregisterCommunityCompetition(packet)
	case MethodUnregisterCommunityCompetitionByID:
		protocol.handleUnregisterCommunityCompetitionByID(packet)
	case MethodGetCommunityCompetitions:
		protocol.handleGetCommunityCompetitions(packet)
	case MethodGetCommunityCompetitionByID:
		protocol.handleGetCommunityCompetitionByID(packet)
	case MethodFindCommunityCompetitionsByParticipant:
		protocol.handleFindCommunityCompetitionsByParticipant(packet)
	case MethodFindCommunityCompetitionsByGatheringID:
		protocol.handleFindCommunityCompetitionsByGatheringID(packet)
	case MethodSelectCommunityCompetitionByOwner:
		protocol.handleSelectCommunityCompetitionByOwner(packet)
	case MethodJoinCommunityCompetition:
		protocol.handleJoinCommunityCompetition(packet)
	case MethodJoinCommunityCompetitionByGatheringID:
		protocol.handleJoinCommunityCompetitionByGatheringID(packet)
	case MethodEndCommunityCompetitionParticipation:
		protocol.handleEndCommunityCompetitionParticipation(packet)
	case MethodEndCommunityCompetitionParticipationByGatheringID:
		protocol.handleEndCommunityCompetitionParticipationByGatheringID(packet)
	case MethodSearchCommunityCompetition:
		protocol.handleSearchCommunityCompetition(packet)
	case MethodPostCommunityCompetitionMatchResult:
		protocol.handlePostCommunityCompetitionMatchResult(packet)
	case MethodGetCommunityCompetitionRanking:
		protocol.handleGetCommunityCompetitionRanking(packet)
	case MethodDebugRegisterCommunityCompetition:
		protocol.handleDebugRegisterCommunityCompetition(packet)
	case MethodDebugUnregisterCommunityCompetition:
		protocol.handleDebugUnregisterCommunityCompetition(packet)
	case MethodDebugJoinCommunityCompetition:
		protocol.handleDebugJoinCommunityCompetition(packet)
	case MethodDebugEndCommunityCompetitionParticipation:
		protocol.handleDebugEndCommunityCompetitionParticipation(packet)
	case MethodDebugPostCommunityCompetitionMatchResult:
		protocol.handleDebugPostCommunityCompetitionMatchResult(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension (Super Smash Bros. 4) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new MatchmakeExtensionSuperSmashBros4 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}
	protocol.matchmakeExtensionProtocol.SetServer(server)

	protocol.Setup()

	return protocol
}
