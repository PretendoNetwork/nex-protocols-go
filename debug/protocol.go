// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
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
	server                           nex.ServerInterface
	EnableAPIRecorder                func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	DisableAPIRecorder               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	IsAPIRecorderEnabled             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetAPICalls                      func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown *nex.DateTime, unknown2 *nex.DateTime) (*nex.RMCMessage, uint32)
	SetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
	GetExcludeJoinedMatchmakeSession func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
	GetAPICallSummary                func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Interface implements the methods present on the Debug protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerEnableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerDisableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerIsAPIRecorderEnabled(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetAPICalls(handler func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown *nex.DateTime, unknown2 *nex.DateTime) (*nex.RMCMessage, uint32))
	SetHandlerSetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
	SetHandlerGetAPICallSummary(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerEnableAPIRecorder sets the handler for the EnableAPIRecorder method
func (protocol *Protocol) SetHandlerEnableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.EnableAPIRecorder = handler
}

// SetHandlerDisableAPIRecorder sets the handler for the DisableAPIRecorder method
func (protocol *Protocol) SetHandlerDisableAPIRecorder(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.DisableAPIRecorder = handler
}

// SetHandlerIsAPIRecorderEnabled sets the handler for the IsAPIRecorderEnabled method
func (protocol *Protocol) SetHandlerIsAPIRecorderEnabled(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.IsAPIRecorderEnabled = handler
}

// SetHandlerGetAPICalls sets the handler for the GetAPICalls method
func (protocol *Protocol) SetHandlerGetAPICalls(handler func(err error, packet nex.PacketInterface, callID uint32, pids []*nex.PID, unknown *nex.DateTime, unknown2 *nex.DateTime) (*nex.RMCMessage, uint32)) {
	protocol.GetAPICalls = handler
}

// SetHandlerSetExcludeJoinedMatchmakeSession sets the handler for the SetExcludeJoinedMatchmakeSession method
func (protocol *Protocol) SetHandlerSetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.SetExcludeJoinedMatchmakeSession = handler
}

// SetHandlerGetExcludeJoinedMatchmakeSession sets the handler for the GetExcludeJoinedMatchmakeSession method
func (protocol *Protocol) SetHandlerGetExcludeJoinedMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetExcludeJoinedMatchmakeSession = handler
}

// SetHandlerGetAPICallSummary sets the handler for the GetAPICallSummary method
func (protocol *Protocol) SetHandlerGetAPICallSummary(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32)) {
	protocol.GetAPICallSummary = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
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
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported Debug method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Debug protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
