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
	Server                         nex.ServerInterface
	RequestProbeInitiation         func(err error, packet nex.PacketInterface, callID uint32, urlTargetList []*nex.StationURL) (*nex.RMCMessage, uint32)
	InitiateProbe                  func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *nex.StationURL) (*nex.RMCMessage, uint32)
	RequestProbeInitiationExt      func(err error, packet nex.PacketInterface, callID uint32, targetList []string, stationToProbe string) (*nex.RMCMessage, uint32)
	ReportNATTraversalResult       func(err error, packet nex.PacketInterface, callID uint32, cid uint32, result bool, rtt uint32) (*nex.RMCMessage, uint32)
	ReportNATProperties            func(err error, packet nex.PacketInterface, callID uint32, natmapping uint32, natfiltering uint32, rtt uint32) (*nex.RMCMessage, uint32)
	GetRelaySignatureKey           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	ReportNATTraversalResultDetail func(err error, packet nex.PacketInterface, callID uint32, cid uint32, result bool, detail int32, rtt uint32) (*nex.RMCMessage, uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			switch message.MethodID {
			case MethodRequestProbeInitiation:
				protocol.handleRequestProbeInitiation(packet)
			case MethodInitiateProbe:
				protocol.handleInitiateProbe(packet)
			case MethodRequestProbeInitiationExt:
				protocol.handleRequestProbeInitiationExt(packet)
			case MethodReportNATTraversalResult:
				protocol.handleReportNATTraversalResult(packet)
			case MethodReportNATProperties:
				protocol.handleReportNATProperties(packet)
			case MethodGetRelaySignatureKey:
				protocol.handleGetRelaySignatureKey(packet)
			case MethodReportNATTraversalResultDetail:
				protocol.handleReportNATTraversalResultDetail(packet)
			default:
				go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported NATTraversal method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewNATTraversalProtocol returns a new NAT Traversal NEX protocol
func NewNATTraversalProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
