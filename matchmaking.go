package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchMakingProtocolID is the protocol ID for the Match Making protocol
	MatchMakingProtocolID = 0x15

	// MatchMakingMethodUnregisterGathering is the method ID for the method UnregisterGathering
	MatchMakingMethodUnregisterGathering = 0x2

	// MatchMakingMethodUnregisterGatherings is the method ID for the method UnregisterGatherings
	MatchMakingMethodUnregisterGatherings = 0x3

	// MatchMakingMethodFindBySingleID is the method ID for the method FindBySingleID
	MatchMakingMethodFindBySingleID = 0x15

	// MatchMakingMethodUpdateSessionHostV1 is the method ID for the method UpdateSessionHostV1
	MatchMakingMethodUpdateSessionHostV1 = 0x28

	// MatchMakingMethodGetSessionURLs is the method ID for the method GetSessionURLs
	MatchMakingMethodGetSessionURLs = 0x29

	// MatchMakingMethodUpdateSessionHost is the method ID for the method UpdateSessionHost
	MatchMakingMethodUpdateSessionHost = 0x2A
)

// AuthenticationProtocol handles the Authentication nex protocol
type MatchMakingProtocol struct {
	server                      *nex.Server
	UnregisterGatheringHandler  func(err error, client *nex.Client, callID uint32, idGathering uint32)
	UnregisterGatheringsHandler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)
	FindBySingleIDHandler       func(err error, client *nex.Client, callID uint32, id uint32)
	UpdateSessionHostV1Handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)
	GetSessionURLsHandler       func(err error, client *nex.Client, callID uint32, gatheringId uint32)
	UpdateSessionHostHandler   func(err error, client *nex.Client, callID uint32, gatheringId uint32)
}

// Setup initializes the protocol
func (matchMakingProtocol *MatchMakingProtocol) Setup() {
	nexServer := matchMakingProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchMakingProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchMakingMethodUnregisterGathering:
				go matchMakingProtocol.handleMatchMakingMethodUnregisterGathering(packet)
			case MatchMakingMethodUnregisterGatherings:
				go matchMakingProtocol.handleMatchMakingMethodUnregisterGatherings(packet)
			case MatchMakingMethodFindBySingleID:
				go matchMakingProtocol.handleFindBySingleID(packet)
			case MatchMakingMethodUpdateSessionHostV1:
				go matchMakingProtocol.handleUpdateSessionHostV1(packet)
			case MatchMakingMethodGetSessionURLs:
				go matchMakingProtocol.handleGetSessionURLs(packet)
			case MatchMakingMethodUpdateSessionHost:
				go matchMakingProtocol.handleUpdateSessionHost(packet)
			default:
				go respondNotImplemented(packet, MatchMakingProtocolID)
				fmt.Printf("Unsupported MatchMaking method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// UnregisterGathering sets the UnregisterGathering handler function
func (matchMakingProtocol *MatchMakingProtocol) UnregisterGathering(handler func(err error, client *nex.Client, callID uint32, idGathering uint32)) {
	matchMakingProtocol.UnregisterGatheringHandler = handler
}

// UnregisterGatherings sets the UnregisterGatherings handler function
func (matchMakingProtocol *MatchMakingProtocol) UnregisterGatherings(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	matchMakingProtocol.UnregisterGatheringsHandler = handler
}

// FindBySingleID sets the FindBySingleID handler function
func (matchMakingProtocol *MatchMakingProtocol) FindBySingleID(handler func(err error, client *nex.Client, callID uint32, id uint32)) {
	matchMakingProtocol.FindBySingleIDHandler = handler
}

// GetSessionURLs sets the GetSessionURLs handler function
func (matchMakingProtocol *MatchMakingProtocol) UpdateSessionHostV1(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	matchMakingProtocol.UpdateSessionHostV1Handler = handler
}

// GetSessionURLs sets the GetSessionURLs handler function
func (matchMakingProtocol *MatchMakingProtocol) GetSessionURLs(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	matchMakingProtocol.GetSessionURLsHandler = handler
}

// GetSessionURLs sets the GetSessionURLs handler function
func (matchMakingProtocol *MatchMakingProtocol) UpdateSessionHost(handler func(err error, client *nex.Client, callID uint32, gatheringId uint32)) {
	matchMakingProtocol.UpdateSessionHostHandler = handler
}

func (matchMakingProtocol *MatchMakingProtocol) handleMatchMakingMethodUnregisterGathering(packet nex.PacketInterface) {
	if matchMakingProtocol.UnregisterGatheringHandler == nil {
		logger.Warning("MatchMakingProtocol::UnregisterGathering not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	idGathering := parametersStream.ReadUInt32LE()

	go matchMakingProtocol.UnregisterGatheringHandler(nil, client, callID, idGathering)
}

func (matchMakingProtocol *MatchMakingProtocol) handleMatchMakingMethodUnregisterGatherings(packet nex.PacketInterface) {
	if matchMakingProtocol.UnregisterGatheringsHandler == nil {
		logger.Warning("MatchMakingProtocol::UnregisterGatherings not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	lstGatherings := parametersStream.ReadListUInt32LE()

	go matchMakingProtocol.UnregisterGatheringsHandler(nil, client, callID, lstGatherings)
}

func (matchMakingProtocol *MatchMakingProtocol) handleFindBySingleID(packet nex.PacketInterface) {
	if matchMakingProtocol.FindBySingleIDHandler == nil {
		logger.Warning("MatchMakingProtocol::FindBySingleID not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	id := parametersStream.ReadUInt32LE()

	go matchMakingProtocol.FindBySingleIDHandler(nil, client, callID, id)
}

func (matchMakingProtocol *MatchMakingProtocol) handleUpdateSessionHostV1(packet nex.PacketInterface) {
	if matchMakingProtocol.UpdateSessionHostV1Handler == nil {
		fmt.Println("[Warning] MatchMakingProtocol::UpdateSessionHostV1 not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	gatheringId := parametersStream.ReadUInt32LE()

	go matchMakingProtocol.UpdateSessionHostV1Handler(nil, client, callID, gatheringId)
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

func (matchMakingProtocol *MatchMakingProtocol) handleUpdateSessionHost(packet nex.PacketInterface) {
	if matchMakingProtocol.UpdateSessionHostHandler == nil {
		fmt.Println("[Warning] MatchMakingProtocol::UpdateSessionHost not implemented")
		go respondNotImplemented(packet, MatchMakingProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingProtocol.server)

	gatheringId := parametersStream.ReadUInt32LE()

	go matchMakingProtocol.UpdateSessionHostHandler(nil, client, callID, gatheringId)
}

// NewMatchMakingProtocol returns a new MatchMakingProtocol
func NewMatchMakingProtocol(server *nex.Server) *MatchMakingProtocol {
	matchMakingProtocol := &MatchMakingProtocol{server: server}

	matchMakingProtocol.Setup()

	return matchMakingProtocol
}
