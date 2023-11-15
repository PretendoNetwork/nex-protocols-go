// Package protocol implements the Match Making Ext protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ProtocolID is the protocol ID for the MatchMakingExt protocol
	ProtocolID = 0x32

	// MethodEndParticipation is the method ID for the method EndParticipation
	MethodEndParticipation = 0x1

	// MethodGetParticipants is the method ID for the method GetParticipants
	MethodGetParticipants = 0x2

	// MethodGetDetailedParticipants is the method ID for the method GetDetailedParticipants
	MethodGetDetailedParticipants = 0x3

	// MethodGetParticipantsURLs is the method ID for the method GetParticipantsURLs
	MethodGetParticipantsURLs = 0x4

	// MethodGetGatheringRelations is the method ID for the method GetGatheringRelations
	MethodGetGatheringRelations = 0x5

	// MethodDeleteFromDeletions is the method ID for the method DeleteFromDeletions
	MethodDeleteFromDeletions = 0x6
)

// Protocol handles the MatchMakingExt protocol
type Protocol struct {
	Server                  nex.ServerInterface
	EndParticipation        func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, strMessage string) (*nex.RMCMessage, uint32)
	GetParticipants         func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, bOnlyActive bool) (*nex.RMCMessage, uint32)
	GetDetailedParticipants func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, bOnlyActive bool) (*nex.RMCMessage, uint32)
	GetParticipantsURLs     func(err error, packet nex.PacketInterface, callID uint32, lstGatherings []uint32) (*nex.RMCMessage, uint32)
	GetGatheringRelations   func(err error, packet nex.PacketInterface, callID uint32, id uint32, descr string) (*nex.RMCMessage, uint32)
	DeleteFromDeletions     func(err error, packet nex.PacketInterface, callID uint32, lstDeletions []uint32, pid *nex.PID) (*nex.RMCMessage, uint32)
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
	case MethodEndParticipation:
		go protocol.handleEndParticipation(packet)
	case MethodGetParticipants:
		go protocol.handleGetParticipants(packet)
	case MethodGetDetailedParticipants:
		go protocol.handleGetDetailedParticipants(packet)
	case MethodGetParticipantsURLs:
		go protocol.handleGetParticipantsURLs(packet)
	case MethodGetGatheringRelations:
		go protocol.handleGetGatheringRelations(packet)
	case MethodDeleteFromDeletions:
		go protocol.handleDeleteFromDeletions(packet)
	default:
		fmt.Printf("Unsupported MatchMakingExt method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Match Making Ext protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
