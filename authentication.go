package nexproto

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

const (
	AuthenticationProtocolID = 0xA

	AuthenticationMethodLogin = 0x1
	AuthenticationMethodLoginEx = 0x2
	AuthenticationMethodRequestTicket = 0x3
	AuthenticationMethodGetPID = 0x4
	AuthenticationMethodGetName = 0x5
	AuthenticationMethodLoginWithParam = 0x6
)

type NintendoLoginData struct {
	token string
}

type AuthenticationProtocol struct {
	server *nex.Server
	loginHandler func(client *nex.Client, callID uint32, username string)
	requestTicketHandler func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32)
}

func (authenticationProtocol *AuthenticationProtocol) Setup() {
	nexServer := authenticationProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if AuthenticationProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case AuthenticationMethodLogin:
				authenticationProtocol.handleLogin(packet)
			case AuthenticationMethodRequestTicket:
				authenticationProtocol.handleRequestTicket(packet)
			default:
				fmt.Printf("Unsupported Authentication method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (authenticationProtocol *AuthenticationProtocol) Login(handler func(client *nex.Client, callID uint32, username string)) {
	authenticationProtocol.loginHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) RequestTicket(handler func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32)) {
	authenticationProtocol.requestTicketHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) handleLogin(packet nex.PacketInterface) {
	if authenticationProtocol.loginHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	username := parametersStream.ReadNEXStringNext()

	authenticationProtocol.loginHandler(client, callID, username)
}

func (authenticationProtocol *AuthenticationProtocol) handleRequestTicket(packet nex.PacketInterface) {
	if authenticationProtocol.requestTicketHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	userPID := parametersStream.ReadU32LENext(1)[0]
	serverPID := parametersStream.ReadU32LENext(1)[0]

	authenticationProtocol.requestTicketHandler(client, callID, userPID, serverPID)
}

func NewAuthenticationProtocol(server *nex.Server) *AuthenticationProtocol {
	authenticationProtocol := &AuthenticationProtocol{server: server}

	authenticationProtocol.Setup()

	return authenticationProtocol
}