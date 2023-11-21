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
	Server                           nex.ServerInterface
	EnableAPIRecorder                func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	DisableAPIRecorder               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	IsAPIRecorderEnabled             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetAPICalls                      func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown *nex.DateTime, unknown2 *nex.DateTime) (*nex.RMCMessage, uint32)
	SetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
	GetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
	GetAPICallSummary                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodEnableAPIRecorder:
		protocol.handleEnableAPIRecorder(packet)
	case MethodDisableAPIRecorder:
		protocol.handleDisableAPIRecorder(packet)
	case MethodIsAPIRecorderEnabled:
		protocol.handleIsAPIRecorderEnabled(packet)
	case MethodGetAPICalls:
		protocol.handleGetAPICalls(packet)
	case MethodSetExcludeJoinedMatchmakeSession:
		protocol.handleSetExcludeJoinedMatchmakeSession(packet)
	case MethodGetExcludeJoinedMatchmakeSession:
		protocol.handleGetExcludeJoinedMatchmakeSession(packet)
	case MethodGetAPICallSummary:
		protocol.handleGetAPICallSummary(packet)
	default:
		fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Debug protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
