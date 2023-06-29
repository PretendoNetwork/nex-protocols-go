package authentication

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Authentication protocol
	ProtocolID = 0xA

	// MethodLogin is the method ID for the method Login
	MethodLogin = 0x1

	// MethodLoginEx is the method ID for the method LoginEx
	MethodLoginEx = 0x2

	// MethodRequestTicket is the method ID for the method RequestTicket
	MethodRequestTicket = 0x3

	// MethodGetPID is the method ID for the method GetPID
	MethodGetPID = 0x4

	// MethodGetName is the method ID for the method GetName
	MethodGetName = 0x5

	// MethodLoginWithParam is the method ID for the method LoginWithParam
	MethodLoginWithParam = 0x6
)

// AuthenticationProtocol handles the Authentication nex protocol
type AuthenticationProtocol struct {
	Server                *nex.Server
	LoginHandler          func(err error, client *nex.Client, callID uint32, strUserName string)
	LoginExHandler        func(err error, client *nex.Client, callID uint32, strUserName string, oExtraData *nex.DataHolder)
	RequestTicketHandler  func(err error, client *nex.Client, callID uint32, idSource uint32, idTarget uint32)
	GetPIDHandler         func(err error, client *nex.Client, callID uint32, strUserName string)
	GetNameHandler        func(err error, client *nex.Client, callID uint32, id uint32)
	LoginWithParamHandler func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *AuthenticationProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket handles in incoming PRUDP packet
func (protocol *AuthenticationProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodLogin:
		go protocol.handleLogin(packet)
	case MethodLoginEx:
		go protocol.handleLoginEx(packet)
	case MethodRequestTicket:
		go protocol.handleRequestTicket(packet)
	case MethodGetPID:
		go protocol.handleGetPID(packet)
	case MethodGetName:
		go protocol.handleGetName(packet)
	case MethodLoginWithParam:
		go protocol.handleLoginWithParam(packet)
	default:
		go globals.RespondNotImplemented(packet, ProtocolID)
		fmt.Printf("Unsupported Authentication method ID: %#v\n", request.MethodID())
	}
}

// NewAuthenticationProtocol returns a new AuthenticationProtocol
func NewAuthenticationProtocol(server *nex.Server) *AuthenticationProtocol {
	protocol := &AuthenticationProtocol{Server: server}

	protocol.Setup()

	return protocol
}
