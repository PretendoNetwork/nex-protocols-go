// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

const (
	// ProtocolID is the protocol ID for the Matchmake Extension protocol
	ProtocolID = 0x6D

	// MethodCloseParticipation is the method ID for method CloseParticipation
	MethodCloseParticipation = 0x1

	// MethodOpenParticipation is the method ID for method OpenParticipation
	MethodOpenParticipation = 0x2

	// MethodAutoMatchmakePostpone is the method ID for method AutoMatchmakePostpone
	MethodAutoMatchmakePostpone = 0x3

	// MethodBrowseMatchmakeSession is the method ID for method BrowseMatchmakeSession
	MethodBrowseMatchmakeSession = 0x4

	// MethodBrowseMatchmakeSessionWithHostURLs is the method ID for method BrowseMatchmakeSessionWithHostURLs
	MethodBrowseMatchmakeSessionWithHostURLs = 0x5

	// MethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MethodCreateMatchmakeSession = 0x6

	// MethodJoinMatchmakeSession is the method ID for method JoinMatchmakeSession
	MethodJoinMatchmakeSession = 0x7

	// MethodModifyCurrentGameAttribute is the method ID for method ModifyCurrentGameAttribute
	MethodModifyCurrentGameAttribute = 0x8

	// MethodUpdateNotificationData is the method ID for method UpdateNotificationData
	MethodUpdateNotificationData = 0x9

	// MethodGetFriendNotificationData is the method ID for method GetFriendNotificationData
	MethodGetFriendNotificationData = 0xA

	// MethodUpdateApplicationBuffer is the method ID for method UpdateApplicationBuffer
	MethodUpdateApplicationBuffer = 0xB

	// MethodUpdateMatchmakeSessionAttribute is the method ID for method UpdateMatchmakeSessionAttribute
	MethodUpdateMatchmakeSessionAttribute = 0xC

	// MethodGetlstFriendNotificationData is the method ID for method GetlstFriendNotificationData
	MethodGetlstFriendNotificationData = 0xD

	// MethodUpdateMatchmakeSession is the method ID for method UpdateMatchmakeSession
	MethodUpdateMatchmakeSession = 0xE

	// MethodAutoMatchmakeWithSearchCriteriaPostpone is the method ID for method AutoMatchmakeWithSearchCriteriaPostpone
	MethodAutoMatchmakeWithSearchCriteriaPostpone = 0xF

	// MethodGetPlayingSession is the method ID for method GetPlayingSession
	MethodGetPlayingSession = 0x10

	// MethodCreateCommunity is the method ID for method CreateCommunity
	MethodCreateCommunity = 0x11

	// MethodUpdateCommunity is the method ID for method UpdateCommunity
	MethodUpdateCommunity = 0x12

	// MethodJoinCommunity is the method ID for method JoinCommunity
	MethodJoinCommunity = 0x13

	// MethodFindCommunityByGatheringID is the method ID for method FindCommunityByGatheringID
	MethodFindCommunityByGatheringID = 0x14

	// MethodFindOfficialCommunity is the method ID for method FindOfficialCommunity
	MethodFindOfficialCommunity = 0x15

	// MethodFindCommunityByParticipant is the method ID for method FindCommunityByParticipant
	MethodFindCommunityByParticipant = 0x16

	// MethodUpdatePrivacySetting is the method ID for method UpdatePrivacySetting
	MethodUpdatePrivacySetting = 0x17

	// MethodGetMyBlockList is the method ID for method GetMyBlockList
	MethodGetMyBlockList = 0x18

	// MethodAddToBlockList is the method ID for method AddToBlockList
	MethodAddToBlockList = 0x19

	// MethodRemoveFromBlockList is the method ID for method RemoveFromBlockList
	MethodRemoveFromBlockList = 0x1A

	// MethodClearMyBlockList is the method ID for method ClearMyBlockList
	MethodClearMyBlockList = 0x1B

	// MethodReportViolation is the method ID for method ReportViolation
	MethodReportViolation = 0x1C

	// MethodIsViolationUser is the method ID for method IsViolationUser
	MethodIsViolationUser = 0x1D

	// MethodJoinMatchmakeSessionEx is the method ID for method JoinMatchmakeSessionEx
	MethodJoinMatchmakeSessionEx = 0x1E

	// MethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MethodGetSimplePlayingSession = 0x1F

	// MethodGetSimpleCommunity is the method ID for method GetSimpleCommunity
	MethodGetSimpleCommunity = 0x20

	// MethodAutoMatchmakeWithGatheringIDPostpone is the method ID for method AutoMatchmakeWithGatheringIDPostpone
	MethodAutoMatchmakeWithGatheringIDPostpone = 0x21

	// MethodUpdateProgressScore is the method ID for method UpdateProgressScore
	MethodUpdateProgressScore = 0x22

	// MethodDebugNotifyEvent is the method ID for method DebugNotifyEvent
	MethodDebugNotifyEvent = 0x23

	// MethodGenerateMatchmakeSessionSystemPassword is the method ID for method GenerateMatchmakeSessionSystemPassword
	MethodGenerateMatchmakeSessionSystemPassword = 0x24

	// MethodClearMatchmakeSessionSystemPassword is the method ID for method ClearMatchmakeSessionSystemPassword
	MethodClearMatchmakeSessionSystemPassword = 0x25

	// MethodCreateMatchmakeSessionWithParam is the method ID for method CreateMatchmakeSessionWithParam
	MethodCreateMatchmakeSessionWithParam = 0x26

	// MethodJoinMatchmakeSessionWithParam is the method ID for method JoinMatchmakeSessionWithParam
	MethodJoinMatchmakeSessionWithParam = 0x27

	// MethodAutoMatchmakeWithParamPostpone is the method ID for method AutoMatchmakeWithParamPostpone
	MethodAutoMatchmakeWithParamPostpone = 0x28

	// MethodFindMatchmakeSessionByGatheringIDDetail is the method ID for method FindMatchmakeSessionByGatheringIDDetail
	MethodFindMatchmakeSessionByGatheringIDDetail = 0x29

	// MethodBrowseMatchmakeSessionNoHolder is the method ID for method BrowseMatchmakeSessionNoHolder
	MethodBrowseMatchmakeSessionNoHolder = 0x2A

	// MethodBrowseMatchmakeSessionWithHostURLsNoHolder is the method ID for method BrowseMatchmakeSessionWithHostURLsNoHolder
	MethodBrowseMatchmakeSessionWithHostURLsNoHolder = 0x2B

	// MethodUpdateMatchmakeSessionPart is the method ID for method UpdateMatchmakeSessionPart
	MethodUpdateMatchmakeSessionPart = 0x2C

	// MethodRequestMatchmaking is the method ID for method RequestMatchmaking
	MethodRequestMatchmaking = 0x2D

	// MethodWithdrawMatchmaking is the method ID for method WithdrawMatchmaking
	MethodWithdrawMatchmaking = 0x2E

	// MethodWithdrawMatchmakingAll is the method ID for method WithdrawMatchmakingAll
	MethodWithdrawMatchmakingAll = 0x2F

	// MethodFindMatchmakeSessionByGatheringID is the method ID for method FindMatchmakeSessionByGatheringID
	MethodFindMatchmakeSessionByGatheringID = 0x30

	// MethodFindMatchmakeSessionBySingleGatheringID is the method ID for method FindMatchmakeSessionBySingleGatheringID
	MethodFindMatchmakeSessionBySingleGatheringID = 0x31

	// MethodFindMatchmakeSessionByOwner is the method ID for method FindMatchmakeSessionByOwner
	MethodFindMatchmakeSessionByOwner = 0x32

	// MethodFindMatchmakeSessionByParticipant is the method ID for method FindMatchmakeSessionByParticipant
	MethodFindMatchmakeSessionByParticipant = 0x33

	// MethodBrowseMatchmakeSessionNoHolderNoResultRange is the method ID for method BrowseMatchmakeSessionNoHolderNoResultRange
	MethodBrowseMatchmakeSessionNoHolderNoResultRange = 0x34

	// MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange is the method ID for method BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange
	MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange = 0x35

	// MethodFindCommunityByOwner is the method ID for method FindCommunityByOwner
	MethodFindCommunityByOwner = 0x36
)

// Protocol stores all the RMC method handlers for the Matchmake Extension protocol and listens for requests
type Protocol struct {
	Server                                                  nex.ServerInterface
	CloseParticipation                                      func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	OpenParticipation                                       func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	AutoMatchmakePostpone                                   func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder, message string) uint32
	BrowseMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32
	BrowseMatchmakeSessionWithHostURLs                      func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32
	CreateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder, message string, participationCount uint16) uint32
	JoinMatchmakeSession                                    func(err error, packet nex.PacketInterface, callID uint32, gid uint32, strMessage string) uint32
	ModifyCurrentGameAttribute                              func(err error, packet nex.PacketInterface, callID uint32, gid uint32, attribIndex uint32, newValue uint32) uint32
	UpdateNotificationData                                  func(err error, packet nex.PacketInterface, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string) uint32
	GetFriendNotificationData                               func(err error, packet nex.PacketInterface, callID uint32, uiType int32) uint32
	UpdateApplicationBuffer                                 func(err error, packet nex.PacketInterface, callID uint32, gid uint32, applicationBuffer []byte) uint32
	UpdateMatchmakeSessionAttribute                         func(err error, packet nex.PacketInterface, callID uint32, gid uint32, attribs []uint32) uint32
	GetlstFriendNotificationData                            func(err error, packet nex.PacketInterface, callID uint32, lstTypes []uint32) uint32
	UpdateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32
	AutoMatchmakeWithSearchCriteriaPostpone                 func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *nex.DataHolder, strMessage string) uint32
	GetPlayingSession                                       func(err error, packet nex.PacketInterface, callID uint32, lstPID []*nex.PID) uint32
	CreateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage string) uint32
	UpdateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) uint32
	JoinCommunity                                           func(err error, packet nex.PacketInterface, callID uint32, gid uint32, strMessage string, strPassword string) uint32
	FindCommunityByGatheringID                              func(err error, packet nex.PacketInterface, callID uint32, lstGID []uint32) uint32
	FindOfficialCommunity                                   func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange) uint32
	FindCommunityByParticipant                              func(err error, packet nex.PacketInterface, callID uint32, pid *nex.PID, resultRange *nex.ResultRange) uint32
	UpdatePrivacySetting                                    func(err error, packet nex.PacketInterface, callID uint32, onlineStatus bool, participationCommunity bool) uint32
	GetMyBlockList                                          func(err error, packet nex.PacketInterface, callID uint32) uint32
	AddToBlockList                                          func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID []*nex.PID) uint32
	RemoveFromBlockList                                     func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID []*nex.PID) uint32
	ClearMyBlockList                                        func(err error, packet nex.PacketInterface, callID uint32) uint32
	ReportViolation                                         func(err error, packet nex.PacketInterface, callID uint32, pid *nex.PID, userName string, violationCode uint32) uint32
	IsViolationUser                                         func(err error, packet nex.PacketInterface, callID uint32) uint32
	JoinMatchmakeSessionEx                                  func(err error, packet nex.PacketInterface, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16) uint32
	GetSimplePlayingSession                                 func(err error, packet nex.PacketInterface, callID uint32, listPID []*nex.PID, includeLoginUser bool) uint32
	GetSimpleCommunity                                      func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList []uint32) uint32
	AutoMatchmakeWithGatheringIDPostpone                    func(err error, packet nex.PacketInterface, callID uint32, lstGID []uint32, anyGathering *nex.DataHolder, strMessage string) uint32
	UpdateProgressScore                                     func(err error, packet nex.PacketInterface, callID uint32, gid uint32, progressScore uint8) uint32
	DebugNotifyEvent                                        func(err error, packet nex.PacketInterface, callID uint32, pid *nex.PID, mainType uint32, subType uint32, param1 uint64, param2 uint64, stringParam string) uint32
	GenerateMatchmakeSessionSystemPassword                  func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	ClearMatchmakeSessionSystemPassword                     func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	CreateMatchmakeSessionWithParam                         func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) uint32
	JoinMatchmakeSessionWithParam                           func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) uint32
	AutoMatchmakeWithParamPostpone                          func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) uint32
	FindMatchmakeSessionByGatheringIDDetail                 func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	BrowseMatchmakeSessionNoHolder                          func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32
	BrowseMatchmakeSessionWithHostURLsNoHolder              func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *nex.ResultRange) uint32
	UpdateMatchmakeSessionPart                              func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) uint32
	RequestMatchmaking                                      func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) uint32
	WithdrawMatchmaking                                     func(err error, packet nex.PacketInterface, callID uint32, requestID uint64) uint32
	WithdrawMatchmakingAll                                  func(err error, packet nex.PacketInterface, callID uint32) uint32
	FindMatchmakeSessionByGatheringID                       func(err error, packet nex.PacketInterface, callID uint32, lstGID []uint32) uint32
	FindMatchmakeSessionBySingleGatheringID                 func(err error, packet nex.PacketInterface, callID uint32, gid uint32) uint32
	FindMatchmakeSessionByOwner                             func(err error, packet nex.PacketInterface, callID uint32, id uint32, resultRange *nex.ResultRange) uint32
	FindMatchmakeSessionByParticipant                       func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) uint32
	BrowseMatchmakeSessionNoHolderNoResultRange             func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) uint32
	BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) uint32
	FindCommunityByOwner                                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32 // TODO - Unknown request/response format
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
	case MethodCloseParticipation:
		go protocol.handleCloseParticipation(packet)
	case MethodOpenParticipation:
		go protocol.handleOpenParticipation(packet)
	case MethodAutoMatchmakePostpone:
		go protocol.handleAutoMatchmakePostpone(packet)
	case MethodBrowseMatchmakeSession:
		go protocol.handleBrowseMatchmakeSession(packet)
	case MethodBrowseMatchmakeSessionWithHostURLs:
		go protocol.handleBrowseMatchmakeSessionWithHostURLs(packet)
	case MethodCreateMatchmakeSession:
		go protocol.handleCreateMatchmakeSession(packet)
	case MethodJoinMatchmakeSession:
		go protocol.handleJoinMatchmakeSession(packet)
	case MethodModifyCurrentGameAttribute:
		go protocol.handleModifyCurrentGameAttribute(packet)
	case MethodUpdateNotificationData:
		go protocol.handleUpdateNotificationData(packet)
	case MethodGetFriendNotificationData:
		go protocol.handleGetFriendNotificationData(packet)
	case MethodUpdateApplicationBuffer:
		go protocol.handleUpdateApplicationBuffer(packet)
	case MethodUpdateMatchmakeSessionAttribute:
		go protocol.handleUpdateMatchmakeSessionAttribute(packet)
	case MethodGetlstFriendNotificationData:
		go protocol.handleGetlstFriendNotificationData(packet)
	case MethodUpdateMatchmakeSession:
		go protocol.handleUpdateMatchmakeSession(packet)
	case MethodAutoMatchmakeWithSearchCriteriaPostpone:
		go protocol.handleAutoMatchmakeWithSearchCriteriaPostpone(packet)
	case MethodGetPlayingSession:
		go protocol.handleGetPlayingSession(packet)
	case MethodCreateCommunity:
		go protocol.handleCreateCommunity(packet)
	case MethodUpdateCommunity:
		go protocol.handleUpdateCommunity(packet)
	case MethodJoinCommunity:
		go protocol.handleJoinCommunity(packet)
	case MethodFindCommunityByGatheringID:
		go protocol.handleFindCommunityByGatheringID(packet)
	case MethodFindOfficialCommunity:
		go protocol.handleFindOfficialCommunity(packet)
	case MethodFindCommunityByParticipant:
		go protocol.handleFindCommunityByParticipant(packet)
	case MethodUpdatePrivacySetting:
		go protocol.handleUpdatePrivacySetting(packet)
	case MethodGetMyBlockList:
		go protocol.handleGetMyBlockList(packet)
	case MethodAddToBlockList:
		go protocol.handleAddToBlockList(packet)
	case MethodRemoveFromBlockList:
		go protocol.handleRemoveFromBlockList(packet)
	case MethodClearMyBlockList:
		go protocol.handleClearMyBlockList(packet)
	case MethodReportViolation:
		go protocol.handleReportViolation(packet)
	case MethodIsViolationUser:
		go protocol.handleIsViolationUser(packet)
	case MethodJoinMatchmakeSessionEx:
		go protocol.handleJoinMatchmakeSessionEx(packet)
	case MethodGetSimplePlayingSession:
		go protocol.handleGetSimplePlayingSession(packet)
	case MethodGetSimpleCommunity:
		go protocol.handleGetSimpleCommunity(packet)
	case MethodAutoMatchmakeWithGatheringIDPostpone:
		go protocol.handleAutoMatchmakeWithGatheringIDPostpone(packet)
	case MethodUpdateProgressScore:
		go protocol.handleUpdateProgressScore(packet)
	case MethodDebugNotifyEvent:
		go protocol.handleDebugNotifyEvent(packet)
	case MethodGenerateMatchmakeSessionSystemPassword:
		go protocol.handleGenerateMatchmakeSessionSystemPassword(packet)
	case MethodClearMatchmakeSessionSystemPassword:
		go protocol.handleClearMatchmakeSessionSystemPassword(packet)
	case MethodCreateMatchmakeSessionWithParam:
		go protocol.handleCreateMatchmakeSessionWithParam(packet)
	case MethodJoinMatchmakeSessionWithParam:
		go protocol.handleJoinMatchmakeSessionWithParam(packet)
	case MethodAutoMatchmakeWithParamPostpone:
		go protocol.handleAutoMatchmakeWithParamPostpone(packet)
	case MethodFindMatchmakeSessionByGatheringIDDetail:
		go protocol.handleFindMatchmakeSessionByGatheringIDDetail(packet)
	case MethodBrowseMatchmakeSessionNoHolder:
		go protocol.handleBrowseMatchmakeSessionNoHolder(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolder:
		go protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolder(packet)
	case MethodUpdateMatchmakeSessionPart:
		go protocol.handleUpdateMatchmakeSessionPart(packet)
	case MethodRequestMatchmaking:
		go protocol.handleRequestMatchmaking(packet)
	case MethodWithdrawMatchmaking:
		go protocol.handleWithdrawMatchmaking(packet)
	case MethodWithdrawMatchmakingAll:
		go protocol.handleWithdrawMatchmakingAll(packet)
	case MethodFindMatchmakeSessionByGatheringID:
		go protocol.handleFindMatchmakeSessionByGatheringID(packet)
	case MethodFindMatchmakeSessionBySingleGatheringID:
		go protocol.handleFindMatchmakeSessionBySingleGatheringID(packet)
	case MethodFindMatchmakeSessionByOwner:
		go protocol.handleFindMatchmakeSessionByOwner(packet)
	case MethodFindMatchmakeSessionByParticipant:
		go protocol.handleFindMatchmakeSessionByParticipant(packet)
	case MethodBrowseMatchmakeSessionNoHolderNoResultRange:
		go protocol.handleBrowseMatchmakeSessionNoHolderNoResultRange(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange:
		go protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet)
	case MethodFindCommunityByOwner:
		go protocol.handleFindCommunityByOwner(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Matchmake Extension protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
