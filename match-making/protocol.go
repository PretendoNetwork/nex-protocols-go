// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

const (
	// ProtocolID is the protocol ID for the Match Making protocol
	ProtocolID = 0x15

	// MethodRegisterGathering is the method ID for method RegisterGathering
	MethodRegisterGathering = 0x1

	// MethodUnregisterGathering is the method ID for method UnregisterGathering
	MethodUnregisterGathering = 0x2

	// MethodUnregisterGatherings is the method ID for method UnregisterGatherings
	MethodUnregisterGatherings = 0x3

	// MethodUpdateGathering is the method ID for method UpdateGathering
	MethodUpdateGathering = 0x4

	// MethodInvite is the method ID for method Invite
	MethodInvite = 0x5

	// MethodAcceptInvitation is the method ID for method AcceptInvitation
	MethodAcceptInvitation = 0x6

	// MethodDeclineInvitation is the method ID for method DeclineInvitation
	MethodDeclineInvitation = 0x7

	// MethodCancelInvitation is the method ID for method CancelInvitation
	MethodCancelInvitation = 0x8

	// MethodGetInvitationsSent is the method ID for method GetInvitationsSent
	MethodGetInvitationsSent = 0x9

	// MethodGetInvitationsReceived is the method ID for method GetInvitationsReceived
	MethodGetInvitationsReceived = 0xA

	// MethodParticipate is the method ID for method Participate
	MethodParticipate = 0xB

	// MethodCancelParticipation is the method ID for method CancelParticipation
	MethodCancelParticipation = 0xC

	// MethodGetParticipants is the method ID for method GetParticipants
	MethodGetParticipants = 0xD

	// MethodAddParticipants is the method ID for method AddParticipants
	MethodAddParticipants = 0xE

	// MethodGetDetailedParticipants is the method ID for method GetDetailedParticipants
	MethodGetDetailedParticipants = 0xF

	// MethodGetParticipantsURLs is the method ID for method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x10

	// MethodFindByType is the method ID for method FindByType
	MethodFindByType = 0x11

	// MethodFindByDescription is the method ID for method FindByDescription
	MethodFindByDescription = 0x12

	// MethodFindByDescriptionRegex is the method ID for method FindByDescriptionRegex
	MethodFindByDescriptionRegex = 0x13

	// MethodFindByID is the method ID for method FindByID
	MethodFindByID = 0x14

	// MethodFindBySingleID is the method ID for method FindBySingleID
	MethodFindBySingleID = 0x15

	// MethodFindByOwner is the method ID for method FindByOwner
	MethodFindByOwner = 0x16

	// MethodFindByParticipants is the method ID for method FindByParticipants
	MethodFindByParticipants = 0x17

	// MethodFindInvitations is the method ID for method FindInvitations
	MethodFindInvitations = 0x18

	// MethodFindBySQLQuery is the method ID for method FindBySQLQuery
	MethodFindBySQLQuery = 0x19

	// MethodLaunchSession is the method ID for method LaunchSession
	MethodLaunchSession = 0x1A

	// MethodUpdateSessionURL is the method ID for method UpdateSessionURL
	MethodUpdateSessionURL = 0x1B

	// MethodGetSessionURL is the method ID for method GetSessionURL
	MethodGetSessionURL = 0x1C

	// MethodGetState is the method ID for method GetState
	MethodGetState = 0x1D

	// MethodSetState is the method ID for method SetState
	MethodSetState = 0x1E

	// MethodReportStats is the method ID for method ReportStats
	MethodReportStats = 0x1F

	// MethodGetStats is the method ID for method GetStats
	MethodGetStats = 0x20

	// MethodDeleteGathering is the method ID for method DeleteGathering
	MethodDeleteGathering = 0x21

	// MethodGetPendingDeletions is the method ID for method GetPendingDeletions
	MethodGetPendingDeletions = 0x22

	// MethodDeleteFromDeletions is the method ID for method DeleteFromDeletions
	MethodDeleteFromDeletions = 0x23

	// MethodMigrateGatheringOwnershipV1 is the method ID for method MigrateGatheringOwnershipV1
	MethodMigrateGatheringOwnershipV1 = 0x24

	// MethodFindByDescriptionLike is the method ID for method FindByDescriptionLike
	MethodFindByDescriptionLike = 0x25

	// MethodRegisterLocalURL is the method ID for method RegisterLocalURL
	MethodRegisterLocalURL = 0x26

	// MethodRegisterLocalURLs is the method ID for method RegisterLocalURLs
	MethodRegisterLocalURLs = 0x27

	// MethodUpdateSessionHostV1 is the method ID for method UpdateSessionHostV1
	MethodUpdateSessionHostV1 = 0x28

	// MethodGetSessionURLs is the method ID for method GetSessionURLs
	MethodGetSessionURLs = 0x29

	// MethodUpdateSessionHost is the method ID for method UpdateSessionHost
	MethodUpdateSessionHost = 0x2A

	// MethodUpdateGatheringOwnership is the method ID for method UpdateGatheringOwnership
	MethodUpdateGatheringOwnership = 0x2B

	// MethodMigrateGatheringOwnership is the method ID for method MigrateGatheringOwnership
	MethodMigrateGatheringOwnership = 0x2C
)

// Protocol stores all the RMC method handlers for the MatchMaking protocol and listens for requests
type Protocol struct {
	Server                             nex.ServerInterface
	registerGatheringHandler           func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32
	unregisterGatheringHandler         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	unregisterGatheringsHandler        func(err error, packet nex.PacketInterface, callID uint32, lstGatherings []uint32) uint32
	updateGatheringHandler             func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32
	inviteHandler                      func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string) uint32
	acceptInvitationHandler            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32
	declineInvitationHandler           func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32
	cancelInvitationHandler            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string) uint32
	getInvitationsSentHandler          func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	getInvitationsReceivedHandler      func(err error, packet nex.PacketInterface, callID uint32) uint32
	participateHandler                 func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32
	cancelParticipationHandler         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) uint32
	getParticipantsHandler             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	addParticipantsHandler             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []uint32, strMessage string) uint32
	getDetailedParticipantsHandler     func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	getParticipantsURLsHandler         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	findByTypeHandler                  func(err error, packet nex.PacketInterface, callID uint32, strType string, resultRange *nex.ResultRange) uint32
	findByDescriptionHandler           func(err error, packet nex.PacketInterface, callID uint32, strDescription string, resultRange *nex.ResultRange) uint32
	findByDescriptionRegexHandler      func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex string, resultRange *nex.ResultRange) uint32
	findByIDHandler                    func(err error, packet nex.PacketInterface, callID uint32, lstID []uint32) uint32
	findBySingleIDHandler              func(err error, packet nex.PacketInterface, callID uint32, id uint32) uint32
	findByOwnerHandler                 func(err error, packet nex.PacketInterface, callID uint32, id uint32, resultRange *nex.ResultRange) uint32
	findByParticipantsHandler          func(err error, packet nex.PacketInterface, callID uint32, pid []uint32) uint32
	findInvitationsHandler             func(err error, packet nex.PacketInterface, callID uint32, resultRange *nex.ResultRange) uint32
	findBySQLQueryHandler              func(err error, packet nex.PacketInterface, callID uint32, strQuery string, resultRange *nex.ResultRange) uint32
	launchSessionHandler               func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strURL string) uint32
	updateSessionURLHandler            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strURL string) uint32
	getSessionURLHandler               func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	getStateHandler                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	setStateHandler                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, uiNewState uint32) uint32
	reportStatsHandler                 func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstStats []*match_making_types.GatheringStats) uint32
	getStatsHandler                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstParticipants []uint32, lstColumns []byte) uint32
	deleteGatheringHandler             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32
	getPendingDeletionsHandler         func(err error, packet nex.PacketInterface, callID uint32, uiReason uint32, resultRange *nex.ResultRange) uint32
	deleteFromDeletionsHandler         func(err error, packet nex.PacketInterface, callID uint32, lstDeletions []uint32) uint32
	migrateGatheringOwnershipV1Handler func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []uint32) uint32
	findByDescriptionLikeHandler       func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike string, resultRange *nex.ResultRange) uint32
	registerLocalURLHandler            func(err error, packet nex.PacketInterface, callID uint32, gid uint32, url *nex.StationURL) uint32
	registerLocalURLsHandler           func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstURLs []*nex.StationURL) uint32
	updateSessionHostV1Handler         func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	getSessionURLsHandler              func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	updateSessionHostHandler           func(err error, packet nex.PacketInterface, callID uint32, gid uint32, isMigrateOwner bool) uint32
	updateGatheringOwnershipHandler    func(err error, packet nex.PacketInterface, callID uint32, gid uint32, participantsOnly bool) uint32
	migrateGatheringOwnershipHandler   func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []uint32, participantsOnly bool) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodRegisterGathering:
		go protocol.handleRegisterGathering(packet)
	case MethodUnregisterGathering:
		go protocol.handleUnregisterGathering(packet)
	case MethodUnregisterGatherings:
		go protocol.handleUnregisterGatherings(packet)
	case MethodUpdateGathering:
		go protocol.handleUpdateGathering(packet)
	case MethodInvite:
		go protocol.handleInvite(packet)
	case MethodAcceptInvitation:
		go protocol.handleAcceptInvitation(packet)
	case MethodDeclineInvitation:
		go protocol.handleDeclineInvitation(packet)
	case MethodCancelInvitation:
		go protocol.handleCancelInvitation(packet)
	case MethodGetInvitationsSent:
		go protocol.handleGetInvitationsSent(packet)
	case MethodGetInvitationsReceived:
		go protocol.handleGetInvitationsReceived(packet)
	case MethodParticipate:
		go protocol.handleParticipate(packet)
	case MethodCancelParticipation:
		go protocol.handleCancelParticipation(packet)
	case MethodGetParticipants:
		go protocol.handleGetParticipants(packet)
	case MethodAddParticipants:
		go protocol.handleAddParticipants(packet)
	case MethodGetDetailedParticipants:
		go protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		go protocol.handleGetParticipantsURLs(packet)
	case MethodFindByType:
		go protocol.handleFindByType(packet)
	case MethodFindByDescription:
		go protocol.handleFindByDescription(packet)
	case MethodFindByDescriptionRegex:
		go protocol.handleFindByDescriptionRegex(packet)
	case MethodFindByID:
		go protocol.handleFindByID(packet)
	case MethodFindBySingleID:
		go protocol.handleFindBySingleID(packet)
	case MethodFindByOwner:
		go protocol.handleFindByOwner(packet)
	case MethodFindByParticipants:
		go protocol.handleFindByParticipants(packet)
	case MethodFindInvitations:
		go protocol.handleFindInvitations(packet)
	case MethodFindBySQLQuery:
		go protocol.handleFindBySQLQuery(packet)
	case MethodLaunchSession:
		go protocol.handleLaunchSession(packet)
	case MethodUpdateSessionURL:
		go protocol.handleUpdateSessionURL(packet)
	case MethodGetSessionURL:
		go protocol.handleGetSessionURL(packet)
	case MethodGetState:
		go protocol.handleGetState(packet)
	case MethodSetState:
		go protocol.handleSetState(packet)
	case MethodReportStats:
		go protocol.handleReportStats(packet)
	case MethodGetStats:
		go protocol.handleGetStats(packet)
	case MethodDeleteGathering:
		go protocol.handleDeleteGathering(packet)
	case MethodGetPendingDeletions:
		go protocol.handleGetPendingDeletions(packet)
	case MethodDeleteFromDeletions:
		go protocol.handleDeleteFromDeletions(packet)
	case MethodMigrateGatheringOwnershipV1:
		go protocol.handleMigrateGatheringOwnershipV1(packet)
	case MethodFindByDescriptionLike:
		go protocol.handleFindByDescriptionLike(packet)
	case MethodRegisterLocalURL:
		go protocol.handleRegisterLocalURL(packet)
	case MethodRegisterLocalURLs:
		go protocol.handleRegisterLocalURLs(packet)
	case MethodUpdateSessionHostV1:
		go protocol.handleUpdateSessionHostV1(packet)
	case MethodGetSessionURLs:
		go protocol.handleGetSessionURLs(packet)
	case MethodUpdateSessionHost:
		go protocol.handleUpdateSessionHost(packet)
	case MethodUpdateGatheringOwnership:
		go protocol.handleUpdateGatheringOwnership(packet)
	case MethodMigrateGatheringOwnership:
		go protocol.handleMigrateGatheringOwnership(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported MatchMaking method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Match Making protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
