package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchMakingProtocolID is the protocol ID for the Match Making protocol
	MatchMakingProtocolID = 0x15

	// MatchMakingMethodGetSessionURLs is the method ID for the method GetSessionURLs
	MatchMakingMethodGetSessionURLs = 0x29
)

// AuthenticationProtocol handles the Authentication nex protocol
type MatchMakingProtocol struct {
	server                *nex.Server
	GetSessionURLsHandler func(err error, client *nex.Client, callID uint32, gatheringId uint32)
}

// Setup initializes the protocol
func (matchMakingProtocol *MatchMakingProtocol) Setup() {
	nexServer := matchMakingProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchMakingProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchMakingMethodGetSessionURLs:
				go matchMakingProtocol.handleGetSessionURLs(packet)
			default:
				go respondNotImplemented(packet, MatchMakingProtocolID)
				fmt.Printf("Unsupported MatchMaking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetSessionURLs sets the GetSessionURLs handler function
func (matchMakingProtocol *MatchMakingProtocol) GetSessionURLs(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	matchMakingProtocol.GetSessionURLsHandler = handler
}

func (matchMakingProtocol *MatchMakingProtocol) handleGetSessionURLs(packet nex.PacketInterface) {
	if matchMakingProtocol.GetSessionURLsHandler == nil {
		logger.Warning("MatchMakingProtocol::GetSessionURLs not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	gatheringId := parametersStream.ReadUInt32LE()

	go matchMakingProtocol.GetSessionURLsHandler(nil, client, callID, gatheringId)
}

// NewMatchMakingProtocol returns a new MatchMakingProtocol
func NewMatchMakingProtocol(server *nex.Server) *MatchMakingProtocol {
	matchMakingProtocol := &MatchMakingProtocol{server: server}

	matchMakingProtocol.Setup()

	return matchMakingProtocol
}
