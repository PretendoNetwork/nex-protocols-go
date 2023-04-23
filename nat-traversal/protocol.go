package nat_traversal

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x3

	// MethodInitiateProbe is the method ID for the method InitiateProbe
	MethodInitiateProbe = 0x2

	// MethodRequestProbeInitiationExt is the method ID for the method RequestProbeInitiationExt
	MethodRequestProbeInitiationExt = 0x3

	// MethodReportNATTraversalResult is the method ID for the method ReportNATTraversalResult
	MethodReportNATTraversalResult = 0x4

	// MethodReportNATProperties is the method ID for the method ReportNATProperties
	MethodReportNATProperties = 0x5

	// MethodGetRelaySignatureKey is the method ID for the method GetRelaySignatureKey
	MethodGetRelaySignatureKey = 0x6
)

// AuthenticationProtocol handles the Authentication nex protocol
type NATTraversalProtocol struct {
	Server                           *nex.Server
	InitiateProbeHandler             func(err error, client *nex.Client, callID uint32)
	RequestProbeInitiationExtHandler func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string)
	ReportNATTraversalResultHandler  func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32)
	ReportNATPropertiesHandler       func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32)
	GetRelaySignatureKeyHandler      func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *NATTraversalProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodRequestProbeInitiationExt:
				go protocol.HandleRequestProbeInitiationExt(packet)
			case MethodReportNATTraversalResult:
				go protocol.HandleReportNATTraversalResult(packet)
			case MethodReportNATProperties:
				go protocol.HandleReportNATProperties(packet)
			case MethodGetRelaySignatureKey:
				go protocol.HandleGetRelaySignatureKey(packet)
			default:
				go globals.RespondNotImplemented(packet, ProtocolID)
				fmt.Printf("Unsupported NATTraversal method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewNATTraversalProtocol returns a new NATTraversalProtocol
func NewNATTraversalProtocol(server *nex.Server) *NATTraversalProtocol {
	protocol := &NATTraversalProtocol{Server: server}

	protocol.Setup()

	return protocol
}