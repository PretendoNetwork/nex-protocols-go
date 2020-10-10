package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NatTraversalProtocolID is the protocol ID for the NatTraversal protocol
	NatTraversalProtocolID = 0x3

	// NatTraversalMethodRequestProbeInitiation is the method ID for the method RequestProbeInitiation
	NatTraversalMethodRequestProbeInitiation = 0x1

	// NatTraversalMethodInitiateProbe is the method ID for the method InitiateProbe
	NatTraversalMethodInitiateProbe = 0x2

	// NatTraversalMethodRequestProbeInitiationExt is the method ID for the method RequestProbeInitiationExt
	NatTraversalMethodRequestProbeInitiationExt = 0x3

	// NatTraversalMethodReportNATTraversalResult is the method ID for the method ReportNATTraversalResult
	NatTraversalMethodReportNATTraversalResult = 0x4

	// NatTraversalMethodReportNATProperties is the method ID for the method ReportNATProperties
	NatTraversalMethodReportNATProperties = 0x5

	// NatTraversalMethodGetRelaySignatureKey is the method ID for the method GetRelaySignatureKey
	NatTraversalMethodGetRelaySignatureKey = 0x6

	// NatTraversalMethodReportNATTraversalResultDetail is the method ID for the method ReportNATTraversalResultDetail
	NatTraversalMethodReportNATTraversalResultDetail = 0x7
)

// NatTraversalProtocol handles the NatTraversal nex protocol
type NatTraversalProtocol struct {
	server                                *nex.Server
	RequestProbeInitiationHandler         func(err error, client *nex.Client, callID uint32, urlTargetList []string)
	InitiateProbeHandler                  func(err error, client *nex.Client, callID uint32, urlStationToProbe string)
	RequestProbeInitiationExtHandler      func(err error, client *nex.Client, callID uint32, urlTargetList []string, urlStationToProbe string)
	ReportNATTraversalResultHandler       func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)
	ReportNATPropertiesHandler            func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)
	GetRelaySignatureKeyHandler           func(err error, client *nex.Client, callID uint32)
	ReportNATTraversalResultDetailHandler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, detail int32, rtt uint32)
}

// Setup intitalizes the protocol
func (natTraversalProtocol *NatTraversalProtocol) Setup() {
	nexServer := natTraversalProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if NatTraversalProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case NatTraversalMethodRequestProbeInitiation:
				go natTraversalProtocol.handleRequestProbeInitiation(packet)
				break
			case NatTraversalMethodInitiateProbe:
				go natTraversalProtocol.handleInitiateProbe(packet)
				break
			case NatTraversalMethodRequestProbeInitiationExt:
				go natTraversalProtocol.handleRequestProbeInitiationExt(packet)
				break
			case NatTraversalMethodReportNATTraversalResult:
				go natTraversalProtocol.handleReportNATTraversalResult(packet)
				break
			case NatTraversalMethodReportNATProperties:
				go natTraversalProtocol.handleReportNATProperties(packet)
				break
			case NatTraversalMethodGetRelaySignatureKey:
				go natTraversalProtocol.handleGetRelaySignatureKey(packet)
				break
			case NatTraversalMethodReportNATTraversalResultDetail:
				go natTraversalProtocol.handleReportNATTraversalResultDetail(packet)
				break
			default:
				fmt.Printf("Unsupported NatTraversal method ID: %#v\n", request.MethodID())
				go natTraversalProtocol.respondNotImplemented(packet)
				break
			}
		}
	})
}

func (natTraversalProtocol *NatTraversalProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.Sender()
	request := packet.RMCRequest()

	rmcResponse := nex.NewRMCResponse(NatTraversalProtocolID, request.CallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.Version() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.Version())
	responsePacket.SetSource(packet.Destination())
	responsePacket.SetDestination(packet.Source())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)
	natTraversalProtocol.server.Send(responsePacket)
}

func (natTraversalProtocol *NatTraversalProtocol) handleRequestProbeInitiation(packet nex.PacketInterface) {
	if natTraversalProtocol.RequestProbeInitiationHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::RequestProbeInitiation not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	urlTargetListCount := parametersStream.ReadUInt32LE()
	urlTargetList := make([]string, urlTargetListCount)
	for i := 0; i < int(urlTargetListCount); i++ {
		var e error
		urlTargetList[i], e = parametersStream.ReadString()
		if e != nil {
			go natTraversalProtocol.RequestProbeInitiationHandler(e, client, callID, nil)
		}
	}

	go natTraversalProtocol.RequestProbeInitiationHandler(nil, client, callID, urlTargetList)
}

func (natTraversalProtocol *NatTraversalProtocol) handleInitiateProbe(packet nex.PacketInterface) {
	if natTraversalProtocol.InitiateProbeHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::InitiateProbe not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	var e error
	urlStationToProbe, e := parametersStream.ReadString()
	if e != nil {
		go natTraversalProtocol.InitiateProbeHandler(e, client, callID, "")
	}

	go natTraversalProtocol.InitiateProbeHandler(nil, client, callID, urlStationToProbe)
}

func (natTraversalProtocol *NatTraversalProtocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if natTraversalProtocol.RequestProbeInitiationExtHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::RequestProbeInitiationExt not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	urlTargetListCount := parametersStream.ReadUInt32LE()
	urlTargetList := make([]string, urlTargetListCount)
	for i := 0; i < int(urlTargetListCount); i++ {
		var e error
		urlTargetList[i], e = parametersStream.ReadString()
		if e != nil {
			go natTraversalProtocol.RequestProbeInitiationExtHandler(e, client, callID, nil, "")
		}
	}
	var e error
	urlStationToProbe, e := parametersStream.ReadString()
	if e != nil {
		go natTraversalProtocol.RequestProbeInitiationExtHandler(e, client, callID, nil, "")
	}

	go natTraversalProtocol.RequestProbeInitiationExtHandler(nil, client, callID, urlTargetList, urlStationToProbe)
}

func (natTraversalProtocol *NatTraversalProtocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNATTraversalResultHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::ReportNATTraversalResult not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	cid := parametersStream.ReadUInt32LE()
	result := parametersStream.ReadByteNext()
	rtt := uint32(0)
	if parametersStream.ByteOffset() < parametersStream.ByteCapacity()-1 {
		rtt = parametersStream.ReadUInt32LE()
	}

	go natTraversalProtocol.ReportNATTraversalResultHandler(nil, client, callID, cid, result&1 == 1, rtt)
}

func (natTraversalProtocol *NatTraversalProtocol) handleReportNATProperties(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNATPropertiesHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::ReportNATProperties not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	natmapping := parametersStream.ReadUInt32LE()
	natfiltering := parametersStream.ReadUInt32LE()
	rtt := parametersStream.ReadUInt32LE()

	go natTraversalProtocol.ReportNATPropertiesHandler(nil, client, callID, natmapping, natfiltering, rtt)
}

func (natTraversalProtocol *NatTraversalProtocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	if natTraversalProtocol.GetRelaySignatureKeyHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::GetRelaySignatureKey not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go natTraversalProtocol.GetRelaySignatureKeyHandler(nil, client, callID)
}

func (natTraversalProtocol *NatTraversalProtocol) handleReportNATTraversalResultDetail(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNATTraversalResultDetailHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::ReportNATTraversalResultDetail not implemented")
		go natTraversalProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)

	cid := parametersStream.ReadUInt32LE()
	result := parametersStream.ReadByteNext()
	detail := int32(parametersStream.ReadUInt32LE())
	rtt := parametersStream.ReadUInt32LE()

	go natTraversalProtocol.ReportNATTraversalResultDetailHandler(nil, client, callID, cid, result&1 == 1, detail, rtt)
}

// RequestProbeInitiation sets the RequestProbeInitiation handler function
func (natTraversalProtocol *NatTraversalProtocol) RequestProbeInitiation(handler func(err error, client *nex.Client, callID uint32, urlTargetList []string)) {
	natTraversalProtocol.RequestProbeInitiationHandler = handler
}

// InitiateProbe sets the InitiateProbe handler function
func (natTraversalProtocol *NatTraversalProtocol) InitiateProbe(handler func(err error, client *nex.Client, callID uint32, urlStationToProbe string)) {
	natTraversalProtocol.InitiateProbeHandler = handler
}

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (natTraversalProtocol *NatTraversalProtocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, urlTargetList []string, urlStationToProbe string)) {
	natTraversalProtocol.RequestProbeInitiationExtHandler = handler
}

// ReportNATTraversalResult sets the ReportNATTraversalResult handler function
func (natTraversalProtocol *NatTraversalProtocol) ReportNATTraversalResult(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)) {
	natTraversalProtocol.ReportNATTraversalResultHandler = handler
}

// ReportNATProperties sets the ReportNATProperties handler function
func (natTraversalProtocol *NatTraversalProtocol) ReportNATProperties(handler func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)) {
	natTraversalProtocol.ReportNATPropertiesHandler = handler
}

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (natTraversalProtocol *NatTraversalProtocol) GetRelaySignatureKey(handler func(err error, client *nex.Client, callID uint32)) {
	natTraversalProtocol.GetRelaySignatureKeyHandler = handler
}

// ReportNATTraversalResultDetail sets the ReportNATTraversalResultDetail handler function
func (natTraversalProtocol *NatTraversalProtocol) ReportNATTraversalResultDetail(handler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, detail int32, rtt uint32)) {
	natTraversalProtocol.ReportNATTraversalResultDetailHandler = handler
}

//NewNatTraversalProtocol returns a new NatTraversalProtocol
func NewNatTraversalProtocol(server *nex.Server) *NatTraversalProtocol {
	natTraversalProtocol := &NatTraversalProtocol{server: server}

	natTraversalProtocol.Setup()

	return natTraversalProtocol
}
