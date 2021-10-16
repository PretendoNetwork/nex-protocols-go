package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchmakeExtensionProtocolID is the protocol ID for the Matchmake Extension protocol
	MatchmakeExtensionProtocolID = 0x6D

	// MatchmakeExtensionMethodCreateMatchmakeSession is the method ID for method CreateMatchmakeSession
	MatchmakeExtensionMethodCreateMatchmakeSession = 0x6

	// MatchmakeExtensionMethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MatchmakeExtensionMethodGetSimplePlayingSession = 0x1F
)

// MatchmakeExtensionProtocol handles the Matchmake Extension nex protocol
type MatchmakeExtensionProtocol struct {
	server                         *nex.Server
	CreateMatchmakeSessionHandler  func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string, participationCount uint16)
	GetSimplePlayingSessionHandler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)
}

// Setup initializes the protocol
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) Setup() {
	nexServer := matchmakeExtensionProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchmakeExtensionProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchmakeExtensionMethodCreateMatchmakeSession:
				go matchmakeExtensionProtocol.handleCreateMatchmakeSession(packet)
			case MatchmakeExtensionMethodGetSimplePlayingSession:
				go matchmakeExtensionProtocol.handleGetSimplePlayingSession(packet)
			default:
				go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
				fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// CreateMatchmakeSession sets the CreateMatchmakeSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) CreateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *MatchmakeSession, message string, participationCount uint16)) {
	matchmakeExtensionProtocol.CreateMatchmakeSessionHandler = handler
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)) {
	matchmakeExtensionProtocol.GetSimplePlayingSessionHandler = handler
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

// NewMatchmakeExtensionProtocol returns a new MatchmakeExtensionProtocol
func NewMatchmakeExtensionProtocol(server *nex.Server) *MatchmakeExtensionProtocol {
	matchmakeExtensionProtocol := &MatchmakeExtensionProtocol{server: server}

	matchmakeExtensionProtocol.Setup()

	return matchmakeExtensionProtocol
}
