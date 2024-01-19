// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
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
	server                         nex.ServerInterface
	RequestProbeInitiation         func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)
	InitiateProbe                  func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, uint32)
	RequestProbeInitiationExt      func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, uint32)
	ReportNATTraversalResult       func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	ReportNATProperties            func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)
	GetRelaySignatureKey           func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	ReportNATTraversalResultDetail func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the NAT Traversal protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerRequestProbeInitiation(handler func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, uint32))
	SetHandlerInitiateProbe(handler func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, uint32))
	SetHandlerRequestProbeInitiationExt(handler func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, uint32))
	SetHandlerReportNATTraversalResult(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerReportNATProperties(handler func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32))
	SetHandlerGetRelaySignatureKey(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerReportNATTraversalResultDetail(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerRequestProbeInitiation sets the handler for the RequestProbeInitiation method
func (protocol *Protocol) SetHandlerRequestProbeInitiation(handler func(err error, packet nex.PacketInterface, callID uint32, urlTargetList *types.List[*types.StationURL]) (*nex.RMCMessage, uint32)) {
	protocol.RequestProbeInitiation = handler
}

// SetHandlerInitiateProbe sets the handler for the InitiateProbe method
func (protocol *Protocol) SetHandlerInitiateProbe(handler func(err error, packet nex.PacketInterface, callID uint32, urlStationToProbe *types.StationURL) (*nex.RMCMessage, uint32)) {
	protocol.InitiateProbe = handler
}

// SetHandlerRequestProbeInitiationExt sets the handler for the RequestProbeInitiationExt method
func (protocol *Protocol) SetHandlerRequestProbeInitiationExt(handler func(err error, packet nex.PacketInterface, callID uint32, targetList *types.List[*types.String], stationToProbe *types.String) (*nex.RMCMessage, uint32)) {
	protocol.RequestProbeInitiationExt = handler
}

// SetHandlerReportNATTraversalResult sets the handler for the ReportNATTraversalResult method
func (protocol *Protocol) SetHandlerReportNATTraversalResult(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ReportNATTraversalResult = handler
}

// SetHandlerReportNATProperties sets the handler for the ReportNATProperties method
func (protocol *Protocol) SetHandlerReportNATProperties(handler func(err error, packet nex.PacketInterface, callID uint32, natmapping *types.PrimitiveU32, natfiltering *types.PrimitiveU32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ReportNATProperties = handler
}

// SetHandlerGetRelaySignatureKey sets the handler for the GetRelaySignatureKey method
func (protocol *Protocol) SetHandlerGetRelaySignatureKey(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetRelaySignatureKey = handler
}

// SetHandlerReportNATTraversalResultDetail sets the handler for the ReportNATTraversalResultDetail method
func (protocol *Protocol) SetHandlerReportNATTraversalResultDetail(handler func(err error, packet nex.PacketInterface, callID uint32, cid *types.PrimitiveU32, result *types.PrimitiveBool, detail *types.PrimitiveS32, rtt *types.PrimitiveU32) (*nex.RMCMessage, uint32)) {
	protocol.ReportNATTraversalResultDetail = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
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
				globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
				fmt.Printf("Unsupported NATTraversal method ID: %#v\n", message.MethodID)
			}
		}
	})
}

// NewProtocol returns a new NAT Traversal NEX protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
