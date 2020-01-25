package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	AuthenticationProtocolID = 0xA

	AuthenticationMethodLogin          = 0x1
	AuthenticationMethodLoginEx        = 0x2
	AuthenticationMethodRequestTicket  = 0x3
	AuthenticationMethodGetPID         = 0x4
	AuthenticationMethodGetName        = 0x5
	AuthenticationMethodLoginWithParam = 0x6
)

type NintendoLoginData struct {
	token string
}

type AuthenticationInfo struct {
	token         string
	tokenType     uint32
	ngsVersion    uint8
	serverVersion uint32
}

type AuthenticationProtocol struct {
	server                *nex.Server
	LoginHandler          func(client *nex.Client, callID uint32, username string)
	LoginExHandler        func(client *nex.Client, callID uint32, username string, authenticationInfo AuthenticationInfo)
	RequestTicketHandler  func(client *nex.Client, callID uint32, userPID uint32, serverPID uint32)
	GetPIDHandler         func(client *nex.Client, callID uint32, username string)
	GetNameHandler        func(client *nex.Client, callID uint32, userPID uint32)
	LoginWithParamHandler func(client *nex.Client, callID uint32)
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
			case AuthenticationMethodGetPID:
				go authenticationProtocol.handleGetPID(packet)
			case AuthenticationMethodGetName:
				go authenticationProtocol.handleGetName(packet)
			case AuthenticationMethodLoginWithParam:
				go authenticationProtocol.handleLoginWithParam(packet)
			default:
				fmt.Printf("Unsupported Authentication method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (authenticationProtocol *AuthenticationProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(AuthenticationProtocolID, request.GetCallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.GetVersion() == 1 {
		responsePacket = nex.NewPacketV0(client, nil)
	} else {
		responsePacket = nex.NewPacketV1(client, nil)
	}

	responsePacket.SetVersion(packet.GetVersion())
	responsePacket.SetSource(packet.GetDestination())
	responsePacket.SetDestination(packet.GetSource())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)

	authenticationProtocol.server.Send(responsePacket)
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

func (authenticationProtocol *AuthenticationProtocol) GetPID(handler func(client *nex.Client, callID uint32, username string)) {
	authenticationProtocol.GetPIDHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) GetName(handler func(client *nex.Client, callID uint32, userPID uint32)) {
	authenticationProtocol.GetNameHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) LoginWithParam(handler func(client *nex.Client, callID uint32)) {
	authenticationProtocol.LoginWithParamHandler = handler
}

func (authenticationProtocol *AuthenticationProtocol) handleLogin(packet nex.PacketInterface) {
	if authenticationProtocol.LoginHandler == nil {
		fmt.Println("[Warning] AuthenticationProtocol::Login not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
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
		fmt.Println("[Warning] AuthenticationProtocol::LoginEx not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
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
		fmt.Println("[Warning] AuthenticationProtocol::RequestTicket not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
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

func (authenticationProtocol *AuthenticationProtocol) handleGetPID(packet nex.PacketInterface) {
	if authenticationProtocol.GetPIDHandler == nil {
		fmt.Println("[Warning] AuthenticationProtocol::GetPID not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	username := parametersStream.ReadNEXStringNext()

	go authenticationProtocol.GetPIDHandler(client, callID, username)
}

func (authenticationProtocol *AuthenticationProtocol) handleGetName(packet nex.PacketInterface) {
	if authenticationProtocol.GetNameHandler == nil {
		fmt.Println("[Warning] AuthenticationProtocol::GetName not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	userPID := parametersStream.ReadU32LENext(1)[0]

	go authenticationProtocol.GetNameHandler(client, callID, userPID)
}

func (authenticationProtocol *AuthenticationProtocol) handleLoginWithParam(packet nex.PacketInterface) {
	if authenticationProtocol.LoginWithParamHandler == nil {
		fmt.Println("[Warning] AuthenticationProtocol::LoginWithParam not implemented")
		go authenticationProtocol.respondNotImplemented(packet)
		return
	}

	// Unsure what data is sent here, or how to trigger the console to send it
}

func NewAuthenticationProtocol(server *nex.Server) *AuthenticationProtocol {
	authenticationProtocol := &AuthenticationProtocol{server: server}

	authenticationProtocol.Setup()

	return authenticationProtocol
}
