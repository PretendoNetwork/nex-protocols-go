package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NatTraversalProtocolID is the protocol ID for the Message Delivery protocol
	NatTraversalProtocolID = 0x3

	// NatTraversalInitiateProbe is the method ID for the method InitiateProbe
	NatTraversalInitiateProbe = 0x2

	// NatTraversalRequestProbeInitiationExt is the method ID for the method RequestProbeInitiationExt
	NatTraversalRequestProbeInitiationExt = 0x3

	// NatTraversalReportNatProperties is the method ID for the method ReportNatProperties
	NatTraversalReportNatProperties = 0x5

	// NatTraversalGetRelaySignatureKey is the method ID for the method GetRelaySignatureKey
	NatTraversalGetRelaySignatureKey = 0x6
)

// AuthenticationProtocol handles the Authentication nex protocol
type NatTraversalProtocol struct {
	server                              *nex.Server
	InitiateProbeHandler                func(err error, client *nex.Client, callID uint32)
	RequestProbeInitiationExtHandler    func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)
	ReportNatPropertiesHandler          func(err error, client *nex.Client, callID uint32)
	GetRelaySignatureKeyHandler         func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (natTraversalProtocol *NatTraversalProtocol) Setup() {
	nexServer := natTraversalProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if NatTraversalProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case NatTraversalRequestProbeInitiationExt:
				go natTraversalProtocol.handleRequestProbeInitiationExt(packet)
			case NatTraversalReportNatProperties:
				go natTraversalProtocol.handleReportNatProperties(packet)
			case NatTraversalGetRelaySignatureKey:
				go natTraversalProtocol.handleGetRelaySignatureKey(packet)
			default:
				go respondNotImplemented(packet, NatTraversalProtocolID)
				fmt.Printf("Unsupported NatTraversal method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (natTraversalProtocol *NatTraversalProtocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)) {
	natTraversalProtocol.RequestProbeInitiationExtHandler = handler
}

// ReportNatProperties sets the ReportNatProperties handler function
func (natTraversalProtocol *NatTraversalProtocol) ReportNatProperties(handler func(err error, client *nex.Client, callID uint32)) {
	natTraversalProtocol.ReportNatPropertiesHandler = handler
}

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (natTraversalProtocol *NatTraversalProtocol) GetRelaySignatureKey(handler func(err error, client *nex.Client, callID uint32)) {
	natTraversalProtocol.GetRelaySignatureKeyHandler = handler
}

func (natTraversalProtocol *NatTraversalProtocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNatPropertiesHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::RequestProbeInitiationExt not implemented")
		go respondNotImplemented(packet, NatTraversalProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, natTraversalProtocol.server)
	targetList := parametersStream.ReadListString()
	stationToProbe, err := parametersStream.ReadString()

	if err != nil {
		go natTraversalProtocol.RequestProbeInitiationExtHandler(err, client, callID, nil, "")
		return
	}

	go natTraversalProtocol.RequestProbeInitiationExtHandler(nil, client, callID, targetList, stationToProbe)
}

func (natTraversalProtocol *NatTraversalProtocol) handleReportNatProperties(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNatPropertiesHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::ReportNatProperties not implemented")
		go respondNotImplemented(packet, NatTraversalProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	//parameters := request.Parameters()

	go natTraversalProtocol.ReportNatPropertiesHandler(nil, client, callID)
}

func (natTraversalProtocol *NatTraversalProtocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	if natTraversalProtocol.GetRelaySignatureKeyHandler == nil {
		fmt.Println("[Warning] NatTraversalProtocol::GetRelaySignatureKey not implemented")
		go respondNotImplemented(packet, NatTraversalProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	//parameters := request.Parameters()

	go natTraversalProtocol.GetRelaySignatureKeyHandler(nil, client, callID)
}

// NewNatTraversalProtocol returns a new NatTraversalProtocol
func NewNatTraversalProtocol(server *nex.Server) *NatTraversalProtocol {
	natTraversalProtocol := &NatTraversalProtocol{server: server}

	natTraversalProtocol.Setup()

	return natTraversalProtocol
}
