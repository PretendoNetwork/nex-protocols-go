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
	Server                        *nex.Server
	reportDataStoreContentHandler func(err error, client *nex.Client, callID uint32)
	reportUserHandler             func(err error, client *nex.Client, callID uint32)
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodReportDataStoreContent:
		go protocol.handleReportDataStoreContent(packet)
	case MethodReportUser:
		go protocol.handleReportUser(packet)
	default:
		fmt.Printf("Unsupported Screening method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Screening protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
