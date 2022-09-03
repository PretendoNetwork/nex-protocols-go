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

	// MatchmakeExtensionMethodOpenParticipation is the method ID for method OpenParticipation
	MatchmakeExtensionMethodOpenParticipation = 0x2

	// MatchmakeExtensionMethodAutoMatchmake_Postpone is the method ID for method AutoMatchmake_Postpone
	MatchmakeExtensionMethodAutoMatchmake_Postpone = 0x3

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodCreateMatchmakeSession = 0x6

	// MatchmakeExtensionMethodUpdateNotificationData is the method ID for method UpdateNotificationData
	MatchmakeExtensionMethodUpdateNotificationData = 0x9

	// MatchmakeExtensionMethodGetFriendNotificationData is the method ID for method GetFriendNotificationData
	MatchmakeExtensionMethodGetFriendNotificationData = 0xA

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodAutoMatchmakeWithSearchCriteria_Postpone = 0xF

	// MatchmakeExtensionMethodJoinMatchmakeSessionEx is the method ID for method JoinMatchmakeSessionEx
	MatchmakeExtensionMethodJoinMatchmakeSessionEx = 0x1E

	// MatchmakeExtensionMethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MatchmakeExtensionMethodGetSimplePlayingSession = 0x1F
)

// MatchmakeExtensionProtocol handles the Matchmake Extension nex protocol
type MatchmakeExtensionProtocol struct {
	server                                          *nex.Server
	OpenParticipationHandler                        func(err error, client *nex.Client, callID uint32, gid uint32)
	AutoMatchmake_PostponeHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string)
	CreateMatchmakeSessionHandler                   func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string, participationCount uint16)
	UpdateNotificationDataHandler                   func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)
	GetFriendNotificationDataHandler                func(err error, client *nex.Client, callID uint32, uiType int32)
	AutoMatchmakeWithSearchCriteria_PostponeHandler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string)
	JoinMatchmakeSessionExHandler                   func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)
	GetSimplePlayingSessionHandler                  func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)
}

// Setup initializes the protocol
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) Setup() {
	nexServer := matchmakeExtensionProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchmakeExtensionProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchmakeExtensionMethodOpenParticipation:
				go matchmakeExtensionProtocol.handleOpenParticipation(packet)
			case MatchmakeExtensionMethodAutoMatchmake_Postpone:
				go matchmakeExtensionProtocol.handleAutoMatchmake_Postpone(packet)
			case MatchmakeExtensionMethodCreateMatchmakeSession:
				go matchmakeExtensionProtocol.handleCreateMatchmakeSession(packet)
			case MatchmakeExtensionMethodUpdateNotificationData:
				go matchmakeExtensionProtocol.handleUpdateNotificationData(packet)
			case MatchmakeExtensionMethodGetFriendNotificationData:
				go matchmakeExtensionProtocol.handleGetFriendNotificationData(packet)
			case MatchmakeExtensionMethodAutoMatchmakeWithSearchCriteria_Postpone:
				go matchmakeExtensionProtocol.handleAutoMatchmakeWithSearchCriteria_Postpone(packet)
			case MatchmakeExtensionMethodJoinMatchmakeSessionEx:
				go matchmakeExtensionProtocol.handleJoinMatchmakeSessionEx(packet)
			case MatchmakeExtensionMethodGetSimplePlayingSession:
				go matchmakeExtensionProtocol.handleGetSimplePlayingSession(packet)
			default:
				go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
				fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID())
			}
		}
	})
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

// UpdateNotificationData sets the UpdateNotificationData handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) UpdateNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType uint32, uiParam1 uint32, uiParam2 uint32, strParam string)) {
	matchmakeExtensionProtocol.UpdateNotificationDataHandler = handler
}

// GetFriendNotificationData sets the GetFriendNotificationData handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetFriendNotificationData(handler func(err error, client *nex.Client, callID uint32, uiType int32)) {
	matchmakeExtensionProtocol.GetFriendNotificationDataHandler = handler
}

// AutoMatchmakeWithSearchCriteria_Postpone sets the AutoMatchmakeWithSearchCriteria_Postpone handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) AutoMatchmakeWithSearchCriteria_Postpone(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string)) {
	matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler = handler
}

// JoinMatchmakeSessionEx sets the JoinMatchmakeSessionEx handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionEx(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)) {
	matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler = handler
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)) {
	matchmakeExtensionProtocol.GetSimplePlayingSessionHandler = handler
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleOpenParticipation(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.OpenParticipationHandler == nil {
		logger.Warning("MatchmakeExtensionProtocol::OpenParticipation not implemented")
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
		logger.Warning("MatchmakeExtensionProtocol::AutoMatchmake_PostponeHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

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

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmake_PostponeHandler(err, client, callID, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, matchmakeExtensionProtocol.server)

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		logger.Error(err.Error())
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

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.CreateMatchmakeSessionHandler == nil {
		logger.Warning("MatchmakeExtensionProtocol::CreateMatchmakeSession not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

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

	message, err := dataHolderContentStream.ReadString()

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
		logger.Warning("MatchmakeExtensionProtocol::UpdateNotificationData not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

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
		logger.Warning("MatchmakeExtensionProtocol::GetFriendNotificationData not implemented")
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

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleAutoMatchmakeWithSearchCriteria_Postpone(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler == nil {
		logger.Warning("MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_PostponeHandler not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()
	logger.Info(hex.EncodeToString(parameters))

	parametersStream := nex.NewStreamIn(parameters, matchmakeExtensionProtocol.server)

	criteriaCount := int(parametersStream.ReadUInt32LE())
	for i := 0; i < criteriaCount; i++ {
		_, _ = parametersStream.ReadStructure(NewMatchmakeSessionSearchCriteria())
	}
	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	logger.Info(dataHolderType)

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_Postpone] Data holder name does not match")
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtensionProtocol::AutoMatchmakeWithSearchCriteria_Postpone] Data holder missing lengths")
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, matchmakeExtensionProtocol.server)

	gatheringStructureInterface, err := dataHolderContentStream.ReadStructure(NewGathering())
	if err != nil {
		logger.Error(err.Error())
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	matchmakeSessionStructureInterface, err := dataHolderContentStream.ReadStructure(NewMatchmakeSession())
	if err != nil {
		logger.Error(err.Error())
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}
	matchmakeSession := matchmakeSessionStructureInterface.(*MatchmakeSession)
	matchmakeSession.Gathering = gatheringStructureInterface.(*Gathering)

	message, err := parametersStream.ReadString()
	if err != nil {
		go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(err, client, callID, nil, "")
		return
	}

	go matchmakeExtensionProtocol.AutoMatchmakeWithSearchCriteria_PostponeHandler(nil, client, callID, matchmakeSession, message)
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
	participationCount := parametersStream.ReadUInt16LE()

	go matchmakeExtensionProtocol.JoinMatchmakeSessionExHandler(nil, client, callID, gid, strMessage, dontCareMyBlockList, participationCount)
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetSimplePlayingSessionHandler == nil {
		logger.Warning("MatchmakeExtensionProtocol::GetSimplePlayingSession not implemented")
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

// NewMatchmakeExtensionProtocol returns a new MatchmakeExtensionProtocol
func NewMatchmakeExtensionProtocol(server *nex.Server) *MatchmakeExtensionProtocol {
	matchmakeExtensionProtocol := &MatchmakeExtensionProtocol{server: server}

	matchmakeExtensionProtocol.Setup()

	return matchmakeExtensionProtocol
}
