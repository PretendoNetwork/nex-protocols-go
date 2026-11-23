// Package protocol implements the GameSession protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	game_session_types "github.com/PretendoNetwork/nex-protocols-go/v2/game-session/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the GameSession protocol
	ProtocolID = 0x42

	// MethodCreateSession is the method ID for the method CreateSession
	MethodCreateSession = 0x1

	// MethodUpdateSession is the method ID for the method UpdateSession
	MethodUpdateSession = 0x2

	// MethodDeleteSession is the method ID for the method DeleteSession
	MethodDeleteSession = 0x3

	// MethodMigrateSession is the method ID for the method MigrateSession
	MethodMigrateSession = 0x4

	// MethodLeaveSession is the method ID for the method LeaveSession
	MethodLeaveSession = 0x5

	// MethodGetSession is the method ID for the method GetSession
	MethodGetSession = 0x6

	// MethodSearchSessions is the method ID for the method SearchSessions
	MethodSearchSessions = 0x7

	// MethodAddParticipants is the method ID for the method AddParticipants
	MethodAddParticipants = 0x8

	// MethodRemoveParticipants is the method ID for the method RemoveParticipants
	MethodRemoveParticipants = 0x9

	// MethodGetParticipantCount is the method ID for the method GetParticipantCount
	MethodGetParticipantCount = 0xa

	// MethodGetParticipants is the method ID for the method GetParticipants
	MethodGetParticipants = 0xb

	// MethodSendInvitation is the method ID for the method SendInvitation
	MethodSendInvitation = 0xc

	// MethodGetInvitationReceivedCount is the method ID for the method GetInvitationReceivedCount
	MethodGetInvitationReceivedCount = 0xd

	// MethodGetInvitationsReceived is the method ID for the method GetInvitationsReceived
	MethodGetInvitationsReceived = 0xe

	// MethodGetInvitationSentCount is the method ID for the method GetInvitationSentCount
	MethodGetInvitationSentCount = 0xf

	// MethodGetInvitationsSent is the method ID for the method GetInvitationsSent
	MethodGetInvitationsSent = 0x10

	// MethodAcceptInvitation is the method ID for the method AcceptInvitation
	MethodAcceptInvitation = 0x11

	// MethodDeclineInvitation is the method ID for the method DeclineInvitation
	MethodDeclineInvitation = 0x12

	// MethodCancelInvitation is the method ID for the method CancelInvitation
	MethodCancelInvitation = 0x13

	// MethodSendTextMessage is the method ID for the method SendTextMessage
	MethodSendTextMessage = 0x14

	// MethodRegisterURLs is the method ID for the method RegisterURLs
	MethodRegisterURLs = 0x15

	// MethodJoinSession is the method ID for the method JoinSession
	MethodJoinSession = 0x16

	// MethodAbandonSession is the method ID for the method AbandonSession
	MethodAbandonSession = 0x17

	// MethodSearchSessionsWithParticipants is the method ID for the method SearchSessionsWithParticipants
	MethodSearchSessionsWithParticipants = 0x18

	// MethodGetSessions is the method ID for the method GetSessions
	MethodGetSessions = 0x19

	// MethodGetParticipantsURLs is the method ID for the method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x1a

	// MethodMigrateSessionHost is the method ID for the method MigrateSessionHost
	MethodMigrateSessionHost = 0x1b

	// MethodSplitSession is the method ID for the method SplitSession
	MethodSplitSession = 0x1c

	// MethodSearchSocialSessions is the method ID for the method SearchSocialSessions
	MethodSearchSocialSessions = 0x1d

	// MethodReportUnsuccessfulJoinSessions is the method ID for the method ReportUnsuccessfulJoinSessions
	MethodReportUnsuccessfulJoinSessions = 0x1e
)

// Protocol handles the GameSession protocol and listens for requests
type Protocol struct {
	endpoint                       nex.EndpointInterface
	CreateSession                  func(err error, packet nex.PacketInterface, callID uint32, gameSession game_session_types.GameSession) (*nex.RMCMessage, *nex.Error)
	UpdateSession                  func(err error, packet nex.PacketInterface, callID uint32, gameSessionUpdate game_session_types.GameSessionUpdate) (*nex.RMCMessage, *nex.Error)
	DeleteSession                  func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	MigrateSession                 func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	LeaveSession                   func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	GetSession                     func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	SearchSessions                 func(err error, packet nex.PacketInterface, callID uint32, gameSessionQuery game_session_types.GameSessionQuery) (*nex.RMCMessage, *nex.Error)
	AddParticipants                func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, publicParticipantIDs types.List[types.PID], privateParticipantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	RemoveParticipants             func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	GetParticipantCount            func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	GetParticipants                func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	SendInvitation                 func(err error, packet nex.PacketInterface, callID uint32, invitation game_session_types.GameSessionInvitation) (*nex.RMCMessage, *nex.Error)
	GetInvitationReceivedCount     func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetInvitationsReceived         func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	GetInvitationSentCount         func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error)
	GetInvitationsSent             func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)
	AcceptInvitation               func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error)
	DeclineInvitation              func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error)
	CancelInvitation               func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationSent) (*nex.RMCMessage, *nex.Error)
	SendTextMessage                func(err error, packet nex.PacketInterface, callID uint32, gameSessionMessage game_session_types.GameSessionMessage) (*nex.RMCMessage, *nex.Error)
	RegisterURLs                   func(err error, packet nex.PacketInterface, callID uint32, stationURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error)
	JoinSession                    func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	AbandonSession                 func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	SearchSessionsWithParticipants func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	GetSessions                    func(err error, packet nex.PacketInterface, callID uint32, gameSessionKeys types.List[game_session_types.GameSessionKey]) (*nex.RMCMessage, *nex.Error)
	GetParticipantsURLs            func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)
	MigrateSessionHost             func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	SplitSession                   func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)
	SearchSocialSessions           func(err error, packet nex.PacketInterface, callID uint32, gameSessionSocialQuery game_session_types.GameSessionSocialQuery) (*nex.RMCMessage, *nex.Error)
	ReportUnsuccessfulJoinSessions func(err error, packet nex.PacketInterface, callID uint32, unsuccessfulJoinSessions types.List[game_session_types.GameSessionUnsuccessfulJoinSession]) (*nex.RMCMessage, *nex.Error)
	Patches                        nex.ServiceProtocol
	PatchedMethods                 []uint32
}

// Interface implements the methods present on the GameSession protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerCreateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSession game_session_types.GameSession) (*nex.RMCMessage, *nex.Error))
	SetHandlerUpdateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionUpdate game_session_types.GameSessionUpdate) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerMigrateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerLeaveSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionQuery game_session_types.GameSessionQuery) (*nex.RMCMessage, *nex.Error))
	SetHandlerAddParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, publicParticipantIDs types.List[types.PID], privateParticipantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipantCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, invitation game_session_types.GameSessionInvitation) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationReceivedCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationsReceived(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationSentCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetInvitationsSent(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerAcceptInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeclineInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error))
	SetHandlerCancelInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationSent) (*nex.RMCMessage, *nex.Error))
	SetHandlerSendTextMessage(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionMessage game_session_types.GameSessionMessage) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegisterURLs(handler func(err error, packet nex.PacketInterface, callID uint32, stationURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerJoinSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerAbandonSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchSessionsWithParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKeys types.List[game_session_types.GameSessionKey]) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error))
	SetHandlerMigrateSessionHost(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSplitSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error))
	SetHandlerSearchSocialSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionSocialQuery game_session_types.GameSessionSocialQuery) (*nex.RMCMessage, *nex.Error))
	SetHandlerReportUnsuccessfulJoinSessions(handler func(err error, packet nex.PacketInterface, callID uint32, unsuccessfulJoinSessions types.List[game_session_types.GameSessionUnsuccessfulJoinSession]) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerCreateSession sets the handler for the CreateSession method
func (protocol *Protocol) SetHandlerCreateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSession game_session_types.GameSession) (*nex.RMCMessage, *nex.Error)) {
	protocol.CreateSession = handler
}

// SetHandlerUpdateSession sets the handler for the UpdateSession method
func (protocol *Protocol) SetHandlerUpdateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionUpdate game_session_types.GameSessionUpdate) (*nex.RMCMessage, *nex.Error)) {
	protocol.UpdateSession = handler
}

// SetHandlerDeleteSession sets the handler for the DeleteSession method
func (protocol *Protocol) SetHandlerDeleteSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteSession = handler
}

// SetHandlerMigrateSession sets the handler for the MigrateSession method
func (protocol *Protocol) SetHandlerMigrateSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.MigrateSession = handler
}

// SetHandlerLeaveSession sets the handler for the LeaveSession method
func (protocol *Protocol) SetHandlerLeaveSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.LeaveSession = handler
}

// SetHandlerGetSession sets the handler for the GetSession method
func (protocol *Protocol) SetHandlerGetSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSession = handler
}

// SetHandlerSearchSessions sets the handler for the SearchSessions method
func (protocol *Protocol) SetHandlerSearchSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionQuery game_session_types.GameSessionQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchSessions = handler
}

// SetHandlerAddParticipants sets the handler for the AddParticipations method
func (protocol *Protocol) SetHandlerAddParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, publicParticipantIDs types.List[types.PID], privateParticipantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.AddParticipants = handler
}

// SetHandlerRemoveParticipants sets the handler for the RemoveParticipants method
func (protocol *Protocol) SetHandlerRemoveParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveParticipants = handler
}

// SetHandlerGetParticipantCount sets the handler for the GetParticipantCount method
func (protocol *Protocol) SetHandlerGetParticipantCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipantCount = handler
}

// SetHandlerGetParticipants sets the handler for the GetParticipants method
func (protocol *Protocol) SetHandlerGetParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipants = handler
}

// SetHandlerSendInvitation sets the handler for the SendInvitation method
func (protocol *Protocol) SetHandlerSendInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, invitation game_session_types.GameSessionInvitation) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendInvitation = handler
}

// SetHandlerGetInvitationReceivedCount sets the handler for the GetInvitationReceivedCount method
func (protocol *Protocol) SetHandlerGetInvitationReceivedCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationReceivedCount = handler
}

// SetHandlerGetInvitationsReceived sets the handler for the GetInvitationsReceived method
func (protocol *Protocol) SetHandlerGetInvitationsReceived(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationsReceived = handler
}

// SetHandlerGetInvitationSentCount sets the handler for the GetInvitationSentCount method
func (protocol *Protocol) SetHandlerGetInvitationSentCount(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationSentCount = handler
}

// SetHandlerGetInvitationsSent sets the handler for the GetInvitationsSent method
func (protocol *Protocol) SetHandlerGetInvitationsSent(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, resultRange types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetInvitationsSent = handler
}

// SetHandlerAcceptInvitation sets the handler for the AcceptInvitation method
func (protocol *Protocol) SetHandlerAcceptInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error)) {
	protocol.AcceptInvitation = handler
}

// SetHandlerDeclineInvitation sets the handler for the DeclineInvitation method
func (protocol *Protocol) SetHandlerDeclineInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationReceived) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeclineInvitation = handler
}

// SetHandlerCancelInvitation sets the handler for the CancelInvitation method
func (protocol *Protocol) SetHandlerCancelInvitation(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionInvitation game_session_types.GameSessionInvitationSent) (*nex.RMCMessage, *nex.Error)) {
	protocol.CancelInvitation = handler
}

// SetHandlerSendTextMessage sets the handler for the SendTextMessage method
func (protocol *Protocol) SetHandlerSendTextMessage(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionMessage game_session_types.GameSessionMessage) (*nex.RMCMessage, *nex.Error)) {
	protocol.SendTextMessage = handler
}

// SetHandlerRegisterURLs sets the handler for the RegisterURLs method
func (protocol *Protocol) SetHandlerRegisterURLs(handler func(err error, packet nex.PacketInterface, callID uint32, stationURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterURLs = handler
}

// SetHandlerJoinSession sets the handler for the JoinSession method
func (protocol *Protocol) SetHandlerJoinSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.JoinSession = handler
}

// SetHandlerAbandonSession sets the handler for the AbandonSession method
func (protocol *Protocol) SetHandlerAbandonSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.AbandonSession = handler
}

// SetHandlerSearchSessionsWithParticipants sets the handler for the SearchSessionsWithParticipants method
func (protocol *Protocol) SetHandlerSearchSessionsWithParticipants(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionTypeID types.UInt32, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchSessionsWithParticipants = handler
}

// SetHandlerGetSessions sets the handler for the GetSessions method
func (protocol *Protocol) SetHandlerGetSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKeys types.List[game_session_types.GameSessionKey]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetSessions = handler
}

// SetHandlerGetParticipantsURLs sets the handler for the GetParticipantsURLs method
func (protocol *Protocol) SetHandlerGetParticipantsURLs(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey, participantIDs types.List[types.PID]) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetParticipantsURLs = handler
}

// SetHandlerMigrateSessionHost sets the handler for the MigrateSessionHost method
func (protocol *Protocol) SetHandlerMigrateSessionHost(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.MigrateSessionHost = handler
}

// SetHandlerSplitSession sets the handler for the SplitSession method
func (protocol *Protocol) SetHandlerSplitSession(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionKey game_session_types.GameSessionKey) (*nex.RMCMessage, *nex.Error)) {
	protocol.SplitSession = handler
}

// SetHandlerSearchSocialSessions sets the handler for the SearchSocialSessions method
func (protocol *Protocol) SetHandlerSearchSocialSessions(handler func(err error, packet nex.PacketInterface, callID uint32, gameSessionSocialQuery game_session_types.GameSessionSocialQuery) (*nex.RMCMessage, *nex.Error)) {
	protocol.SearchSocialSessions = handler
}

// SetHandlerReportUnsuccessfulJoinSessions sets the handler for the ReportUnsuccessfulJoinSessions method
func (protocol *Protocol) SetHandlerReportUnsuccessfulJoinSessions(handler func(err error, packet nex.PacketInterface, callID uint32, unsuccessfulJoinSessions types.List[game_session_types.GameSessionUnsuccessfulJoinSession]) (*nex.RMCMessage, *nex.Error)) {
	protocol.ReportUnsuccessfulJoinSessions = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodCreateSession:
		protocol.handleCreateSession(packet)
	case MethodUpdateSession:
		protocol.handleUpdateSession(packet)
	case MethodDeleteSession:
		protocol.handleDeleteSession(packet)
	case MethodMigrateSession:
		protocol.handleMigrateSession(packet)
	case MethodLeaveSession:
		protocol.handleLeaveSession(packet)
	case MethodGetSession:
		protocol.handleGetSession(packet)
	case MethodSearchSessions:
		protocol.handleSearchSessions(packet)
	case MethodAddParticipants:
		protocol.handleAddParticipants(packet)
	case MethodRemoveParticipants:
		protocol.handleRemoveParticipants(packet)
	case MethodGetParticipantCount:
		protocol.handleGetParticipantCount(packet)
	case MethodGetParticipants:
		protocol.handleGetParticipants(packet)
	case MethodSendInvitation:
		protocol.handleSendInvitation(packet)
	case MethodGetInvitationReceivedCount:
		protocol.handleGetInvitationReceivedCount(packet)
	case MethodGetInvitationsReceived:
		protocol.handleGetInvitationsReceived(packet)
	case MethodGetInvitationSentCount:
		protocol.handleGetInvitationSentCount(packet)
	case MethodGetInvitationsSent:
		protocol.handleGetInvitationsSent(packet)
	case MethodAcceptInvitation:
		protocol.handleAcceptInvitation(packet)
	case MethodDeclineInvitation:
		protocol.handleDeclineInvitation(packet)
	case MethodCancelInvitation:
		protocol.handleCancelInvitation(packet)
	case MethodSendTextMessage:
		protocol.handleSendTextMessage(packet)
	case MethodRegisterURLs:
		protocol.handleRegisterURLs(packet)
	case MethodJoinSession:
		protocol.handleJoinSession(packet)
	case MethodAbandonSession:
		protocol.handleAbandonSession(packet)
	case MethodSearchSessionsWithParticipants:
		protocol.handleSearchSessionsWithParticipants(packet)
	case MethodGetSessions:
		protocol.handleGetSessions(packet)
	case MethodGetParticipantsURLs:
		protocol.handleGetParticipantsURLs(packet)
	case MethodMigrateSessionHost:
		protocol.handleMigrateSessionHost(packet)
	case MethodSplitSession:
		protocol.handleSplitSession(packet)
	case MethodSearchSocialSessions:
		protocol.handleSearchSocialSessions(packet)
	case MethodReportUnsuccessfulJoinSessions:
		protocol.handleReportUnsuccessfulJoinSessions(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported GameSession method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new GameSession protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
