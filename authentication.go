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

type AuthenticationInfo struct {
	token string
	tokenType uint32
	ngsVersion uint8
	serverVersion uint32
}

type AuthenticationProtocol struct {
	server *nex.Server
	LoginHandler func(client *nex.Client, callID uint32, username string)
	LoginExHandler func(client *nex.Client, callID uint32, username string, authenticationInfo AuthenticationInfo)
	RequestTicketHandler func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32)
}

func (authenticationProtocol *AuthenticationProtocol) Setup() {
	nexServer := authenticationProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if AuthenticationProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case AuthenticationMethodLogin:
				go authenticationProtocol.handleLogin(packet)
			case AuthenticationMethodLoginEx:
				go authenticationProtocol.handleLoginEx(packet)
			case AuthenticationMethodRequestTicket:
				go authenticationProtocol.handleRequestTicket(packet)
			default:
				fmt.Printf("Unsupported Authentication method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (authenticationProtocol *AuthenticationProtocol) Login(handler func(client *nex.Client, callID uint32, username string)) {
	authenticationProtocol.LoginHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) LoginEx(handler func(client *nex.Client, callID uint32, username string, authenticationInfo AuthenticationInfo)) {
	authenticationProtocol.LoginExHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) RequestTicket(handler func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32)) {
	authenticationProtocol.RequestTicketHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) handleLogin(packet nex.PacketInterface) {
	if authenticationProtocol.LoginHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	username := parametersStream.ReadNEXStringNext()

	go authenticationProtocol.LoginHandler(client, callID, username)
}

func (authenticationProtocol *AuthenticationProtocol) handleLoginEx(packet nex.PacketInterface) {
	if authenticationProtocol.LoginExHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	username := parametersStream.ReadNEXStringNext()
	authenticationInfo := AuthenticationInfo{}

	go authenticationProtocol.LoginExHandler(client, callID, username, authenticationInfo)
}

func (authenticationProtocol *AuthenticationProtocol) handleRequestTicket(packet nex.PacketInterface) {
	if authenticationProtocol.RequestTicketHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	userPID := parametersStream.ReadU32LENext(1)[0]
	serverPID := parametersStream.ReadU32LENext(1)[0]

	go authenticationProtocol.RequestTicketHandler(client, callID, userPID, serverPID)
}

func NewAuthenticationProtocol(server *nex.Server) *AuthenticationProtocol {
	authenticationProtocol := &AuthenticationProtocol{server: server}

	authenticationProtocol.Setup()

	return authenticationProtocol
}