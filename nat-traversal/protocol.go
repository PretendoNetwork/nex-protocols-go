// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Message Delivery protocol
	ProtocolID = 0x3

	// MethodRequestProbeInitiation is the method ID for the method RequestProbeInitiation
	MethodRequestProbeInitiation = 0x1

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

	// MethodReportNATTraversalResultDetail is the method ID for the method ReportNATTraversalResultDetail
	MethodReportNATTraversalResultDetail = 0x7
)

// Protocol stores all the RMC method handlers for the NAT Traversal protocol and listens for requests
type Protocol struct {
	Server                                *nex.Server
	requestProbeInitiationHandler         func(err error, client *nex.Client, callID uint32, urlTargetList []*nex.StationURL) uint32
	initiateProbeHandler                  func(err error, client *nex.Client, callID uint32, urlStationToProbe *nex.StationURL) uint32
	requestProbeInitiationExtHandler      func(err error, client *nex.Client, callID uint32, targetList []string, stationToProbe string) uint32
	reportNATTraversalResultHandler       func(err error, client *nex.Client, callID uint32, cid uint32, result bool, rtt uint32) uint32
	reportNATPropertiesHandler            func(err error, client *nex.Client, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32) uint32
	getRelaySignatureKeyHandler           func(err error, client *nex.Client, callID uint32) uint32
	reportNATTraversalResultDetailHandler func(err error, client *nex.Client, callID uint32, cid uint32, result bool, detail int32, rtt uint32) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			switch request.MethodID() {
			case MethodRequestProbeInitiation:
				go protocol.handleRequestProbeInitiation(packet)
			case MethodInitiateProbe:
				go protocol.handleInitiateProbe(packet)
			case MethodRequestProbeInitiationExt:
				go protocol.handleRequestProbeInitiationExt(packet)
			case MethodReportNATTraversalResult:
				go protocol.handleReportNATTraversalResult(packet)
			case MethodReportNATProperties:
				go protocol.handleReportNATProperties(packet)
			case MethodGetRelaySignatureKey:
				go protocol.handleGetRelaySignatureKey(packet)
			case MethodReportNATTraversalResultDetail:
				go protocol.handleReportNATTraversalResultDetail(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported NATTraversal method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// NewNATTraversalProtocol returns a new NAT Traversal NEX protocol
func NewNATTraversalProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
