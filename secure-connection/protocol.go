// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection protocol
	ProtocolID = 0xB

	// MethodRegister is the method ID for the method Register
	MethodRegister = 0x1

	// MethodRequestConnectionData is the method ID for the method RequestConnectionData
	MethodRequestConnectionData = 0x2

	// MethodRequestURLs is the method ID for the method RequestURLs
	MethodRequestURLs = 0x3

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x4

	// MethodTestConnectivity is the method ID for the method TestConnectivity
	MethodTestConnectivity = 0x5

	// MethodUpdateURLs is the method ID for the method UpdateURLs
	MethodUpdateURLs = 0x6

	// MethodReplaceURL is the method ID for the method ReplaceURL
	MethodReplaceURL = 0x7

	// MethodSendReport is the method ID for the method SendReport
	MethodSendReport = 0x8
)

// Protocol stores all the RMC method handlers for the Secure Connection protocol and listens for requests
type Protocol struct {
	server                nex.ServerInterface
	Register              func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)
	RequestConnectionData func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32)
	RequestURLs           func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32)
	RegisterEx            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, uint32)
	TestConnectivity      func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	UpdateURLs            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)
	ReplaceURL            func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, uint32)
	SendReport            func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Secure Connection protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32))
	SetHandlerRequestConnectionData(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerRequestURLs(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32))
	SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, uint32))
	SetHandlerTestConnectivity(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateURLs(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32))
	SetHandlerReplaceURL(handler func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, uint32))
	SetHandlerSendReport(handler func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.Buffer) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerRegister sets the handler for the Register method
func (protocol *Protocol) SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)) {
	protocol.Register = handler
}

// SetHandlerRequestConnectionData sets the handler for the RequestConnectionData method
func (protocol *Protocol) SetHandlerRequestConnectionData(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.RequestConnectionData = handler
}

// SetHandlerRequestURLs sets the handler for the RequestURLs method
func (protocol *Protocol) SetHandlerRequestURLs(handler func(err error, packet nex.PacketInterface, callID uint32, cidTarget *types.PrimitiveU32, pidTarget *types.PID) (*nex.RMCMessage, uint32)) {
	protocol.RequestURLs = handler
}

// SetHandlerRegisterEx sets the handler for the RegisterEx method
func (protocol *Protocol) SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL], hCustomData *types.AnyDataHolder) (*nex.RMCMessage, uint32)) {
	protocol.RegisterEx = handler
}

// SetHandlerTestConnectivity sets the handler for the TestConnectivity method
func (protocol *Protocol) SetHandlerTestConnectivity(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.TestConnectivity = handler
}

// SetHandlerUpdateURLs sets the handler for the UpdateURLs method
func (protocol *Protocol) SetHandlerUpdateURLs(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)) {
	protocol.UpdateURLs = handler
}

// SetHandlerReplaceURL sets the handler for the ReplaceURL method
func (protocol *Protocol) SetHandlerReplaceURL(handler func(err error, packet nex.PacketInterface, callID uint32, target *types.StationURL, url *types.StationURL) (*nex.RMCMessage, uint32)) {
	protocol.ReplaceURL = handler
}

// SetHandlerSendReport sets the handler for the SendReport method
func (protocol *Protocol) SetHandlerSendReport(handler func(err error, packet nex.PacketInterface, callID uint32, reportID *types.PrimitiveU32, reportData *types.QBuffer) (*nex.RMCMessage, uint32)) {
	protocol.SendReport = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	switch message.MethodID {
	case MethodRegister:
		protocol.handleRegister(packet)
	case MethodRequestConnectionData:
		protocol.handleRequestConnectionData(packet)
	case MethodRequestURLs:
		protocol.handleRequestURLs(packet)
	case MethodRegisterEx:
		protocol.handleRegisterEx(packet)
	case MethodTestConnectivity:
		protocol.handleTestConnectivity(packet)
	case MethodUpdateURLs:
		protocol.handleUpdateURLs(packet)
	case MethodReplaceURL:
		protocol.handleReplaceURL(packet)
	case MethodSendReport:
		protocol.handleSendReport(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		fmt.Printf("Unsupported SecureConnection method ID: %#v\n", message.MethodID)
	}
}

// NewProtocol returns a new Secure Connection protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	return &Protocol{server: server}
}
