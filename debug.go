package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DebugProtocolID is the protocol ID for the Debug protocol
	DebugProtocolID = 0x74

	// DebugMethodEnableApiRecorder is the method ID for the method EnableApiRecorder
	DebugMethodEnableApiRecorder = 0x1

	// DebugMethodDisableApiRecorder is the method ID for the method DisableApiRecorder
	DebugMethodDisableApiRecorder = 0x2

	// DebugMethodIsApiRecorderEnabled is the method ID for the method IsApiRecorderEnabled
	DebugMethodIsApiRecorderEnabled = 0x3

	// DebugMethodGetApiCalls is the method ID for the method GetApiCalls
	DebugMethodGetApiCalls = 0x4

	// DebugMethodSetExcludeJoinedMatchmakeSession is the method ID for the method SetExcludeJoinedMatchmakeSession
	DebugMethodSetExcludeJoinedMatchmakeSession = 0x5

	// DebugMethodGetExcludeJoinedMatchmakeSession is the method ID for the method GetExcludeJoinedMatchmakeSession
	DebugMethodGetExcludeJoinedMatchmakeSession = 0x6

	// DebugMethodGetApiCallSummary is the method ID for the method GetApiCallSummary
	DebugMethodGetApiCallSummary = 0x7
)

// DebugProtocol handles the Debug protocol
type DebugProtocol struct {
	server                                  *nex.Server
	EnableApiRecorderHandler                func(err error, client *nex.Client, callID uint32)
	DisableApiRecorderHandler               func(err error, client *nex.Client, callID uint32)
	IsApiRecorderEnabledHandler             func(err error, client *nex.Client, callID uint32)
	GetApiCallsHandler                      func(err error, client *nex.Client, callID uint32, pids []uint32, dateUnk1 uint64, dateUnk2 uint64)
	SetExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	GetExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	GetApiCallSummaryHandler                func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (debugProtocol *DebugProtocol) Setup() {
	nexServer := debugProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DebugProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DebugMethodEnableApiRecorder:
				go debugProtocol.handleEnableApiRecorder(packet)
				break
			case DebugMethodDisableApiRecorder:
				go debugProtocol.handleDisableApiRecorder(packet)
				break
			case DebugMethodIsApiRecorderEnabled:
				go debugProtocol.handleIsApiRecorderEnabled(packet)
				break
			case DebugMethodGetApiCalls:
				go debugProtocol.handleGetApiCalls(packet)
				break
			case DebugMethodSetExcludeJoinedMatchmakeSession:
				go debugProtocol.handleSetExcludeJoinedMatchmakeSession(packet)
				break
			case DebugMethodGetExcludeJoinedMatchmakeSession:
				go debugProtocol.handleGetExcludeJoinedMatchmakeSession(packet)
				break
			case DebugMethodGetApiCallSummary:
				go debugProtocol.handleGetApiCallSummary(packet)
				break
			default:
				fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (debugProtocol *DebugProtocol) handleEnableApiRecorder(packet nex.PacketInterface) {
	if debugProtocol.EnableApiRecorderHandler == nil {
		logger.Warning("DebugProtocol::EnableApiRecorder not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go debugProtocol.EnableApiRecorderHandler(nil, client, callID)
}

func (debugProtocol *DebugProtocol) handleDisableApiRecorder(packet nex.PacketInterface) {
	if debugProtocol.DisableApiRecorderHandler == nil {
		logger.Warning("DebugProtocol::DisableApiRecorder not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go debugProtocol.DisableApiRecorderHandler(nil, client, callID)
}

func (debugProtocol *DebugProtocol) handleIsApiRecorderEnabled(packet nex.PacketInterface) {
	if debugProtocol.IsApiRecorderEnabledHandler == nil {
		logger.Warning("DebugProtocol::IsApiRecorderEnabled not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go debugProtocol.IsApiRecorderEnabledHandler(nil, client, callID)
}

func (debugProtocol *DebugProtocol) handleGetApiCalls(packet nex.PacketInterface) {
	if debugProtocol.GetApiCallsHandler == nil {
		logger.Warning("DebugProtocol::GetApiCalls not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, debugProtocol.server)

	pidsCount := parametersStream.ReadUInt32LE()
	pids := make([]uint32, pidsCount)
	for i := 0; uint32(i) < pidsCount; i++ {
		pids[i] = parametersStream.ReadUInt32LE()
	}

	dateUnk1 := parametersStream.ReadUInt64LE()

	dateUnk2 := parametersStream.ReadUInt64LE()

	go debugProtocol.GetApiCallsHandler(nil, client, callID, pids, dateUnk1, dateUnk2)
}

func (debugProtocol *DebugProtocol) handleSetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	logger.Warning("DebugProtocol::SetExcludeJoinedMatchmakeSession STUBBED")

	if debugProtocol.SetExcludeJoinedMatchmakeSessionHandler == nil {
		logger.Warning("DebugProtocol::SetExcludeJoinedMatchmakeSession not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

}

func (debugProtocol *DebugProtocol) handleGetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	logger.Warning("DebugProtocol::GetExcludeJoinedMatchmakeSession STUBBED")

	if debugProtocol.GetExcludeJoinedMatchmakeSessionHandler == nil {
		logger.Warning("DebugProtocol::GetExcludeJoinedMatchmakeSession not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

}

func (debugProtocol *DebugProtocol) handleGetApiCallSummary(packet nex.PacketInterface) {
	logger.Warning("DebugProtocol::GetApiCallSummary STUBBED")

	if debugProtocol.GetApiCallSummaryHandler == nil {
		logger.Warning("DebugProtocol::GetApiCallSummary not implemented")
		go respondNotImplemented(packet, DebugProtocolID)
		return
	}

}

// EnableApiRecorder sets the EnableApiRecorder handler function
func (debugProtocol *DebugProtocol) EnableApiRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.EnableApiRecorderHandler = handler
}

// DisableApiRecorder sets the DisableApiRecorder handler function
func (debugProtocol *DebugProtocol) DisableApiRecorder(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.DisableApiRecorderHandler = handler
}

// IsApiRecorderEnabled sets the IsApiRecorderEnabled handler function
func (debugProtocol *DebugProtocol) IsApiRecorderEnabled(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.IsApiRecorderEnabledHandler = handler
}

// GetApiCalls sets the GetApiCalls handler function
func (debugProtocol *DebugProtocol) GetApiCalls(handler func(err error, client *nex.Client, callID uint32, pids []uint32, dateUnk1 uint64, dateUnk2 uint64)) {
	debugProtocol.GetApiCallsHandler = handler
}

// SetExcludeJoinedMatchmakeSession sets the SetExcludeJoinedMatchmakeSession handler function
func (debugProtocol *DebugProtocol) SetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.SetExcludeJoinedMatchmakeSessionHandler = handler
}

// GetExcludeJoinedMatchmakeSession sets the GetExcludeJoinedMatchmakeSession handler function
func (debugProtocol *DebugProtocol) GetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.GetExcludeJoinedMatchmakeSessionHandler = handler
}

// GetApiCallSummary sets the GetApiCallSummary handler function
func (debugProtocol *DebugProtocol) GetApiCallSummary(handler func(err error, client *nex.Client, callID uint32)) {
	debugProtocol.GetApiCallSummaryHandler = handler
}

// NewDebugProtocol returns a new DebugProtocol
func NewDebugProtocol(server *nex.Server) *DebugProtocol {
	debugProtocol := &DebugProtocol{server: server}

	debugProtocol.Setup()

	return debugProtocol
}
