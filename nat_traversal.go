package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NATTraversalProtocolID is the protocol ID for the Message Delivery protocol
	NATTraversalProtocolID = 0x3

	// NATTraversalMethodInitiateProbe is the method ID for the method InitiateProbe
	NATTraversalMethodInitiateProbe = 0x2

	// NATTraversalMethodRequestProbeInitiationExt is the method ID for the method RequestProbeInitiationExt
	NATTraversalMethodRequestProbeInitiationExt = 0x3

	// NATTraversalMethodReportNATProperties is the method ID for the method ReportNATProperties
	NATTraversalMethodReportNATProperties = 0x5

	// NATTraversalMethodGetRelaySignatureKey is the method ID for the method GetRelaySignatureKey
	NATTraversalMethodGetRelaySignatureKey = 0x6
)

// AuthenticationProtocol handles the Authentication nex protocol
type NATTraversalProtocol struct {
	server                           *nex.Server
	InitiateProbeHandler             func(err error, client *nex.Client, callID uint32)
	RequestProbeInitiationExtHandler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)
	ReportNATPropertiesHandler       func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)
	GetRelaySignatureKeyHandler      func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (natTraversalProtocol *NATTraversalProtocol) Setup() {
	nexServer := natTraversalProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if NATTraversalProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case NATTraversalMethodRequestProbeInitiationExt:
				go natTraversalProtocol.handleRequestProbeInitiationExt(packet)
			case NATTraversalMethodReportNATProperties:
				go natTraversalProtocol.handleReportNATProperties(packet)
			case NATTraversalMethodGetRelaySignatureKey:
				go natTraversalProtocol.handleGetRelaySignatureKey(packet)
			default:
				go respondNotImplemented(packet, NATTraversalProtocolID)
				fmt.Printf("Unsupported NATTraversal method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// RequestProbeInitiationExt sets the RequestProbeInitiationExt handler function
func (natTraversalProtocol *NATTraversalProtocol) RequestProbeInitiationExt(handler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)) {
	natTraversalProtocol.RequestProbeInitiationExtHandler = handler
}

// ReportNATProperties sets the ReportNATProperties handler function
func (natTraversalProtocol *NATTraversalProtocol) ReportNATProperties(handler func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)) {
	natTraversalProtocol.ReportNATPropertiesHandler = handler
}

// GetRelaySignatureKey sets the GetRelaySignatureKey handler function
func (natTraversalProtocol *NATTraversalProtocol) GetRelaySignatureKey(handler func(err error, client *nex.Client, callID uint32)) {
	natTraversalProtocol.GetRelaySignatureKeyHandler = handler
}

func (natTraversalProtocol *NATTraversalProtocol) handleRequestProbeInitiationExt(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNATPropertiesHandler == nil {
		fmt.Println("[Warning] NATTraversalProtocol::RequestProbeInitiationExt not implemented")
		go respondNotImplemented(packet, NATTraversalProtocolID)
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

func (natTraversalProtocol *NATTraversalProtocol) handleReportNATProperties(packet nex.PacketInterface) {
	if natTraversalProtocol.ReportNATPropertiesHandler == nil {
		fmt.Println("[Warning] NATTraversalProtocol::ReportNATProperties not implemented")
		go respondNotImplemented(packet, NATTraversalProtocolID)
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

func (natTraversalProtocol *NATTraversalProtocol) handleGetRelaySignatureKey(packet nex.PacketInterface) {
	if natTraversalProtocol.GetRelaySignatureKeyHandler == nil {
		fmt.Println("[Warning] NATTraversalProtocol::GetRelaySignatureKey not implemented")
		go respondNotImplemented(packet, NATTraversalProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	//parameters := request.Parameters()

	go natTraversalProtocol.GetRelaySignatureKeyHandler(nil, client, callID)
}

// NewNATTraversalProtocol returns a new NATTraversalProtocol
func NewNATTraversalProtocol(server *nex.Server) *NATTraversalProtocol {
	natTraversalProtocol := &NATTraversalProtocol{server: server}

	natTraversalProtocol.Setup()

	return natTraversalProtocol
}
