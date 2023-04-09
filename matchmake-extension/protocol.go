package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

const (
	// ProtocolID is the protocol ID for the Matchmake Extension protocol
	ProtocolID = 0x6D

	// MethodCloseParticipation is the method ID for method CloseParticipation
	MethodCloseParticipation = 0x1

	// MethodOpenParticipation is the method ID for method OpenParticipation
	MethodOpenParticipation = 0x2

	// MethodAutoMatchmake_Postpone is the method ID for method AutoMatchmake_Postpone
	MethodAutoMatchmake_Postpone = 0x3

	// MethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MethodCreateMatchmakeSession = 0x6

	// MethodUpdateNotificationData is the method ID for method UpdateNotificationData
	MethodUpdateNotificationData = 0x9

	// MethodGetFriendNotificationData is the method ID for method GetFriendNotificationData
	MethodGetFriendNotificationData = 0xA

	// MethodAutoMatchmakeWithSearchCriteria_Postpone is the method ID for method AutoMatchmakeWithSearchCriteria_Postpone
	MethodAutoMatchmakeWithSearchCriteria_Postpone = 0xF

	// MethodGetPlayingSession is the method ID for method GetPlayingSession
	MethodGetPlayingSession = 0x10

	// MethodCreateCommunity is the method ID for method CreateCommunity
	MethodCreateCommunity = 0x11

	// MethodFindCommunityByGatheringID is the method ID for method FindCommunityByGatheringID
	MethodFindCommunityByGatheringID = 0x14

	// MethodFindOfficialCommunity is the method ID for method FindOfficialCommunity
	MethodFindOfficialCommunity = 0x15

	// MethodFindCommunityByParticipant is the method ID for method FindCommunityByParticipant
	MethodFindCommunityByParticipant = 0x16

	// MethodJoinMatchmakeSessionEx is the method ID for method JoinMatchmakeSessionEx
	MethodJoinMatchmakeSessionEx = 0x1E

	// MethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MethodGetSimplePlayingSession = 0x1F

	// MethodUpdateProgressScore is the method ID for method UpdateProgressScore
	MethodUpdateProgressScore = 0x22

	// MethodCreateMatchmakeSessionWithParam is the method ID for method CreateMatchmakeSessionWithParam
	MethodCreateMatchmakeSessionWithParam = 0x26

	// MethodJoinMatchmakeSessionWithParam is the method ID for method JoinMatchmakeSessionWithParam
	MethodJoinMatchmakeSessionWithParam = 0x27

	// MethodAutoMatchmakeWithParam_Postpone is the method ID for method AutoMatchmakeWithParam_Postpone
	MethodAutoMatchmakeWithParam_Postpone = 0x28
)

// MatchmakeExtensionProtocol handles the Matchmake Extension nex protocol
type MatchmakeExtensionProtocol struct {
	Server                                          *nex.Server
	CloseParticipationHandler                       func(err error, client *nex.Client, callID uint32, gid uint32)
	OpenParticipationHandler                        func(err error, client *nex.Client, callID uint32, gid uint32)
	AutoMatchmake_PostponeHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string)
	CreateMatchmakeSessionHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string, participationCount uint16)
	UpdateNotificationDataHandler                   func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)
	GetFriendNotificationDataHandler                func(err error, client *nex.Client, callID uint32, uiType int32)
	AutoMatchmakeWithSearchCriteria_PostponeHandler func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string)
	GetPlayingSessionHandler                        func(err error, client *nex.Client, callID uint32, lstPID []uint32)
	CreateCommunityHandler                          func(err error, client *nex.Client, callID uint32, community *match_making.PersistentGathering, strMessage string)
	FindCommunityByGatheringIDHandler               func(err error, client *nex.Client, callID uint32, lstGID []uint32)
	FindOfficialCommunityHandler                    func(err error, client *nex.Client, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange)
	FindCommunityByParticipantHandler               func(err error, client *nex.Client, callID uint32, pid uint32, resultRange *nex.ResultRange)
	JoinMatchmakeSessionExHandler                   func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)
	GetSimplePlayingSessionHandler                  func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)
	UpdateProgressScoreHandler                      func(err error, client *nex.Client, callID uint32, GID uint32, progressScore uint8)
	CreateMatchmakeSessionWithParamHandler          func(err error, client *nex.Client, callID uint32, createMatchmakeSessionParam *match_making.CreateMatchmakeSessionParam)
	JoinMatchmakeSessionWithParamHandler            func(err error, client *nex.Client, callID uint32, joinMatchmakeSessionParam *match_making.JoinMatchmakeSessionParam)
	AutoMatchmakeWithParam_PostponeHandler          func(err error, client *nex.Client, callID uint32, autoMatchmakeParam *match_making.AutoMatchmakeParam)
}

// Setup initializes the protocol
func (protocol *MatchmakeExtensionProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *MatchmakeExtensionProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodCloseParticipation:
		go protocol.HandleCloseParticipation(packet)
	case MethodOpenParticipation:
		go protocol.HandleOpenParticipation(packet)
	case MethodAutoMatchmake_Postpone:
		go protocol.HandleAutoMatchmake_Postpone(packet)
	case MethodCreateMatchmakeSession:
		go protocol.HandleCreateMatchmakeSession(packet)
	case MethodUpdateNotificationData:
		go protocol.HandleUpdateNotificationData(packet)
	case MethodGetFriendNotificationData:
		go protocol.HandleGetFriendNotificationData(packet)
	case MethodAutoMatchmakeWithSearchCriteria_Postpone:
		go protocol.HandleAutoMatchmakeWithSearchCriteria_Postpone(packet)
	case MethodGetPlayingSession:
		go protocol.HandleGetPlayingSession(packet)
	case MethodCreateCommunity:
		go protocol.HandleCreateCommunity(packet)
	case MethodFindCommunityByGatheringID:
		go protocol.HandleFindCommunityByGatheringID(packet)
	case MethodFindOfficialCommunity:
		go protocol.HandleFindOfficialCommunity(packet)
	case MethodFindCommunityByParticipant:
		go protocol.HandleFindCommunityByParticipant(packet)
	case MethodJoinMatchmakeSessionEx:
		go protocol.HandleJoinMatchmakeSessionEx(packet)
	case MethodGetSimplePlayingSession:
		go protocol.HandleGetSimplePlayingSession(packet)
	case MethodUpdateProgressScore:
		go protocol.HandleUpdateProgressScore(packet)
	case MethodCreateMatchmakeSessionWithParam:
		go protocol.HandleCreateMatchmakeSessionWithParam(packet)
	case MethodJoinMatchmakeSessionWithParam:
		go protocol.HandleJoinMatchmakeSessionWithParam(packet)
	case MethodAutoMatchmakeWithParam_Postpone:
		go protocol.HandleAutoMatchmakeWithParam_Postpone(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID())
	}
}

// NewMatchmakeExtensionProtocol returns a new MatchmakeExtensionProtocol
func NewMatchmakeExtensionProtocol(server *nex.Server) *MatchmakeExtensionProtocol {
	protocol := &MatchmakeExtensionProtocol{Server: server}

	protocol.Setup()

	return protocol
}
