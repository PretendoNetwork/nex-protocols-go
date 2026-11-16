// Package protocol implements the Simple Authentcation protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Simple Authentication protocol
	ProtocolID = 0x74

	// MethodAuthenticate is the method ID for the method Authenticate
	MethodAuthenticate = 0x1

	// MethodLoginWithToken is the method Id for the method LoginWithToken
	MethodLoginWithToken = 0x2

	// MethodLoginWithTokenEx is the method ID for the method LoginWithToken
	MethodLoginWithTokenEx = 0x3

	// MethodLogin is the method ID for the method Login
	MethodLogin = 0x4

	// MethodLoginWithSubAccount is the method ID for the method LoginWithSubAccount
	MethodLoginWithSubAccount = 0x5

	// MethodRegister is the method ID for the method RegisterEx
	MethodRegister = 0x6

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x7

	// MethodLoginWithTokenCafe is the method Id for the method LoginWithTokenCafe
	MethodLoginWithTokenCafe = 0x8

	// MethodLoginWithTokenCafeEx is the method ID for the method LoginWithTokenCafeEx
	MethodLoginWithTokenCafeEx = 0x9
)

// Protocol stores all the RMC method handlers for the Simple Authentication protocol and listens for requests
type Protocol struct {
	endpoint             nex.EndpointInterface
	Authenticate         func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String) (*nex.RMCMessage, *nex.Error)
	LoginWithToken       func(err error, packet nex.PacketInterface, callID uint32, strToken types.String) (*nex.RMCMessage, *nex.Error)
	LoginWithTokenEx     func(err error, packet nex.PacketInterface, callID uint32, strToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error)
	Login                func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String, strPassword types.String) (*nex.RMCMessage, *nex.Error)
	LoginWithSubAccount  func(err error, packet nex.PacketInterface, callId uint32, loginData types.DataHolder) (*nex.RMCMessage, *nex.Error)
	Register             func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error)
	RegisterEx           func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL], hCustomData types.DataHolder) (*nex.RMCMessage, *nex.Error)
	LoginWithTokenCafe   func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String) (*nex.RMCMessage, *nex.Error)
	LoginWithTokenCafeEx func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error)
	Patches              nex.ServiceProtocol
	PatchedMethods       []uint32
}

// Interface implements the methods present on the Simple Authentication protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerAuthenticate(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithToken(handler func(err error, packet nex.PacketInterface, callID uint32, strToken types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithTokenEx(handler func(err error, packet nex.PacketInterface, callID uint32, strToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String, strPassword types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithSubAccount(handler func(err error, packet nex.PacketInterface, callId uint32, loginData types.DataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error))
	SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL], hCustomData types.DataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithTokenCafe(handler func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerLoginWithTokenCafeEx(handler func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerAuthenticate sets the handler for the Authenticate method
func (protocol *Protocol) SetHandlerAuthenticate(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Authenticate = handler
}

// SetHandlerLoginWithToken sets the handler for the LoginWithToken method
func (protocol *Protocol) SetHandlerLoginWithToken(handler func(err error, packet nex.PacketInterface, callID uint32, strToken types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithToken = handler
}

// SetHandlerLoginWithTokenEx sets the handler for the LoginWithTokenEx method
func (protocol *Protocol) SetHandlerLoginWithTokenEx(handler func(err error, packet nex.PacketInterface, callID uint32, strToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithTokenEx = handler
}

// SetHandlerLogin sets the handler for the Login method
func (protocol *Protocol) SetHandlerLogin(handler func(err error, packet nex.PacketInterface, callID uint32, strUserName types.String, strPassword types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.Login = handler
}

// SetHandlerLoginWithSubAccount sets the handler for the LoginWithSubAccount method
func (protocol *Protocol) SetHandlerLoginWithSubAccount(handler func(err error, packet nex.PacketInterface, callId uint32, loginData types.DataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithSubAccount = handler
}

// SetHandlerRegister sets the handler for the Register method
func (protocol *Protocol) SetHandlerRegister(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL]) (*nex.RMCMessage, *nex.Error)) {
	protocol.Register = handler
}

// SetHandlerRegisterEx sets the handler for the RegisterEx method
func (protocol *Protocol) SetHandlerRegisterEx(handler func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs types.List[types.StationURL], hCustomData types.DataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.RegisterEx = handler
}

// SetHandlerLoginWithTokenCafe sets the handler for the LoginWithTokenCafe method
func (protocol *Protocol) SetHandlerLoginWithTokenCafe(handler func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithTokenCafe = handler
}

// SetHandlerLoginWithTokenCafeEx sets the handler for the LoginWithTokenCafeEx method
func (protocol *Protocol) SetHandlerLoginWithTokenCafeEx(handler func(err error, packet nex.PacketInterface, callID uint32, strNintendoToken types.String, oAnyData types.DataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.LoginWithTokenCafeEx = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodAuthenticate:
		protocol.handleAuthenticate(packet)
	case MethodLoginWithToken:
		protocol.handleLoginWithToken(packet)
	case MethodLoginWithTokenEx:
		protocol.handleLoginWithTokenEx(packet)
	case MethodLogin:
		protocol.handleLogin(packet)
	case MethodLoginWithSubAccount:
		protocol.handleLoginWithSubAccount(packet)
	case MethodRegister:
		protocol.handleRegister(packet)
	case MethodRegisterEx:
		protocol.handleRegisterEx(packet)
	case MethodLoginWithTokenCafe:
		protocol.handleLoginWithTokenCafe(packet)
	case MethodLoginWithTokenCafeEx:
		protocol.handleLoginWithTokenCafeEx(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported SimpleAuthentication method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Simple Authentication protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
