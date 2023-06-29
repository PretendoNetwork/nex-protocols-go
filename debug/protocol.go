package debug

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Debug protocol
	ProtocolID = 0x74

	// MethodEnableApiRecorder is the method ID for the method EnableApiRecorder
	MethodEnableApiRecorder = 0x1

	// MethodDisableApiRecorder is the method ID for the method DisableApiRecorder
	MethodDisableApiRecorder = 0x2

	// MethodIsApiRecorderEnabled is the method ID for the method IsApiRecorderEnabled
	MethodIsApiRecorderEnabled = 0x3

	// MethodGetApiCalls is the method ID for the method GetApiCalls
	MethodGetApiCalls = 0x4

	// MethodSetExcludeJoinedMatchmakeSession is the method ID for the method SetExcludeJoinedMatchmakeSession
	MethodSetExcludeJoinedMatchmakeSession = 0x5

	// MethodGetExcludeJoinedMatchmakeSession is the method ID for the method GetExcludeJoinedMatchmakeSession
	MethodGetExcludeJoinedMatchmakeSession = 0x6

	// MethodGetApiCallSummary is the method ID for the method GetApiCallSummary
	MethodGetApiCallSummary = 0x7
)

// DebugProtocol handles the Debug protocol
type DebugProtocol struct {
	Server                                  *nex.Server
	EnableApiRecorderHandler                func(err error, client *nex.Client, callID uint32)
	DisableApiRecorderHandler               func(err error, client *nex.Client, callID uint32)
	IsApiRecorderEnabledHandler             func(err error, client *nex.Client, callID uint32)
	GetApiCallsHandler                      func(err error, client *nex.Client, callID uint32, pids []uint32, unknown *nex.DateTime, unknown2 *nex.DateTime)
	SetExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	GetExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	GetApiCallSummaryHandler                func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *DebugProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *DebugProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodEnableApiRecorder:
		go protocol.handleEnableApiRecorder(packet)
	case MethodDisableApiRecorder:
		go protocol.handleDisableApiRecorder(packet)
	case MethodIsApiRecorderEnabled:
		go protocol.handleIsApiRecorderEnabled(packet)
	case MethodGetApiCalls:
		go protocol.handleGetApiCalls(packet)
	case MethodSetExcludeJoinedMatchmakeSession:
		go protocol.handleSetExcludeJoinedMatchmakeSession(packet)
	case MethodGetExcludeJoinedMatchmakeSession:
		go protocol.handleGetExcludeJoinedMatchmakeSession(packet)
	case MethodGetApiCallSummary:
		go protocol.handleGetApiCallSummary(packet)
	default:
		fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID())
	}
}

// NewDebugProtocol returns a new DebugProtocol
func NewDebugProtocol(server *nex.Server) *DebugProtocol {
	protocol := &DebugProtocol{Server: server}

	protocol.Setup()

	return protocol
}
