// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection protocol
	ProtocolID = 0xB

	// MethodRegister is the method ID for the method Register
	MethodRegister = 0x1

	// MethodRequestConnectionData is the method ID for the method RequestConnectionData
	MethodRequestConnectionData = 0x2

	// MethodRequestURLs is the method ID for the method RequestURLs
	MethodRequestURLs = 0x3

	// MethodRegisterEx is the method ID for the method RegisterEx
	MethodRegisterEx = 0x4

	// MethodTestConnectivity is the method ID for the method TestConnectivity
	MethodTestConnectivity = 0x5

	// MethodUpdateURLs is the method ID for the method UpdateURLs
	MethodUpdateURLs = 0x6

	// MethodReplaceURL is the method ID for the method ReplaceURL
	MethodReplaceURL = 0x7

	// MethodSendReport is the method ID for the method SendReport
	MethodSendReport = 0x8
)

// Protocol stores all the RMC method handlers for the Secure Connection protocol and listens for requests
type Protocol struct {
	Server                       nex.ServerInterface
	registerHandler              func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs []*nex.StationURL) uint32
	requestConnectionDataHandler func(err error, packet nex.PacketInterface, callID uint32, cidTarget uint32, pidTarget *nex.PID) uint32
	requestURLsHandler           func(err error, packet nex.PacketInterface, callID uint32, cidTarget uint32, pidTarget *nex.PID) uint32
	registerExHandler            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs []*nex.StationURL, hCustomData *nex.DataHolder) uint32
	testConnectivityHandler      func(err error, packet nex.PacketInterface, callID uint32) uint32
	updateURLsHandler            func(err error, packet nex.PacketInterface, callID uint32, vecMyURLs []*nex.StationURL) uint32
	replaceURLHandler            func(err error, packet nex.PacketInterface, callID uint32, target *nex.StationURL, url *nex.StationURL) uint32
	sendReportHandler            func(err error, packet nex.PacketInterface, callID uint32, reportID uint32, reportData []byte) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		if request.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodRegister:
		go protocol.handleRegister(packet)
	case MethodRequestConnectionData:
		go protocol.handleRequestConnectionData(packet)
	case MethodRequestURLs:
		go protocol.handleRequestURLs(packet)
	case MethodRegisterEx:
		go protocol.handleRegisterEx(packet)
	case MethodTestConnectivity:
		go protocol.handleTestConnectivity(packet)
	case MethodUpdateURLs:
		go protocol.handleUpdateURLs(packet)
	case MethodReplaceURL:
		go protocol.handleReplaceURL(packet)
	case MethodSendReport:
		go protocol.handleSendReport(packet)
	default:
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported SecureConnection method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Secure Connection protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
