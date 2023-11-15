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
	Server                      nex.ServerInterface
	RegisterGathering           func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) (*nex.RMCMessage, uint32)
	UnregisterGathering         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	UnregisterGatherings        func(err error, packet nex.PacketInterface, callID uint32, lstGatherings []uint32) (*nex.RMCMessage, uint32)
	UpdateGathering             func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) (*nex.RMCMessage, uint32)
	Invite                      func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []*nex.PID, strMessage string) (*nex.RMCMessage, uint32)
	AcceptInvitation            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) (*nex.RMCMessage, uint32)
	DeclineInvitation           func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) (*nex.RMCMessage, uint32)
	CancelInvitation            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []*nex.PID, strMessage string) (*nex.RMCMessage, uint32)
	GetInvitationsSent          func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	GetInvitationsReceived      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	Participate                 func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) (*nex.RMCMessage, uint32)
	CancelParticipation         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) (*nex.RMCMessage, uint32)
	GetParticipants             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	AddParticipants             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstPrincipals []*nex.PID, strMessage string) (*nex.RMCMessage, uint32)
	GetDetailedParticipants     func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	GetParticipantsURLs         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	FindByType                  func(err error, packet nex.PacketInterface, callID uint32, strType string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	FindByDescription           func(err error, packet nex.PacketInterface, callID uint32, strDescription string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	FindByDescriptionRegex      func(err error, packet nex.PacketInterface, callID uint32, strDescriptionRegex string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	FindByID                    func(err error, packet nex.PacketInterface, callID uint32, lstID []uint32) (*nex.RMCMessage, uint32)
	FindBySingleID              func(err error, packet nex.PacketInterface, callID uint32, id uint32) (*nex.RMCMessage, uint32)
	FindByOwner                 func(err error, packet nex.PacketInterface, callID uint32, id *nex.PID, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	FindByParticipants          func(err error, packet nex.PacketInterface, callID uint32, pid []*nex.PID) (*nex.RMCMessage, uint32)
	FindInvitations             func(err error, packet nex.PacketInterface, callID uint32, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	FindBySQLQuery              func(err error, packet nex.PacketInterface, callID uint32, strQuery string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	LaunchSession               func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strURL string) (*nex.RMCMessage, uint32)
	UpdateSessionURL            func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strURL string) (*nex.RMCMessage, uint32)
	GetSessionURL               func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	GetState                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	SetState                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, uiNewState uint32) (*nex.RMCMessage, uint32)
	ReportStats                 func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstStats []*match_making_types.GatheringStats) (*nex.RMCMessage, uint32)
	GetStats                    func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, lstParticipants []*nex.PID, lstColumns []byte) (*nex.RMCMessage, uint32)
	DeleteGathering             func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) (*nex.RMCMessage, uint32)
	GetPendingDeletions         func(err error, packet nex.PacketInterface, callID uint32, uiReason uint32, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	DeleteFromDeletions         func(err error, packet nex.PacketInterface, callID uint32, lstDeletions []uint32) (*nex.RMCMessage, uint32)
	MigrateGatheringOwnershipV1 func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []*nex.PID) (*nex.RMCMessage, uint32)
	FindByDescriptionLike       func(err error, packet nex.PacketInterface, callID uint32, strDescriptionLike string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	RegisterLocalURL            func(err error, packet nex.PacketInterface, callID uint32, gid uint32, url *nex.StationURL) (*nex.RMCMessage, uint32)
	RegisterLocalURLs           func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstURLs []*nex.StationURL) (*nex.RMCMessage, uint32)
	UpdateSessionHostV1         func(err error, packet nex.PacketInterface, callID uint32, gid uint32) (*nex.RMCMessage, uint32)
	GetSessionURLs              func(err error, packet nex.PacketInterface, callID uint32, gid uint32) (*nex.RMCMessage, uint32)
	UpdateSessionHost           func(err error, packet nex.PacketInterface, callID uint32, gid uint32, isMigrateOwner bool) (*nex.RMCMessage, uint32)
	UpdateGatheringOwnership    func(err error, packet nex.PacketInterface, callID uint32, gid uint32, participantsOnly bool) (*nex.RMCMessage, uint32)
	MigrateGatheringOwnership   func(err error, packet nex.PacketInterface, callID uint32, gid uint32, lstPotentialNewOwnersID []*nex.PID, participantsOnly bool) (*nex.RMCMessage, uint32)
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
		protocol.handleRegisterGathering(packet)
	case MethodUnregisterGathering:
		protocol.handleUnregisterGathering(packet)
	case MethodUnregisterGatherings:
		protocol.handleUnregisterGatherings(packet)
	case MethodUpdateGathering:
		protocol.handleUpdateGathering(packet)
	case MethodInvite:
		protocol.handleInvite(packet)
	case MethodAcceptInvitation:
		protocol.handleAcceptInvitation(packet)
	case MethodDeclineInvitation:
		protocol.handleDeclineInvitation(packet)
	case MethodCancelInvitation:
		protocol.handleCancelInvitation(packet)
	case MethodGetInvitationsSent:
		protocol.handleGetInvitationsSent(packet)
	case MethodGetInvitationsReceived:
		protocol.handleGetInvitationsReceived(packet)
	case MethodParticipate:
		protocol.handleParticipate(packet)
	case MethodCancelParticipation:
		protocol.handleCancelParticipation(packet)
	case MethodGetParticipants:
		protocol.handleGetParticipants(packet)
	case MethodAddParticipants:
		protocol.handleAddParticipants(packet)
	case MethodGetDetailedParticipants:
		protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		protocol.handleGetParticipantsURLs(packet)
	case MethodFindByType:
		protocol.handleFindByType(packet)
	case MethodFindByDescription:
		protocol.handleFindByDescription(packet)
	case MethodFindByDescriptionRegex:
		protocol.handleFindByDescriptionRegex(packet)
	case MethodFindByID:
		protocol.handleFindByID(packet)
	case MethodFindBySingleID:
		protocol.handleFindBySingleID(packet)
	case MethodFindByOwner:
		protocol.handleFindByOwner(packet)
	case MethodFindByParticipants:
		protocol.handleFindByParticipants(packet)
	case MethodFindInvitations:
		protocol.handleFindInvitations(packet)
	case MethodFindBySQLQuery:
		protocol.handleFindBySQLQuery(packet)
	case MethodLaunchSession:
		protocol.handleLaunchSession(packet)
	case MethodUpdateSessionURL:
		protocol.handleUpdateSessionURL(packet)
	case MethodGetSessionURL:
		protocol.handleGetSessionURL(packet)
	case MethodGetState:
		protocol.handleGetState(packet)
	case MethodSetState:
		protocol.handleSetState(packet)
	case MethodReportStats:
		protocol.handleReportStats(packet)
	case MethodGetStats:
		protocol.handleGetStats(packet)
	case MethodDeleteGathering:
		protocol.handleDeleteGathering(packet)
	case MethodGetPendingDeletions:
		protocol.handleGetPendingDeletions(packet)
	case MethodDeleteFromDeletions:
		protocol.handleDeleteFromDeletions(packet)
	case MethodMigrateGatheringOwnershipV1:
		protocol.handleMigrateGatheringOwnershipV1(packet)
	case MethodFindByDescriptionLike:
		protocol.handleFindByDescriptionLike(packet)
	case MethodRegisterLocalURL:
		protocol.handleRegisterLocalURL(packet)
	case MethodRegisterLocalURLs:
		protocol.handleRegisterLocalURLs(packet)
	case MethodUpdateSessionHostV1:
		protocol.handleUpdateSessionHostV1(packet)
	case MethodGetSessionURLs:
		protocol.handleGetSessionURLs(packet)
	case MethodUpdateSessionHost:
		protocol.handleUpdateSessionHost(packet)
	case MethodUpdateGatheringOwnership:
		protocol.handleUpdateGatheringOwnership(packet)
	case MethodMigrateGatheringOwnership:
		protocol.handleMigrateGatheringOwnership(packet)
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
