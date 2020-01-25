package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	SecureProtocolID = 0xB

	SecureMethodRegister              = 0x1
	SecureMethodRequestConnectionData = 0x2
	SecureMethodRequestURLs           = 0x3
	SecureMethodRegisterEx            = 0x4
	SecureMethodTestConnectivity      = 0x5
	SecureMethodUpdateURLs            = 0x6
	SecureMethodReplaceURL            = 0x7
	SecureMethodSendReport            = 0x8
)

type SecureProtocol struct {
	server                       *nex.Server
	ConnectionIDCounter          *nex.Counter
	RegisterHandler              func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)
	RequestConnectionDataHandler func(client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)
	RequestURLsHandler           func(client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)
	RegisterExHandler            func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData NintendoLoginData)
	TestConnectivityHandler      func(client *nex.Client, callID uint32)
	UpdateURLsHandler            func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)
	ReplaceURLHandler            func(client *nex.Client, callID uint32, oldStation *nex.StationURL, newStation *nex.StationURL)
	SendReportHandler            func(client *nex.Client, callID uint32, reportID uint32, report []byte)
}

func (secureProtocol *SecureProtocol) Setup() {
	nexServer := secureProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.GetRMCRequest()

		if SecureProtocolID == request.GetProtocolID() {
			switch request.GetMethodID() {
			case SecureMethodRegister:
				go secureProtocol.handleRegister(packet)
			case SecureMethodRequestConnectionData:
				go secureProtocol.handleRequestConnectionData(packet)
			case SecureMethodRequestURLs:
				go secureProtocol.handleRequestURLs(packet)
			case SecureMethodRegisterEx:
				go secureProtocol.handleRegisterEx(packet)
			case SecureMethodTestConnectivity:
				go secureProtocol.handleTestConnectivity(packet)
			case SecureMethodUpdateURLs:
				go secureProtocol.handleUpdateURLs(packet)
			case SecureMethodReplaceURL:
				go secureProtocol.handleReplaceURL(packet)
			case SecureMethodSendReport:
				go secureProtocol.handleSendReport(packet)
			default:
				fmt.Printf("Unsupported Secure method ID: %#v\n", request.GetMethodID())
			}
		}
	})
}

func (secureProtocol *SecureProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.GetSender()
	request := packet.GetRMCRequest()

	rmcResponse := nex.NewRMCResponse(SecureProtocolID, request.GetCallID())
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

	secureProtocol.server.Send(responsePacket)
}

func (secureProtocol *SecureProtocol) Register(handler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)) {
	secureProtocol.RegisterHandler = handler
}

func (secureProtocol *SecureProtocol) RequestConnectionData(handler func(client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)) {
	secureProtocol.RequestConnectionDataHandler = handler
}

func (secureProtocol *SecureProtocol) RequestURLs(handler func(client *nex.Client, callID uint32, stationCID uint32, stationPID uint32)) {
	secureProtocol.RequestURLsHandler = handler
}

func (secureProtocol *SecureProtocol) RegisterEx(handler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL, loginData NintendoLoginData)) {
	secureProtocol.RegisterExHandler = handler
}

func (secureProtocol *SecureProtocol) TestConnectivity(handler func(client *nex.Client, callID uint32)) {
	secureProtocol.TestConnectivityHandler = handler
}

func (secureProtocol *SecureProtocol) UpdateURLs(handler func(client *nex.Client, callID uint32, stationUrls []*nex.StationURL)) {
	secureProtocol.UpdateURLsHandler = handler
}

func (secureProtocol *SecureProtocol) ReplaceURL(handler func(client *nex.Client, callID uint32, oldStation *nex.StationURL, newStation *nex.StationURL)) {
	secureProtocol.ReplaceURLHandler = handler
}

func (secureProtocol *SecureProtocol) SendReport(handler func(client *nex.Client, callID uint32, reportID uint32, report []byte)) {
	secureProtocol.SendReportHandler = handler
}

func (secureProtocol *SecureProtocol) handleRegister(packet nex.PacketInterface) {
	if secureProtocol.RegisterHandler == nil {
		fmt.Println("[Warning] SecureProtocol::Register not implemented")
		go secureProtocol.respondNotImplemented(packet)
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

	go secureProtocol.RegisterHandler(client, callID, stationUrls)
}

func (secureProtocol *SecureProtocol) handleRequestConnectionData(packet nex.PacketInterface) {
	if secureProtocol.RequestConnectionDataHandler == nil {
		fmt.Println("[Warning] SecureProtocol::RequestConnectionData not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	stationCID := parametersStream.ReadU32LENext(1)[0]
	stationPID := parametersStream.ReadU32LENext(1)[0]

	go secureProtocol.RequestConnectionDataHandler(client, callID, stationCID, stationPID)
}

func (secureProtocol *SecureProtocol) handleRequestURLs(packet nex.PacketInterface) {
	if secureProtocol.RequestURLsHandler == nil {
		fmt.Println("[Warning] SecureProtocol::RequestURLs not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	stationCID := parametersStream.ReadU32LENext(1)[0]
	stationPID := parametersStream.ReadU32LENext(1)[0]

	go secureProtocol.RequestURLsHandler(client, callID, stationCID, stationPID)
}

func (secureProtocol *SecureProtocol) handleRegisterEx(packet nex.PacketInterface) {
	if secureProtocol.RegisterExHandler == nil {
		fmt.Println("[Warning] SecureProtocol::RegisterEx not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
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

	go secureProtocol.RegisterExHandler(client, callID, stationUrls, loginData)
}

func (secureProtocol *SecureProtocol) handleTestConnectivity(packet nex.PacketInterface) {
	if secureProtocol.TestConnectivityHandler == nil {
		fmt.Println("[Warning] SecureProtocol::TestConnectivity not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()

	go secureProtocol.TestConnectivityHandler(client, callID)
}

func (secureProtocol *SecureProtocol) handleUpdateURLs(packet nex.PacketInterface) {
	if secureProtocol.UpdateURLsHandler == nil {
		fmt.Println("[Warning] SecureProtocol::UpdateURLs not implemented")
		go secureProtocol.respondNotImplemented(packet)
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

	go secureProtocol.UpdateURLsHandler(client, callID, stationUrls)
}

func (secureProtocol *SecureProtocol) handleReplaceURL(packet nex.PacketInterface) {
	if secureProtocol.ReplaceURLHandler == nil {
		fmt.Println("[Warning] SecureProtocol::ReplaceURL not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	oldStation := nex.NewStationURL(parametersStream.ReadNEXStringNext())
	newStation := nex.NewStationURL(parametersStream.ReadNEXStringNext())

	go secureProtocol.ReplaceURLHandler(client, callID, oldStation, newStation)
}

func (secureProtocol *SecureProtocol) handleSendReport(packet nex.PacketInterface) {
	if secureProtocol.SendReportHandler == nil {
		fmt.Println("[Warning] SecureProtocol::SendReport not implemented")
		go secureProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.GetSender()
	request := packet.GetRMCRequest()

	callID := request.GetCallID()
	parameters := request.GetParameters()

	parametersStream := nex.NewStream(parameters)

	reportID := parametersStream.ReadU32LENext(1)[0]
	report := parametersStream.ReadNEXQBufferNext()

	go secureProtocol.SendReportHandler(client, callID, reportID, report)
}

func NewSecureProtocol(server *nex.Server) *SecureProtocol {
	secureProtocol := &SecureProtocol{
		server:              server,
		ConnectionIDCounter: nex.NewCounter(10),
	}

	secureProtocol.Setup()

	return secureProtocol
}
