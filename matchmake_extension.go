package nexproto

import (
	"encoding/hex"
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchmakeExtensionProtocolID is the protocol ID for the Matchmake Extension protocol
	MatchmakeExtensionProtocolID = 0x6D

	// MatchmakeExtensionMethodCloseParticipation is the method ID for method CloseParticipation
	MatchmakeExtensionMethodCloseParticipation = 0x1

	// MatchmakeExtensionMethodOpenParticipation is the method ID for method OpenParticipation
	MatchmakeExtensionMethodOpenParticipation = 0x2

	// MatchmakeExtensionMethodAutoMatchmake_Postpone is the method ID for method AutoMatchmake_Postpone
	MatchmakeExtensionMethodAutoMatchmake_Postpone = 0x3

	// MatchmakeExtensionMethodBrowseMatchmakeSession is the method ID for method BrowseMatchmakeSession
	MatchmakeExtensionMethodBrowseMatchmakeSession = 0x4

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodCreateMatchmakeSession = 0x6

	// MatchmakeExtensionMethodUpdateNotificationData is the method ID for method UpdateNotificationData
	MatchmakeExtensionMethodUpdateNotificationData = 0x8

	// MatchmakeExtensionMethodGetFriendNotificationData is the method ID for method GetFriendNotificationData
	MatchmakeExtensionMethodGetFriendNotificationData = 0xA

	// MatchmakeExtensionMethodUpdateApplicationBuffer is the method ID for method UpdateApplicationBuffer
	MatchmakeExtensionMethodUpdateApplicationBuffer = 0xb

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodAutoMatchmakeWithSearchCriteria_Postpone = 0xF

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodGetPlayingSession = 0x10

	// MatchmakeExtensionMethodCreateCommunity is the method ID for method CreateCommunity
	MatchmakeExtensionMethodCreateCommunity = 0x11

	// MatchmakeExtensionMethodFindCommunityByGatheringId is the method ID for method FindCommunityByGatheringId
	MatchmakeExtensionMethodFindCommunityByGatheringId = 0x14

	// MatchmakeExtensionMethodFindOfficialCommunity is the method ID for method FindOfficialCommunity
	MatchmakeExtensionMethodFindOfficialCommunity = 0x15

	// MatchmakeExtensionMethodFindCommunityByParticipant is the method ID for method FindCommunityByParticipant
	MatchmakeExtensionMethodFindCommunityByParticipant = 0x16

	// MatchmakeExtensionMethodJoinMatchmakeSessionEx is the method ID for method JoinMatchmakeSessionEx
	MatchmakeExtensionMethodJoinMatchmakeSessionEx = 0x1E
	
	// MatchmakeExtensionMethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MatchmakeExtensionMethodGetSimplePlayingSession = 0x1F

	// MatchmakeExtensionMethodGetSimpleCommunity is the method ID for method GetSimpleCommunity
	MatchmakeExtensionMethodGetSimpleCommunity = 0x20

	// MatchmakeExtensionMethodUpdateProgressScore is the method ID for method UpdateProgressScore
	MatchmakeExtensionMethodUpdateProgressScore = 0x22

	// MatchmakeExtensionMethodCreateMatchmakeSessionWithParam is the method ID for method CreateMatchmakeSessionWithParam
	MatchmakeExtensionMethodCreateMatchmakeSessionWithParam = 0x26

	// MatchmakeExtensionMethodJoinMatchmakeSessionWithParam is the method ID for method JoinMatchmakeSessionWithParam
	MatchmakeExtensionMethodJoinMatchmakeSessionWithParam = 0x27

	// AutoMatchmakeWithParam_Postpone is the method ID for method AutoMatchmakeWithParam_Postpone
	MatchmakeExtensionMethodAutoMatchmakeWithParam_Postpone = 0x28
)

// MatchmakeExtensionProtocol handles the Matchmake Extension nex protocol
type MatchmakeExtensionProtocol struct {
	server                                          *nex.Server
	CloseParticipationHandler                       func(err error, client *nex.Client, callID uint32, gid uint32)
	OpenParticipationHandler                        func(err error, client *nex.Client, callID uint32, gid uint32)
	AutoMatchmake_PostponeHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string)
	BrowseMatchmakeSessionHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria)
	CreateMatchmakeSessionHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string, participationCount uint16)
	UpdateNotificationDataHandler                   func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)
	GetFriendNotificationDataHandler                func(err error, client *nex.Client, callID uint32, uiType int32)
	UpdateApplicationBufferHandler                  func(err error, client *nex.Client, callID uint32, gid uint32, applicationBuffer []byte)
	AutoMatchmakeWithSearchCriteria_PostponeHandler func(err error, client *nex.Client, callID uint32, matchmakeSessionSearchCriteria []*MatchmakeSessionSearchCriteria, matchmakeSession *MatchmakeSession, message string)
	FindCommunityByGatheringIdHandler               func(err error, client *nex.Client, callID uint32, lstGid []uint32)
	FindOfficialCommunityHandler                    func(err error, client *nex.Client, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange)
	FindCommunityByParticipantHandler               func(err error, client *nex.Client, callID uint32, pid uint32, resultRange *nex.ResultRange)
	GetSimplePlayingSessionHandler                  func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)
	JoinMatchmakeSessionExHandler                   func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)
	GetSimpleCommunityHandler                       func(err error, client *nex.Client, callID uint32, gatheringIdList []uint32)
	UpdateProgressScoreHandler                      func(err error, client *nex.Client, callID uint32, gid uint32, progressScore uint8)
	GetPlayingSessionHandler                        func(err error, client *nex.Client, callID uint32, listPID []uint32)
	CreateCommunityHandler                          func(err error, client *nex.Client, callID uint32, community *PersistentGathering, strMessage string)
	CreateMatchmakeSessionWithParamHandler          func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession)
	JoinMatchmakeSessionWithParamHandler            func(err error, client *nex.Client, callID uint32, gid uint32)
	AutoMatchmakeWithParam_PostponeHandler          func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, sourceGid uint32)
}

// Setup initializes the protocol
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) Setup() {
	nexServer := matchmakeExtensionProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchmakeExtensionProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchmakeExtensionMethodCloseParticipation:
				go matchmakeExtensionProtocol.handleCloseParticipation(packet)
			case MatchmakeExtensionMethodOpenParticipation:
				go matchmakeExtensionProtocol.handleOpenParticipation(packet)
			case MatchmakeExtensionMethodAutoMatchmake_Postpone:
				go matchmakeExtensionProtocol.handleAutoMatchmake_Postpone(packet)
			case MatchmakeExtensionMethodBrowseMatchmakeSession:
				go matchmakeExtensionProtocol.handleBrowseMatchmakeSession(packet)
			case MatchmakeExtensionMethodCreateMatchmakeSession:
				go matchmakeExtensionProtocol.handleCreateMatchmakeSession(packet)
			case MatchmakeExtensionMethodUpdateNotificationData:
				go matchmakeExtensionProtocol.handleUpdateNotificationData(packet)
			case MatchmakeExtensionMethodGetFriendNotificationData:
				go matchmakeExtensionProtocol.handleGetFriendNotificationData(packet)
			case MatchmakeExtensionMethodUpdateApplicationBuffer:
				go matchmakeExtensionProtocol.handleUpdateApplicationBuffer(packet)
			case MatchmakeExtensionMethodAutoMatchmakeWithSearchCriteria_Postpone:
				go matchmakeExtensionProtocol.handleAutoMatchmakeWithSearchCriteria_Postpone(packet)
			case MatchmakeExtensionMethodJoinMatchmakeSessionEx:
				go matchmakeExtensionProtocol.handleJoinMatchmakeSessionEx(packet)
			case MatchmakeExtensionMethodGetPlayingSession:
				go matchmakeExtensionProtocol.handleGetPlayingSession(packet)
			case MatchmakeExtensionMethodCreateCommunity:
				go matchmakeExtensionProtocol.handleCreateCommunity(packet)
			case MatchmakeExtensionMethodFindCommunityByGatheringId:
				go matchmakeExtensionProtocol.handleFindCommunityByGatheringId(packet)
			case MatchmakeExtensionMethodFindOfficialCommunity:
				go matchmakeExtensionProtocol.handleFindOfficialCommunity(packet)
			case MatchmakeExtensionMethodFindCommunityByParticipant:
				go matchmakeExtensionProtocol.handleFindCommunityByParticipant(packet)
			case MatchmakeExtensionMethodGetSimplePlayingSession:
				go matchmakeExtensionProtocol.handleGetSimplePlayingSession(packet)
			case MatchmakeExtensionMethodGetSimpleCommunity:
				go matchmakeExtensionProtocol.handleGetSimpleCommunity(packet)
			case MatchmakeExtensionMethodUpdateProgressScore:
				go matchmakeExtensionProtocol.handleUpdateProgressScore(packet)
			case MatchmakeExtensionMethodCreateMatchmakeSessionWithParam:
				go matchmakeExtensionProtocol.handleCreateMatchmakeSessionWithParam(packet)
			case MatchmakeExtensionMethodJoinMatchmakeSessionWithParam:
				go matchmakeExtensionProtocol.handleJoinMatchmakeSessionWithParam(packet)
			case MatchmakeExtensionMethodAutoMatchmakeWithParam_Postpone:
				go matchmakeExtensionProtocol.handleAutoMatchmakeWithParam_Postpone(packet)
			default:
				go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
				fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// OpenParticipation sets the OpenParticipation handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) CloseParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	matchmakeExtensionProtocol.CloseParticipationHandler = handler
}

// OpenParticipation sets the OpenParticipation handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) OpenParticipation(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	matchmakeExtensionProtocol.OpenParticipationHandler = handler
}

// AutoMatchmake_Postpone sets the AutoMatchmake_Postpone handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) AutoMatchmake_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string)) {
	matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler = handler
}

// CreateMatchmakeSession sets the CreateMatchmakeSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) CreateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string, participationCount uint16)) {
	matchmakeExtensionProtocol.CreateMatchmakeSessionHandler = handler
}

// BrowseMatchmakeSession sets the CreateMatchmakeSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) BrowseMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, matchmakeSessionSearchCriteria *MatchmakeSessionSearchCriteria)) {
	matchmakeExtensionProtocol.BrowseMatchmakeSessionHandler = handler
}

// UpdateNotificationData sets the UpdateNotificationData handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) UpdateNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)) {
	matchmakeExtensionProtocol.UpdateNotificationDataHandler = handler
}

// GetFriendNotificationData sets the GetFriendNotificationData handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetFriendNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType int32)) {
	matchmakeExtensionProtocol.GetFriendNotificationDataHandler = handler
}

// UpdateApplicationBuffer sets the UpdateApplicationBuffer handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) UpdateApplicationBuffer(handler func(err error, client *nex.Client, callID uint32, gid uint32, applicationBuffer[]byte)) {
	matchmakeExtensionProtocol.UpdateApplicationBufferHandler = handler
}

// AutoMatchmakeWithSearchCriteria_Postpone sets the AutoMatchmakeWithSearchCriteria_Postpone handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) AutoMatchmakeWithSearchCriteria_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSessionSearchCriteria []*MatchmakeSessionSearchCriteria, matchmakeSession *MatchmakeSession, message string)) {
	matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler = handler
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetPlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32)) {
	matchmakeExtensionProtocol.GetPlayingSessionHandler = handler
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) CreateCommunity(handler func(err error, client *nex.Client, callID uint32, community *PersistentGathering, strMessage string)) {
	matchmakeExtensionProtocol.CreateCommunityHandler = handler
}

// FindOfficialCommunity sets the FindOfficialCommunity handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) FindOfficialCommunity(handler func(err error, client *nex.Client, callID uint32, isAvailableOnly bool, resultRange *nex.ResultRange)) {
	matchmakeExtensionProtocol.FindOfficialCommunityHandler = handler
}

// FindCommunityByGatheringId sets the FindCommunityByGatheringId handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) FindCommunityByGatheringId(handler func(err error, client *nex.Client, callID uint32, lstGid []uint32)) {
	matchmakeExtensionProtocol.FindCommunityByGatheringIdHandler = handler
}

// FindCommunityByParticipant sets the FindCommunityByParticipant handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) FindCommunityByParticipant(handler func(err error, client *nex.Client, callID uint32, pid uint32, resultRange *nex.ResultRange)) {
	matchmakeExtensionProtocol.FindCommunityByParticipantHandler = handler
}

// JoinMatchmakeSessionEx sets the JoinMatchmakeSessionEx handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionEx(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)) {
	matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler = handler
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)) {
	matchmakeExtensionProtocol.GetSimplePlayingSessionHandler = handler
}

// GetSimpleCommunity sets the GetSimpleCommunity handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetSimpleCommunity(handler func(err error, client *nex.Client, callID uint32, gatheringIdList []uint32)) {
	matchmakeExtensionProtocol.GetSimpleCommunityHandler = handler
}

// UpdateProgressScore sets the UpdateProgressScore handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) UpdateProgressScore(handler func(err error, client *nex.Client, callID uint32, gid uint32, progressScore uint8)) {
	matchmakeExtensionProtocol.UpdateProgressScoreHandler = handler
}

// CreateMatchmakeSessionWithParam sets the CreateMatchmakeSessionWithParam handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) CreateMatchmakeSessionWithParam(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession)) {
	matchmakeExtensionProtocol.CreateMatchmakeSessionWithParamHandler = handler
}

// JoinMatchmakeSessionWithParam sets the JoinMatchmakeSessionWithParam handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionWithParam(handler func(err error, client *nex.Client, callID uint32, gid uint32)) {
	matchmakeExtensionProtocol.JoinMatchmakeSessionWithParamHandler = handler
}
// AutoMatchmakeWithParam_Postpone sets the AutoMatchmakeWithParam_Postpone handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) AutoMatchmakeWithParam_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, sourceGid uint32)) {
	matchmakeExtensionProtocol.AutoMatchmakeWithParam_PostponeHandler = handler
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleCloseParticipation(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.CloseParticipationHandler == nil {
		fmt.Println("[Warning] MatchMakingProtocol::CloseParticipation not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	//parameters := request.Parameters()

	//parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	gatheringId := (uint32)(0) //parametersStream.ReadUInt32LE()

	go matchmakeExtensionProtocol.CloseParticipationHandler(nil, client, callID, gatheringId)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleOpenParticipation(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.OpenParticipationHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::OpenParticipation not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	gid := parametersStream.ReadUInt32LE()

	go matchmakeExtensionProtocol.OpenParticipationHandler(nil, client, callID, gid)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleAutoMatchmake_Postpone(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::AutoMatchmake_PostponeHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmake_Postpone] Data holder name does not match")
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmake_Postpone] Data holder missing lengths")
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	//fmt.Println("%i", parametersStream.ReadUInt32LE())
	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, matchmakeExtensionProtocol.server)

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*MatchmakeSession)

	message, err := parametersStream.ReadString()
	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(nil, client, callID, matchmakeSession, message)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleBrowseMatchmakeSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.BrowseMatchmakeSessionHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::BrowseMatchmakeSessionHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	criteriaTmp, _ := parametersStream.ReadStructure(NewMatchmakeSessionSearchCriteria())
	criteria := criteriaTmp.(*MatchmakeSessionSearchCriteria)

	go matchmakeExtensionProtocol.BrowseMatchmakeSessionHandler(nil, client, callID, criteria)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.CreateMatchmakeSessionHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::CreateMatchmakeSession not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtensionProtocol::CreateMatchmakeSession] Data holder name does not match")
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtensionProtocol::CreateMatchmakeSession] Data holder missing lengths")
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, matchmakeExtensionProtocol.server)

	matchmakeSession, err := dataHolderContentStream.ReadStructure(NewMatchmakeSession())

	if err != nil {
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	message, err := parametersStream.ReadString()

	if err != nil {
		go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	var participationCount uint16 = 0

	if matchmakeExtensionProtocol.server.NexVersion() >= 30500 {
		participationCount = dataHolderContentStream.ReadUInt16LE()
	}

	go matchmakeExtensionProtocol.CreateMatchmakeSessionHandler(nil, client, callID, matchmakeSession.(*MatchmakeSession), message, participationCount)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.UpdateNotificationDataHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::UpdateNotificationData not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	
	go matchmakeExtensionProtocol.UpdateNotificationDataHandler(nil, client, callID, 0, 0, 0, "")
	return

	uiType := parametersStream.ReadUInt32LE()
	uiParam1 := parametersStream.ReadUInt32LE()
	uiParam2 := parametersStream.ReadUInt32LE()
	strParam, err := parametersStream.ReadString()
	if err != nil {
		go matchmakeExtensionProtocol.UpdateNotificationDataHandler(err, client, callID, 0, 0, 0, "")
		return
	}

	go matchmakeExtensionProtocol.UpdateNotificationDataHandler(nil, client, callID, uiType, uiParam1, uiParam2, strParam)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetFriendNotificationData(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetFriendNotificationDataHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::GetFriendNotificationData not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	uiType := int32(parametersStream.ReadUInt32LE())

	go matchmakeExtensionProtocol.GetFriendNotificationDataHandler(nil, client, callID, uiType)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleUpdateApplicationBuffer(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.UpdateApplicationBufferHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::UpdateApplicationBuffer not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	gid := parametersStream.ReadUInt32LE()
	applicationBuffer, _ := parametersStream.ReadBuffer()

	go matchmakeExtensionProtocol.UpdateApplicationBufferHandler(nil, client, callID, gid, applicationBuffer)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithSearchCriteria_Postpone(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_PostponeHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	criteriaCount := int(parametersStream.ReadUInt32LE())
	criteriaList := []*MatchmakeSessionSearchCriteria{}
	for i := 0; i < criteriaCount; i++ {
		criteria, _ := parametersStream.ReadStructure(NewMatchmakeSessionSearchCriteria())
		//criteria = 
		criteriaList = append(criteriaList, criteria.(*MatchmakeSessionSearchCriteria))
	}
	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}

	fmt.Println(dataHolderType)

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_Postpone] Data holder name does not match")
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_Postpone] Data holder missing lengths")
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, matchmakeExtensionProtocol.server)

	//gatheringStructureInterface, err := dataHolderContentStream.ReadStructure(NewGathering())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*MatchmakeSession)
	//matchmakeSession.Gathering = gatheringStructureInterface.(*Gathering)

	/*message, err := parametersStream.ReadString()
	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, nil, "")
		return
	}*/

	go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, criteriaList, matchmakeSession, "message")
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetPlayingSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetPlayingSessionHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::GetPlayingSession not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	listPID := parametersStream.ReadListUInt32LE()

	go matchmakeExtensionProtocol.GetPlayingSessionHandler(nil, client, callID, listPID)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleCreateCommunity(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.CreateCommunityHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::CreateCommunity not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	
	communityStructureInterface, err := parametersStream.ReadStructure(NewPersistentGathering())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.CreateCommunityHandler(err, client, callID, nil, "")
		return
	}
	community := communityStructureInterface.(*PersistentGathering)
	strMessage, err := parametersStream.ReadString()
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.CreateCommunityHandler(err, client, callID, nil, "")
		return
	}

	go matchmakeExtensionProtocol.CreateCommunityHandler(nil, client, callID, community, strMessage)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleFindCommunityByGatheringId(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.FindCommunityByGatheringIdHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::FindCommunityByGatheringId not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	lstGid := parametersStream.ReadListUInt32LE()

	go matchmakeExtensionProtocol.FindCommunityByGatheringIdHandler(nil, client, callID, lstGid)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleFindOfficialCommunity(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.FindOfficialCommunityHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::FindOfficialCommunity not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	isAvailableOnly := parametersStream.ReadBool()
	resultRangeStructureInterface, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.FindCommunityByParticipantHandler(err, client, callID, 0, nil)
		return
	}
	resultRange := resultRangeStructureInterface.(*nex.ResultRange)

	go matchmakeExtensionProtocol.FindOfficialCommunityHandler(nil, client, callID, isAvailableOnly, resultRange)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleFindCommunityByParticipant(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.FindCommunityByParticipantHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::FindCommunityByParticipant not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	pid := parametersStream.ReadUInt32LE()
	resultRangeStructureInterface, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.FindCommunityByParticipantHandler(err, client, callID, 0, nil)
		return
	}
	resultRange := resultRangeStructureInterface.(*nex.ResultRange)

	go matchmakeExtensionProtocol.FindCommunityByParticipantHandler(nil, client, callID, pid, resultRange)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler == nil {
		logger.Warning("MatchmakeExtensionProtocol::JoinMatchmakeSessionEx not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	gid := parametersStream.ReadUInt32LE()
	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler(err, client, callID, 0, "", false, 0)
		return
	}

	dontCareMyBlockList := parametersStream.ReadBool()
	//participationCount := parametersStream.ReadUInt16LE()

	go matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler(nil, client, callID, gid, strMessage, dontCareMyBlockList, 0)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetSimplePlayingSessionHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::GetSimplePlayingSession not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	listPID := parametersStream.ReadListUInt32LE()
	includeLoginUser := parametersStream.ReadUInt8() == 1

	go matchmakeExtensionProtocol.GetSimplePlayingSessionHandler(nil, client, callID, listPID, includeLoginUser)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetSimpleCommunity(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetSimpleCommunityHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::GetSimpleCommunity not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	gatheringIdList := parametersStream.ReadListUInt32LE()

	go matchmakeExtensionProtocol.GetSimpleCommunityHandler(nil, client, callID, gatheringIdList)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleUpdateProgressScore(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.UpdateProgressScoreHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::UpdateProgressScore not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	gid := parametersStream.ReadUInt32LE()
	progressScore := parametersStream.ReadUInt8()

	go matchmakeExtensionProtocol.UpdateProgressScoreHandler(nil, client, callID, gid, progressScore)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleCreateMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.CreateMatchmakeSessionWithParamHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::CreateMatchmakeSessionWithParamHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	parametersStream.SeekByte(5, false)

	gatheringStructureInterface, err := parametersStream.ReadStructure(NewGathering())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.CreateMatchmakeSessionWithParamHandler(err, client, callID, nil)
		return
	}

	matchmakeSessionStructureInterface, err := parametersStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.CreateMatchmakeSessionWithParamHandler(err, client, callID, nil)
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*MatchmakeSession)
	matchmakeSession.Gathering = gatheringStructureInterface.(*Gathering)

	go matchmakeExtensionProtocol.CreateMatchmakeSessionWithParamHandler(nil, client, callID, matchmakeSessionStructureInterface.(*MatchmakeSession))
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleJoinMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.JoinMatchmakeSessionWithParamHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::JoinMatchmakeSessionWithParamHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	parametersStream.SeekByte(5, false)
	gid := parametersStream.ReadUInt32LE()

	go matchmakeExtensionProtocol.JoinMatchmakeSessionWithParamHandler(nil, client, callID, gid)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithParam_Postpone(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.AutoMatchmakeWithParam_Postpone == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::AutoMatchmakeWithParam_Postpone not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	fmt.Println(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)
	parametersStream.SeekByte(5, false)

	gatheringStructureInterface, err := parametersStream.ReadStructure(NewGathering())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmakeWithParam_PostponeHandler(err, client, callID, nil, 0)
		return
	}

	matchmakeSessionStructureInterface, err := parametersStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		fmt.Println(err)
		go matchmakeExtensionProtocol.AutoMatchmakeWithParam_PostponeHandler(err, client, callID, nil, 0)
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*MatchmakeSession)
	matchmakeSession.Gathering = gatheringStructureInterface.(*Gathering)

	//hacky thing to get splatfest source gathering ID
	var sourceGid uint32
	parametersStream.SeekByte(0x82, false)
	if(parametersStream.ReadUInt8() != 0){
		parametersStream.SeekByte(0x82, false)
		parametersStream.SeekByte((int64)(0x86 + (parametersStream.ReadUInt32LE()*4)), false)
		sourceGid = parametersStream.ReadUInt32LE()
	}

	go matchmakeExtensionProtocol.AutoMatchmakeWithParam_PostponeHandler(nil, client, callID, matchmakeSession, sourceGid)
}

// NewMatchmakeExtensionProtocol returns a new MatchmakeExtensionProtocol
func NewMatchmakeExtensionProtocol(server *nex.Server) *MatchmakeExtensionProtocol {
	matchmakeExtensionProtocol := &MatchmakeExtensionProtocol{server: server}

	matchmakeExtensionProtocol.Setup()

	return matchmakeExtensionProtocol
}
