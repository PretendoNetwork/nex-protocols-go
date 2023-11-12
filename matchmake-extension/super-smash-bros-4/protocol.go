// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
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
	Server nex.ServerInterface
	matchmakeExtensionProtocol
	getTournamentHandler                                     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentReplayIDHandler                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentResultHandler                               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	setTournamentReplayIDHandler                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentProfilesHandler                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	joinOrCreateMatchmakeSessionHandler                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	registerTournamentPlayerInfoHandler                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	registerTournamentBotHandler                             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	reportTournamentBotRoundResultHandler                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	replaceTournamentLeafNodeHandler                         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	startTournamentHandler                                   func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	autoTournamentMatchmakeHandler                           func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	simpleFindByIDHandler                                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentCompetitionsHandler                         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentCompetitionHandler                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getTournamentReplayIDsHandler                            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	registerCommunityCompetitionHandler                      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	unregisterCommunityCompetitionHandler                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	unregisterCommunityCompetitionByIDHandler                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getCommunityCompetitionsHandler                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getCommunityCompetitionByIDHandler                       func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	findCommunityCompetitionsByParticipantHandler            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	findCommunityCompetitionsByGatheringIDHandler            func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	selectCommunityCompetitionByOwnerHandler                 func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	joinCommunityCompetitionHandler                          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	joinCommunityCompetitionByGatheringIDHandler             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	endCommunityCompetitionParticipationHandler              func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	endCommunityCompetitionParticipationByGatheringIDHandler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	searchCommunityCompetitionHandler                        func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	postCommunityCompetitionMatchResultHandler               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	getCommunityCompetitionRankingHandler                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	debugRegisterCommunityCompetitionHandler                 func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	debugUnregisterCommunityCompetitionHandler               func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	debugJoinCommunityCompetitionHandler                     func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	debugEndCommunityCompetitionParticipationHandler         func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
	debugPostCommunityCompetitionMatchResultHandler          func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			if slices.Contains(patchedMethods, request.MethodID) {
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
		go protocol.handleGetTournament(packet)
	case MethodGetTournamentReplayID:
		go protocol.handleGetTournamentReplayID(packet)
	case MethodGetTournamentResult:
		go protocol.handleGetTournamentResult(packet)
	case MethodSetTournamentReplayID:
		go protocol.handleSetTournamentReplayID(packet)
	case MethodGetTournamentProfiles:
		go protocol.handleGetTournamentProfiles(packet)
	case MethodJoinOrCreateMatchmakeSession:
		go protocol.handleJoinOrCreateMatchmakeSession(packet)
	case MethodRegisterTournamentPlayerInfo:
		go protocol.handleRegisterTournamentPlayerInfo(packet)
	case MethodRegisterTournamentBot:
		go protocol.handleRegisterTournamentBot(packet)
	case MethodReportTournamentBotRoundResult:
		go protocol.handleReportTournamentBotRoundResult(packet)
	case MethodReplaceTournamentLeafNode:
		go protocol.handleReplaceTournamentLeafNode(packet)
	case MethodStartTournament:
		go protocol.handleStartTournament(packet)
	case MethodAutoTournamentMatchmake:
		go protocol.handleAutoTournamentMatchmake(packet)
	case MethodSimpleFindByID:
		go protocol.handleSimpleFindByID(packet)
	case MethodGetTournamentCompetitions:
		go protocol.handleGetTournamentCompetitions(packet)
	case MethodGetTournamentCompetition:
		go protocol.handleGetTournamentCompetition(packet)
	case MethodGetTournamentReplayIDs:
		go protocol.handleGetTournamentReplayIDs(packet)
	case MethodRegisterCommunityCompetition:
		go protocol.handleRegisterCommunityCompetition(packet)
	case MethodUnregisterCommunityCompetition:
		go protocol.handleUnregisterCommunityCompetition(packet)
	case MethodUnregisterCommunityCompetitionByID:
		go protocol.handleUnregisterCommunityCompetitionByID(packet)
	case MethodGetCommunityCompetitions:
		go protocol.handleGetCommunityCompetitions(packet)
	case MethodGetCommunityCompetitionByID:
		go protocol.handleGetCommunityCompetitionByID(packet)
	case MethodFindCommunityCompetitionsByParticipant:
		go protocol.handleFindCommunityCompetitionsByParticipant(packet)
	case MethodFindCommunityCompetitionsByGatheringID:
		go protocol.handleFindCommunityCompetitionsByGatheringID(packet)
	case MethodSelectCommunityCompetitionByOwner:
		go protocol.handleSelectCommunityCompetitionByOwner(packet)
	case MethodJoinCommunityCompetition:
		go protocol.handleJoinCommunityCompetition(packet)
	case MethodJoinCommunityCompetitionByGatheringID:
		go protocol.handleJoinCommunityCompetitionByGatheringID(packet)
	case MethodEndCommunityCompetitionParticipation:
		go protocol.handleEndCommunityCompetitionParticipation(packet)
	case MethodEndCommunityCompetitionParticipationByGatheringID:
		go protocol.handleEndCommunityCompetitionParticipationByGatheringID(packet)
	case MethodSearchCommunityCompetition:
		go protocol.handleSearchCommunityCompetition(packet)
	case MethodPostCommunityCompetitionMatchResult:
		go protocol.handlePostCommunityCompetitionMatchResult(packet)
	case MethodGetCommunityCompetitionRanking:
		go protocol.handleGetCommunityCompetitionRanking(packet)
	case MethodDebugRegisterCommunityCompetition:
		go protocol.handleDebugRegisterCommunityCompetition(packet)
	case MethodDebugUnregisterCommunityCompetition:
		go protocol.handleDebugUnregisterCommunityCompetition(packet)
	case MethodDebugJoinCommunityCompetition:
		go protocol.handleDebugJoinCommunityCompetition(packet)
	case MethodDebugEndCommunityCompetitionParticipation:
		go protocol.handleDebugEndCommunityCompetitionParticipation(packet)
	case MethodDebugPostCommunityCompetitionMatchResult:
		go protocol.handleDebugPostCommunityCompetitionMatchResult(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension (Super Smash Bros. 4) method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new MatchmakeExtensionSuperSmashBros4 protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.matchmakeExtensionProtocol.Server = server

	protocol.Setup()

	return protocol
}
