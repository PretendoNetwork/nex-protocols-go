// Package protocol implements the Ranking (Legacy) protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Ranking (Legacy) protocol
	ProtocolID = 0x70

	// MethodUploadCommonData is the method ID for method UploadCommonData
	MethodUploadCommonData = 0x5

	// MethodUnknown0xE is the method ID for unknown method 0xE
	MethodUnknown0xE = 0xE

	// MethodUnknown0xF is the method ID for unknown method 0xF
	MethodUnknown0xF = 0xF

	// MethodGetTotal is the method ID for method GetTotal
	MethodGetTotal = 0x10

	// MethodUploadScoreWithLimit is the method ID for method UploadScoreWithLimit
	MethodUploadScoreWithLimit = 0x11

	// MethodUploadSpecificPeriodScore is the method ID for method UploadSpecificPeriodScore
	MethodUploadSpecificPeriodScore = 0x14

	// MethodGetSpecificPeriodDataList is the method ID for method GetSpecificPeriodDataList
	MethodGetSpecificPeriodDataList = 0x16

	// MethodGetSpecificPeriodTotal is the method ID for method GetSpecificPeriodTotal
	MethodGetSpecificPeriodTotal = 0x19
)

// Protocol stores all the RMC method handlers for the Ranking (Legacy) protocol and listens for requests
type Protocol struct {
	server                    nex.ServerInterface
	UploadCommonData          func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, commonData []byte) (*nex.RMCMessage, uint32)
	Unknown0xE                func(err error, packet nex.PacketInterface, callID uint32, rankingMode uint8, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, offset uint32, length uint8) (*nex.RMCMessage, uint32)
	Unknown0xF                func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, length uint8) (*nex.RMCMessage, uint32)
	GetTotal                  func(err error, packet nex.PacketInterface, callID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint32) (*nex.RMCMessage, uint32)
	UploadScoreWithLimit      func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scores []uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32)
	UploadSpecificPeriodScore func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, score uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32)
	GetSpecificPeriodDataList func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, offset uint32, length uint8) (*nex.RMCMessage, uint32)
	GetSpecificPeriodTotal    func(err error, packet nex.PacketInterface, callID uint32, category uint32) (*nex.RMCMessage, uint32)
}

type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, commonData []byte) (*nex.RMCMessage, uint32))
	SetHandlerUnknown0xE(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode uint8, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, offset uint32, length uint8) (*nex.RMCMessage, uint32))
	SetHandlerUnknown0xF(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, length uint8) (*nex.RMCMessage, uint32))
	SetHandlerGetTotal(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint32) (*nex.RMCMessage, uint32))
	SetHandlerUploadScoreWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scores []uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32))
	SetHandlerUploadSpecificPeriodScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, score uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32))
	SetHandlerGetSpecificPeriodDataList(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, offset uint32, length uint8) (*nex.RMCMessage, uint32))
	SetHandlerGetSpecificPeriodTotal(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerUploadCommonData sets the handler for the UploadCommonData method
func (protocol *Protocol) SetHandlerUploadCommonData(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, commonData []byte) (*nex.RMCMessage, uint32)) {
	protocol.UploadCommonData = handler
}

// SetHandlerUUnknown0xE sets the handler for the Unknown0xE method
func (protocol *Protocol) SetHandlerUnknown0xE(handler func(err error, packet nex.PacketInterface, callID uint32, rankingMode uint8, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, offset uint32, length uint8) (*nex.RMCMessage, uint32)) {
	protocol.Unknown0xE = handler
}

// SetHandlerUUnknown0xF sets the handler for the Unknown0xF method
func (protocol *Protocol) SetHandlerUnknown0xF(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scoreIndex uint8, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint8, unknown5 uint8, unknown6 uint32, length uint8) (*nex.RMCMessage, uint32)) {
	protocol.Unknown0xF = handler
}

// SetHandlerGetTotal sets the handler for the GetTotal method
func (protocol *Protocol) SetHandlerGetTotal(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, unknown4 uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetTotal = handler
}

// SetHandlerUploadScoreWithLimit sets the handler for the UploadScoreWithLimit method
func (protocol *Protocol) SetHandlerUploadScoreWithLimit(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, scores []uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32)) {
	protocol.UploadScoreWithLimit = handler
}

// SetHandlerUploadSpecificPeriodScore sets the handler for the UploadSpecificPeriodScore method
func (protocol *Protocol) SetHandlerUploadSpecificPeriodScore(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, score uint32, unknown1 uint8, unknown2 uint32, unknown3 uint16) (*nex.RMCMessage, uint32)) {
	protocol.UploadSpecificPeriodScore = handler
}

// SetHandlerGetSpecificPeriodDataList sets the handler for the GetSpecificPeriodDataList method
func (protocol *Protocol) SetHandlerGetSpecificPeriodDataList(handler func(err error, packet nex.PacketInterface, callID uint32, uniqueID uint32, category uint32, unknown1 uint8, unknown2 uint8, unknown3 uint8, offset uint32, length uint8) (*nex.RMCMessage, uint32)) {
	protocol.GetSpecificPeriodDataList = handler
}

// SetHandlerGetSpecificPeriodTotal sets the handler for the GetSpecificPeriodTotal method
func (protocol *Protocol) SetHandlerGetSpecificPeriodTotal(handler func(err error, packet nex.PacketInterface, callID uint32, category uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetSpecificPeriodTotal = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	if request.ProtocolID == ProtocolID {
		switch request.MethodID {
		case MethodUploadCommonData:
			protocol.handleUploadCommonData(packet)
		case MethodUnknown0xE:
			protocol.handleUnknown0xE(packet)
		case MethodUnknown0xF:
			protocol.handleUnknown0xF(packet)
		case MethodGetTotal:
			protocol.handleGetTotal(packet)
		case MethodUploadScoreWithLimit:
			protocol.handleUploadScoreWithLimit(packet)
		case MethodUploadSpecificPeriodScore:
			protocol.handleUploadSpecificPeriodScore(packet)
		case MethodGetSpecificPeriodDataList:
			protocol.handleGetSpecificPeriodDataList(packet)
		case MethodGetSpecificPeriodTotal:
			protocol.handleGetSpecificPeriodTotal(packet)
		default:
			globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
			fmt.Printf("Unsupported Ranking (Legacy) method ID: %#v\n", request.MethodID)
		}
	}
}

// NewProtocol returns a new Ranking (Legacy) protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
