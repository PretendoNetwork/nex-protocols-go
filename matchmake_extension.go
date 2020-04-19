package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchmakeExtensionProtocolID is the protocol ID for the Matchmake Extension protocol
	MatchmakeExtensionProtocolID = 0x6D

	// MatchmakeExtensionMethodGetSimplePlayingSession is the method ID for method GetSimplePlayingSession
	MatchmakeExtensionMethodGetSimplePlayingSession = 0x1F
)

// MatchmakeExtensionProtocol handles the Matchmake Extension nex protocol
type MatchmakeExtensionProtocol struct {
	server *nex.Server

	GetSimplePlayingSessionHandler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)
}

// Setup initializes the protocol
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) Setup() {
	nexServer := matchmakeExtensionProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if MatchmakeExtensionProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case MatchmakeExtensionMethodGetSimplePlayingSession:
				go matchmakeExtensionProtocol.handleGetSimplePlayingSession(packet)
			default:
				fmt.Printf("Unsupported Matchmake Extension method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

// GetSimplePlayingSession sets the GetSimplePlayingSession handler function
func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) GetSimplePlayingSession(handler func(err error, client *nex.Client, callID uint32, listPID []uint32, includeLoginUser bool)) {
	matchmakeExtensionProtocol.GetSimplePlayingSessionHandler = handler
}

func (matchmakeExtensionProtocol *MatchmakeExtensionProtocol) handleGetSimplePlayingSession(packet nex.PacketInterface) {
	if matchmakeExtensionProtocol.GetSimplePlayingSessionHandler == nil {
		fmt.Println("[Warning] MatchmakeExtensionProtocol::GetSimplePlayingSession not implemented")
		go respondNotImplemented(packet, MatchmakeExtensionProtocolID)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

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
