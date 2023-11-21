// Package protocol implements the Screening protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the Screening protocol
	ProtocolID = 0x7C

	// MethodReportDataStoreContent is the method ID for the method ReportDataStoreContent
	MethodReportDataStoreContent = 0x1

	// MethodReportUser is the method ID for the method ReportUser
	MethodReportUser = 0x2
)

// Protocol handles the Screening protocol
type Protocol struct {
	Server                 nex.ServerInterface
	ReportDataStoreContent func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
	ReportUser             func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, uint32) // TODO - Unknown request/response format
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodReportDataStoreContent:
		protocol.handleReportDataStoreContent(packet)
	case MethodReportUser:
		protocol.handleReportUser(packet)
	default:
		fmt.Printf("Unsupported Screening method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Screening protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
