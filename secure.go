package nexproto

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

const (
	SecureProtocolID = 0xB

	SecureMethodRegister = 0x1
	SecureMethodRequestConnectionData = 0x2
	SecureMethodRequestURLs = 0x3
	SecureMethodRegisterEx = 0x4
	SecureMethodTestConnectivity = 0x5
	SecureMethodReplaceURL = 0x6
	SecureMethodSendReport = 0x7
)

type SecureProtocol struct {
	server *nex.Server
	ConnectionIDCounter *nex.Counter
	RegisterHandler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)
	RegisterExHandler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData NintendoLoginData)
}

func (secureProtocol *SecureProtocol) Setup() {
	nexServer := secureProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if SecureProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case SecureMethodRegister:
				secureProtocol.handleRegister(packet)
			case SecureMethodRegisterEx:
				secureProtocol.handleRegisterEx(packet)
			default:
				fmt.Printf("Unsupported Secure method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (secureProtocol *SecureProtocol) Register(handler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)) {
	secureProtocol.RegisterHandler = handler
}

func (secureProtocol *SecureProtocol) RegisterEx(handler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData NintendoLoginData)) {
	secureProtocol.RegisterExHandler = handler
}

func (secureProtocol *SecureProtocol) handleRegister(packet nex.PacketInterface) {
	if secureProtocol.RegisterHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	stationURLCount := parametersStream.ReadU32LENext(1)[0]
	stationUrls := make([]*nex.StationURL, 0)

	for i := 0; i < int(stationURLCount); i++ {
		station := nex.NewStationURL(parametersStream.ReadNEXStringNext())
		stationUrls = append(stationUrls, station)
	}

	secureProtocol.RegisterHandler(client, callID, stationUrls)
}

func (secureProtocol *SecureProtocol) handleRegisterEx(packet nex.PacketInterface) {
	if secureProtocol.RegisterExHandler == nil {
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	stationURLCount := parametersStream.ReadU32LENext(1)[0]
	stationUrls := make([]*nex.StationURL, 0)

	for i := 0; i < int(stationURLCount); i++ {
		station := nex.NewStationURL(parametersStream.ReadNEXStringNext())
		stationUrls = append(stationUrls, station)
	}

	dataHolderType := parametersStream.ReadNEXStringNext()

	if dataHolderType != "NintendoLoginData" {
		// Error log?
		return
	}

	_ = parametersStream.ReadU32LENext(1)[0] // Length including next buffer length field
	dataHolderInner := parametersStream.ReadNEXBufferNext()
	dataHolderStream := nex.NewStream(dataHolderInner)

	loginData := NintendoLoginData{token: dataHolderStream.ReadNEXStringNext()}

	secureProtocol.RegisterExHandler(client, callID, stationUrls, loginData)
}

func NewSecureProtocol(server *nex.Server) *SecureProtocol {
	secureProtocol := &SecureProtocol{
		server: server,
		ConnectionIDCounter: nex.NewCounter(10),
	}

	secureProtocol.Setup()

	return secureProtocol
}