// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server                                                  nex.ServerInterface
	CloseParticipation                                      func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	OpenParticipation                                       func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	AutoMatchmakePostpone                                   func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSessionWithHostURLs                      func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	CreateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32)
	JoinMatchmakeSession                                    func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string) (*nex.RMCMessage, uint32)
	ModifyCurrentGameAttribute                              func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	UpdateNotificationData                                  func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam string) (*nex.RMCMessage, uint32)
	GetFriendNotificationData                               func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, uint32)
	UpdateApplicationBuffer                                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer []byte) (*nex.RMCMessage, uint32)
	UpdateMatchmakeSessionAttribute                         func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	GetlstFriendNotificationData                            func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	UpdateMatchmakeSession                                  func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, uint32)
	AutoMatchmakeWithSearchCriteriaPostpone                 func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32)
	GetPlayingSession                                       func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	CreateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage string) (*nex.RMCMessage, uint32)
	UpdateCommunity                                         func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, uint32)
	JoinCommunity                                           func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, strPassword string) (*nex.RMCMessage, uint32)
	FindCommunityByGatheringID                              func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	FindOfficialCommunity                                   func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	FindCommunityByParticipant                              func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	UpdatePrivacySetting                                    func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetMyBlockList                                          func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	AddToBlockList                                          func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	RemoveFromBlockList                                     func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32)
	ClearMyBlockList                                        func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	ReportViolation                                         func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName string, violationCode *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	IsViolationUser                                         func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	JoinMatchmakeSessionEx                                  func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32)
	GetSimplePlayingSession                                 func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, uint32)
	GetSimpleCommunity                                      func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	AutoMatchmakeWithGatheringIDPostpone                    func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32)
	UpdateProgressScore                                     func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, uint32)
	DebugNotifyEvent                                        func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam string) (*nex.RMCMessage, uint32)
	GenerateMatchmakeSessionSystemPassword                  func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	ClearMatchmakeSessionSystemPassword                     func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	CreateMatchmakeSessionWithParam                         func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, uint32)
	JoinMatchmakeSessionWithParam                           func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, uint32)
	AutoMatchmakeWithParamPostpone                          func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32)
	FindMatchmakeSessionByGatheringIDDetail                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSessionNoHolder                          func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSessionWithHostURLsNoHolder              func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	UpdateMatchmakeSessionPart                              func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, uint32)
	RequestMatchmaking                                      func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32)
	WithdrawMatchmaking                                     func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, uint32)
	WithdrawMatchmakingAll                                  func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	FindMatchmakeSessionByGatheringID                       func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)
	FindMatchmakeSessionBySingleGatheringID                 func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	FindMatchmakeSessionByOwner                             func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)
	FindMatchmakeSessionByParticipant                       func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSessionNoHolderNoResultRange             func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32)
	BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32)
	FindCommunityByOwner                                    func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Interface implements the methods present on the Matchmake Extension protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerCloseParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerOpenParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerAutoMatchmakePostpone(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSessionWithHostURLs(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerCreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32))
	SetHandlerJoinMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string) (*nex.RMCMessage, uint32))
	SetHandlerModifyCurrentGameAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam string) (*nex.RMCMessage, uint32))
	SetHandlerGetFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateApplicationBuffer(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer []byte) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMatchmakeSessionAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerGetlstFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, uint32))
	SetHandlerAutoMatchmakeWithSearchCriteriaPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32))
	SetHandlerGetPlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerCreateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, uint32))
	SetHandlerJoinCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, strPassword string) (*nex.RMCMessage, uint32))
	SetHandlerFindCommunityByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerFindOfficialCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerFindCommunityByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerUpdatePrivacySetting(handler func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerAddToBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerRemoveFromBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32))
	SetHandlerClearMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerReportViolation(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName string, violationCode *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerIsViolationUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerJoinMatchmakeSessionEx(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32))
	SetHandlerGetSimplePlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, uint32))
	SetHandlerGetSimpleCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerAutoMatchmakeWithGatheringIDPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateProgressScore(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, uint32))
	SetHandlerDebugNotifyEvent(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam string) (*nex.RMCMessage, uint32))
	SetHandlerGenerateMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerClearMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerCreateMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, uint32))
	SetHandlerJoinMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, uint32))
	SetHandlerAutoMatchmakeWithParamPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32))
	SetHandlerFindMatchmakeSessionByGatheringIDDetail(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSessionNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerUpdateMatchmakeSessionPart(handler func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, uint32))
	SetHandlerRequestMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32))
	SetHandlerWithdrawMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, uint32))
	SetHandlerWithdrawMatchmakingAll(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerFindMatchmakeSessionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32))
	SetHandlerFindMatchmakeSessionBySingleGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerFindMatchmakeSessionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerFindMatchmakeSessionByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32))
	SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32))
	SetHandlerFindCommunityByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerCloseParticipation sets the handler for the CloseParticipation method
func (protocol *Protocol) SetHandlerCloseParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.CloseParticipation = handler
}

// SetHandlerOpenParticipation sets the handler for the OpenParticipation method
func (protocol *Protocol) SetHandlerOpenParticipation(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.OpenParticipation = handler
}

// SetHandlerAutoMatchmakePostpone sets the handler for the AutoMatchmakePostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakePostpone(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string) (*nex.RMCMessage, uint32)) {
	protocol.AutoMatchmakePostpone = handler
}

// SetHandlerBrowseMatchmakeSession sets the handler for the BrowseMatchmakeSession method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSession = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLs sets the handler for the BrowseMatchmakeSessionWithHostURLs method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLs(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSessionWithHostURLs = handler
}

// SetHandlerCreateMatchmakeSession sets the handler for the CreateMatchmakeSession method
func (protocol *Protocol) SetHandlerCreateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder, message string, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32)) {
	protocol.CreateMatchmakeSession = handler
}

// SetHandlerJoinMatchmakeSession sets the handler for the JoinMatchmakeSession method
func (protocol *Protocol) SetHandlerJoinMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string) (*nex.RMCMessage, uint32)) {
	protocol.JoinMatchmakeSession = handler
}

// SetHandlerModifyCurrentGameAttribute sets the handler for the ModifyCurrentGameAttribute method
func (protocol *Protocol) SetHandlerModifyCurrentGameAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribIndex *types.PrimitiveU32, newValue *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ModifyCurrentGameAttribute = handler
}

// SetHandlerUpdateNotificationData sets the handler for the UpdateNotificationData method
func (protocol *Protocol) SetHandlerUpdateNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveU32, uiParam1 *types.PrimitiveU32, uiParam2 *types.PrimitiveU32, strParam string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateNotificationData = handler
}

// SetHandlerGetFriendNotificationData sets the handler for the GetFriendNotificationData method
func (protocol *Protocol) SetHandlerGetFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, uiType *types.PrimitiveS32) (*nex.RMCMessage, uint32)) {
	protocol.GetFriendNotificationData = handler
}

// SetHandlerUpdateApplicationBuffer sets the handler for the UpdateApplicationBuffer method
func (protocol *Protocol) SetHandlerUpdateApplicationBuffer(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, applicationBuffer []byte) (*nex.RMCMessage, uint32)) {
	protocol.UpdateApplicationBuffer = handler
}

// SetHandlerUpdateMatchmakeSessionAttribute sets the handler for the UpdateMatchmakeSessionAttribute method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSessionAttribute(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, attribs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMatchmakeSessionAttribute = handler
}

// SetHandlerGetlstFriendNotificationData sets the handler for the GetlstFriendNotificationData method
func (protocol *Protocol) SetHandlerGetlstFriendNotificationData(handler func(err error, packet nex.PacketInterface, callID uint32, lstTypes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.GetlstFriendNotificationData = handler
}

// SetHandlerUpdateMatchmakeSession sets the handler for the UpdateMatchmakeSession method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *types.AnyDataHolder) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMatchmakeSession = handler
}

// SetHandlerAutoMatchmakeWithSearchCriteriaPostpone sets the handler for the AutoMatchmakeWithSearchCriteriaPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithSearchCriteriaPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstSearchCriteria []*match_making_types.MatchmakeSessionSearchCriteria, anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32)) {
	protocol.AutoMatchmakeWithSearchCriteriaPostpone = handler
}

// SetHandlerGetPlayingSession sets the handler for the GetPlayingSession method
func (protocol *Protocol) SetHandlerGetPlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, lstPID *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.GetPlayingSession = handler
}

// SetHandlerCreateCommunity sets the handler for the CreateCommunity method
func (protocol *Protocol) SetHandlerCreateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering, strMessage string) (*nex.RMCMessage, uint32)) {
	protocol.CreateCommunity = handler
}

// SetHandlerUpdateCommunity sets the handler for the UpdateCommunity method
func (protocol *Protocol) SetHandlerUpdateCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, community *match_making_types.PersistentGathering) (*nex.RMCMessage, uint32)) {
	protocol.UpdateCommunity = handler
}

// SetHandlerJoinCommunity sets the handler for the JoinCommunity method
func (protocol *Protocol) SetHandlerJoinCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, strPassword string) (*nex.RMCMessage, uint32)) {
	protocol.JoinCommunity = handler
}

// SetHandlerFindCommunityByGatheringID sets the handler for the FindCommunityByGatheringID method
func (protocol *Protocol) SetHandlerFindCommunityByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.FindCommunityByGatheringID = handler
}

// SetHandlerFindOfficialCommunity sets the handler for the FindOfficialCommunity method
func (protocol *Protocol) SetHandlerFindOfficialCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, isAvailableOnly *types.PrimitiveBool, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.FindOfficialCommunity = handler
}

// SetHandlerFindCommunityByParticipant sets the handler for the FindCommunityByParticipant method
func (protocol *Protocol) SetHandlerFindCommunityByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.FindCommunityByParticipant = handler
}

// SetHandlerUpdatePrivacySetting sets the handler for the UpdatePrivacySetting method
func (protocol *Protocol) SetHandlerUpdatePrivacySetting(handler func(err error, packet nex.PacketInterface, callID uint32, onlineStatus *types.PrimitiveBool, participationCommunity *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.UpdatePrivacySetting = handler
}

// SetHandlerGetMyBlockList sets the handler for the GetMyBlockList method
func (protocol *Protocol) SetHandlerGetMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetMyBlockList = handler
}

// SetHandlerAddToBlockList sets the handler for the AddToBlockList method
func (protocol *Protocol) SetHandlerAddToBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.AddToBlockList = handler
}

// SetHandlerRemoveFromBlockList sets the handler for the RemoveFromBlockList method
func (protocol *Protocol) SetHandlerRemoveFromBlockList(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipalID *types.List[*types.PID]) (*nex.RMCMessage, uint32)) {
	protocol.RemoveFromBlockList = handler
}

// SetHandlerClearMyBlockList sets the handler for the ClearMyBlockList method
func (protocol *Protocol) SetHandlerClearMyBlockList(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.ClearMyBlockList = handler
}

// SetHandlerReportViolation sets the handler for the ReportViolation method
func (protocol *Protocol) SetHandlerReportViolation(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, userName string, violationCode *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ReportViolation = handler
}

// SetHandlerIsViolationUser sets the handler for the IsViolationUser method
func (protocol *Protocol) SetHandlerIsViolationUser(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.IsViolationUser = handler
}

// SetHandlerJoinMatchmakeSessionEx sets the handler for the JoinMatchmakeSessionEx method
func (protocol *Protocol) SetHandlerJoinMatchmakeSessionEx(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, strMessage string, dontCareMyBlockList *types.PrimitiveBool, participationCount *types.PrimitiveU16) (*nex.RMCMessage, uint32)) {
	protocol.JoinMatchmakeSessionEx = handler
}

// SetHandlerGetSimplePlayingSession sets the handler for the GetSimplePlayingSession method
func (protocol *Protocol) SetHandlerGetSimplePlayingSession(handler func(err error, packet nex.PacketInterface, callID uint32, listPID *types.List[*types.PID], includeLoginUser *types.PrimitiveBool) (*nex.RMCMessage, uint32)) {
	protocol.GetSimplePlayingSession = handler
}

// SetHandlerGetSimpleCommunity sets the handler for the GetSimpleCommunity method
func (protocol *Protocol) SetHandlerGetSimpleCommunity(handler func(err error, packet nex.PacketInterface, callID uint32, gatheringIDList *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.GetSimpleCommunity = handler
}

// SetHandlerAutoMatchmakeWithGatheringIDPostpone sets the handler for the AutoMatchmakeWithGatheringIDPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithGatheringIDPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32], anyGathering *types.AnyDataHolder, strMessage string) (*nex.RMCMessage, uint32)) {
	protocol.AutoMatchmakeWithGatheringIDPostpone = handler
}

// SetHandlerUpdateProgressScore sets the handler for the UpdateProgressScore method
func (protocol *Protocol) SetHandlerUpdateProgressScore(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, progressScore *types.PrimitiveU8) (*nex.RMCMessage, uint32)) {
	protocol.UpdateProgressScore = handler
}

// SetHandlerDebugNotifyEvent sets the handler for the DebugNotifyEvent method
func (protocol *Protocol) SetHandlerDebugNotifyEvent(handler func(err error, packet nex.PacketInterface, callID uint32, pid *types.PID, mainType *types.PrimitiveU32, subType *types.PrimitiveU32, param1 *types.PrimitiveU64, param2 *types.PrimitiveU64, stringParam string) (*nex.RMCMessage, uint32)) {
	protocol.DebugNotifyEvent = handler
}

// SetHandlerGenerateMatchmakeSessionSystemPassword sets the handler for the GenerateMatchmakeSessionSystemPassword method
func (protocol *Protocol) SetHandlerGenerateMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.GenerateMatchmakeSessionSystemPassword = handler
}

// SetHandlerClearMatchmakeSessionSystemPassword sets the handler for the ClearMatchmakeSessionSystemPassword method
func (protocol *Protocol) SetHandlerClearMatchmakeSessionSystemPassword(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ClearMatchmakeSessionSystemPassword = handler
}

// SetHandlerCreateMatchmakeSessionWithParam sets the handler for the CreateMatchmakeSessionWithParam method
func (protocol *Protocol) SetHandlerCreateMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, createMatchmakeSessionParam *match_making_types.CreateMatchmakeSessionParam) (*nex.RMCMessage, uint32)) {
	protocol.CreateMatchmakeSessionWithParam = handler
}

// SetHandlerJoinMatchmakeSessionWithParam sets the handler for the JoinMatchmakeSessionWithParam method
func (protocol *Protocol) SetHandlerJoinMatchmakeSessionWithParam(handler func(err error, packet nex.PacketInterface, callID uint32, joinMatchmakeSessionParam *match_making_types.JoinMatchmakeSessionParam) (*nex.RMCMessage, uint32)) {
	protocol.JoinMatchmakeSessionWithParam = handler
}

// SetHandlerAutoMatchmakeWithParamPostpone sets the handler for the AutoMatchmakeWithParamPostpone method
func (protocol *Protocol) SetHandlerAutoMatchmakeWithParamPostpone(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32)) {
	protocol.AutoMatchmakeWithParamPostpone = handler
}

// SetHandlerFindMatchmakeSessionByGatheringIDDetail sets the handler for the FindMatchmakeSessionByGatheringIDDetail method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByGatheringIDDetail(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.FindMatchmakeSessionByGatheringIDDetail = handler
}

// SetHandlerBrowseMatchmakeSessionNoHolder sets the handler for the BrowseMatchmakeSessionNoHolder method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSessionNoHolder = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder sets the handler for the BrowseMatchmakeSessionWithHostURLsNoHolder method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolder(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSessionWithHostURLsNoHolder = handler
}

// SetHandlerUpdateMatchmakeSessionPart sets the handler for the UpdateMatchmakeSessionPart method
func (protocol *Protocol) SetHandlerUpdateMatchmakeSessionPart(handler func(err error, packet nex.PacketInterface, callID uint32, updateMatchmakeSessionParam *match_making_types.UpdateMatchmakeSessionParam) (*nex.RMCMessage, uint32)) {
	protocol.UpdateMatchmakeSessionPart = handler
}

// SetHandlerRequestMatchmaking sets the handler for the RequestMatchmaking method
func (protocol *Protocol) SetHandlerRequestMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, autoMatchmakeParam *match_making_types.AutoMatchmakeParam) (*nex.RMCMessage, uint32)) {
	protocol.RequestMatchmaking = handler
}

// SetHandlerWithdrawMatchmaking sets the handler for the WithdrawMatchmaking method
func (protocol *Protocol) SetHandlerWithdrawMatchmaking(handler func(err error, packet nex.PacketInterface, callID uint32, requestID *types.PrimitiveU64) (*nex.RMCMessage, uint32)) {
	protocol.WithdrawMatchmaking = handler
}

// SetHandlerWithdrawMatchmakingAll sets the handler for the WithdrawMatchmakingAll method
func (protocol *Protocol) SetHandlerWithdrawMatchmakingAll(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.WithdrawMatchmakingAll = handler
}

// SetHandlerFindMatchmakeSessionByGatheringID sets the handler for the FindMatchmakeSessionByGatheringID method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, lstGID *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, uint32)) {
	protocol.FindMatchmakeSessionByGatheringID = handler
}

// SetHandlerFindMatchmakeSessionBySingleGatheringID sets the handler for the FindMatchmakeSessionBySingleGatheringID method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionBySingleGatheringID(handler func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.FindMatchmakeSessionBySingleGatheringID = handler
}

// SetHandlerFindMatchmakeSessionByOwner sets the handler for the FindMatchmakeSessionByOwner method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, id *types.PrimitiveU32, resultRange *types.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.FindMatchmakeSessionByOwner = handler
}

// SetHandlerFindMatchmakeSessionByParticipant sets the handler for the FindMatchmakeSessionByParticipant method
func (protocol *Protocol) SetHandlerFindMatchmakeSessionByParticipant(handler func(err error, packet nex.PacketInterface, callID uint32, param *match_making_types.FindMatchmakeSessionByParticipantParam) (*nex.RMCMessage, uint32)) {
	protocol.FindMatchmakeSessionByParticipant = handler
}

// SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange sets the handler for the BrowseMatchmakeSessionNoHolderNoResultRange method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSessionNoHolderNoResultRange = handler
}

// SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange sets the handler for the BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange method
func (protocol *Protocol) SetHandlerBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(handler func(err error, packet nex.PacketInterface, callID uint32, searchCriteria *match_making_types.MatchmakeSessionSearchCriteria) (*nex.RMCMessage, uint32)) {
	protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange = handler
}

// SetHandlerFindCommunityByOwner sets the handler for the FindCommunityByOwner method
func (protocol *Protocol) SetHandlerFindCommunityByOwner(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.FindCommunityByOwner = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodCloseParticipation:
		protocol.handleCloseParticipation(packet)
	case MethodOpenParticipation:
		protocol.handleOpenParticipation(packet)
	case MethodAutoMatchmakePostpone:
		protocol.handleAutoMatchmakePostpone(packet)
	case MethodBrowseMatchmakeSession:
		protocol.handleBrowseMatchmakeSession(packet)
	case MethodBrowseMatchmakeSessionWithHostURLs:
		protocol.handleBrowseMatchmakeSessionWithHostURLs(packet)
	case MethodCreateMatchmakeSession:
		protocol.handleCreateMatchmakeSession(packet)
	case MethodJoinMatchmakeSession:
		protocol.handleJoinMatchmakeSession(packet)
	case MethodModifyCurrentGameAttribute:
		protocol.handleModifyCurrentGameAttribute(packet)
	case MethodUpdateNotificationData:
		protocol.handleUpdateNotificationData(packet)
	case MethodGetFriendNotificationData:
		protocol.handleGetFriendNotificationData(packet)
	case MethodUpdateApplicationBuffer:
		protocol.handleUpdateApplicationBuffer(packet)
	case MethodUpdateMatchmakeSessionAttribute:
		protocol.handleUpdateMatchmakeSessionAttribute(packet)
	case MethodGetlstFriendNotificationData:
		protocol.handleGetlstFriendNotificationData(packet)
	case MethodUpdateMatchmakeSession:
		protocol.handleUpdateMatchmakeSession(packet)
	case MethodAutoMatchmakeWithSearchCriteriaPostpone:
		protocol.handleAutoMatchmakeWithSearchCriteriaPostpone(packet)
	case MethodGetPlayingSession:
		protocol.handleGetPlayingSession(packet)
	case MethodCreateCommunity:
		protocol.handleCreateCommunity(packet)
	case MethodUpdateCommunity:
		protocol.handleUpdateCommunity(packet)
	case MethodJoinCommunity:
		protocol.handleJoinCommunity(packet)
	case MethodFindCommunityByGatheringID:
		protocol.handleFindCommunityByGatheringID(packet)
	case MethodFindOfficialCommunity:
		protocol.handleFindOfficialCommunity(packet)
	case MethodFindCommunityByParticipant:
		protocol.handleFindCommunityByParticipant(packet)
	case MethodUpdatePrivacySetting:
		protocol.handleUpdatePrivacySetting(packet)
	case MethodGetMyBlockList:
		protocol.handleGetMyBlockList(packet)
	case MethodAddToBlockList:
		protocol.handleAddToBlockList(packet)
	case MethodRemoveFromBlockList:
		protocol.handleRemoveFromBlockList(packet)
	case MethodClearMyBlockList:
		protocol.handleClearMyBlockList(packet)
	case MethodReportViolation:
		protocol.handleReportViolation(packet)
	case MethodIsViolationUser:
		protocol.handleIsViolationUser(packet)
	case MethodJoinMatchmakeSessionEx:
		protocol.handleJoinMatchmakeSessionEx(packet)
	case MethodGetSimplePlayingSession:
		protocol.handleGetSimplePlayingSession(packet)
	case MethodGetSimpleCommunity:
		protocol.handleGetSimpleCommunity(packet)
	case MethodAutoMatchmakeWithGatheringIDPostpone:
		protocol.handleAutoMatchmakeWithGatheringIDPostpone(packet)
	case MethodUpdateProgressScore:
		protocol.handleUpdateProgressScore(packet)
	case MethodDebugNotifyEvent:
		protocol.handleDebugNotifyEvent(packet)
	case MethodGenerateMatchmakeSessionSystemPassword:
		protocol.handleGenerateMatchmakeSessionSystemPassword(packet)
	case MethodClearMatchmakeSessionSystemPassword:
		protocol.handleClearMatchmakeSessionSystemPassword(packet)
	case MethodCreateMatchmakeSessionWithParam:
		protocol.handleCreateMatchmakeSessionWithParam(packet)
	case MethodJoinMatchmakeSessionWithParam:
		protocol.handleJoinMatchmakeSessionWithParam(packet)
	case MethodAutoMatchmakeWithParamPostpone:
		protocol.handleAutoMatchmakeWithParamPostpone(packet)
	case MethodFindMatchmakeSessionByGatheringIDDetail:
		protocol.handleFindMatchmakeSessionByGatheringIDDetail(packet)
	case MethodBrowseMatchmakeSessionNoHolder:
		protocol.handleBrowseMatchmakeSessionNoHolder(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolder:
		protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolder(packet)
	case MethodUpdateMatchmakeSessionPart:
		protocol.handleUpdateMatchmakeSessionPart(packet)
	case MethodRequestMatchmaking:
		protocol.handleRequestMatchmaking(packet)
	case MethodWithdrawMatchmaking:
		protocol.handleWithdrawMatchmaking(packet)
	case MethodWithdrawMatchmakingAll:
		protocol.handleWithdrawMatchmakingAll(packet)
	case MethodFindMatchmakeSessionByGatheringID:
		protocol.handleFindMatchmakeSessionByGatheringID(packet)
	case MethodFindMatchmakeSessionBySingleGatheringID:
		protocol.handleFindMatchmakeSessionBySingleGatheringID(packet)
	case MethodFindMatchmakeSessionByOwner:
		protocol.handleFindMatchmakeSessionByOwner(packet)
	case MethodFindMatchmakeSessionByParticipant:
		protocol.handleFindMatchmakeSessionByParticipant(packet)
	case MethodBrowseMatchmakeSessionNoHolderNoResultRange:
		protocol.handleBrowseMatchmakeSessionNoHolderNoResultRange(packet)
	case MethodBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange:
		protocol.handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet)
	case MethodFindCommunityByOwner:
		protocol.handleFindCommunityByOwner(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Matchmake Extension protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
