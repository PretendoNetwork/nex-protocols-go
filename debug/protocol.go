// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Debug protocol
	ProtocolID = 0x74

	// MethodEnableAPIRecorder is the method ID for the method EnableAPIRecorder
	MethodEnableAPIRecorder = 0x1

	// MethodDisableAPIRecorder is the method ID for the method DisableAPIRecorder
	MethodDisableAPIRecorder = 0x2

	// MethodIsAPIRecorderEnabled is the method ID for the method IsAPIRecorderEnabled
	MethodIsAPIRecorderEnabled = 0x3

	// MethodGetAPICalls is the method ID for the method GetAPICalls
	MethodGetAPICalls = 0x4

	// MethodSetExcludeJoinedMatchmakeSession is the method ID for the method SetExcludeJoinedMatchmakeSession
	MethodSetExcludeJoinedMatchmakeSession = 0x5

	// MethodGetExcludeJoinedMatchmakeSession is the method ID for the method GetExcludeJoinedMatchmakeSession
	MethodGetExcludeJoinedMatchmakeSession = 0x6

	// MethodGetAPICallSummary is the method ID for the method GetAPICallSummary
	MethodGetAPICallSummary = 0x7
)

// Protocol handles the Debug protocol
type Protocol struct {
	Server                                  *nex.Server
	enableAPIRecorderHandler                func(err error, client *nex.Client, callID uint32)
	disableAPIRecorderHandler               func(err error, client *nex.Client, callID uint32)
	isAPIRecorderEnabledHandler             func(err error, client *nex.Client, callID uint32)
	getAPICallsHandler                      func(err error, client *nex.Client, callID uint32, pids []uint32, unknown *nex.DateTime, unknown2 *nex.DateTime)
	setExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	getExcludeJoinedMatchmakeSessionHandler func(err error, client *nex.Client, callID uint32)
	getAPICallSummaryHandler                func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodEnableAPIRecorder:
		go protocol.handleEnableAPIRecorder(packet)
	case MethodDisableAPIRecorder:
		go protocol.handleDisableAPIRecorder(packet)
	case MethodIsAPIRecorderEnabled:
		go protocol.handleIsAPIRecorderEnabled(packet)
	case MethodGetAPICalls:
		go protocol.handleGetAPICalls(packet)
	case MethodSetExcludeJoinedMatchmakeSession:
		go protocol.handleSetExcludeJoinedMatchmakeSession(packet)
	case MethodGetExcludeJoinedMatchmakeSession:
		go protocol.handleGetExcludeJoinedMatchmakeSession(packet)
	case MethodGetAPICallSummary:
		go protocol.handleGetAPICallSummary(packet)
	default:
		fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Debug protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
