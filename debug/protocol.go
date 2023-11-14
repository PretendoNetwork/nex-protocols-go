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
	EnableAPIRecorder                func(err error, packet nex.PacketInterface, callID uint32) uint32
	DisableAPIRecorder               func(err error, packet nex.PacketInterface, callID uint32) uint32
	IsAPIRecorderEnabled             func(err error, packet nex.PacketInterface, callID uint32) uint32
	GetAPICalls                      func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown *nex.DateTime, unknown2 *nex.DateTime) uint32
	SetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32 // TODO - Unknown request/response format
	GetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32 // TODO - Unknown request/response format
	GetAPICallSummary                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32 // TODO - Unknown request/response format
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
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
		fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Debug protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
