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
	Server                         *nex.Server
	endParticipationHandler        func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string) uint32
	getParticipantsHandler         func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool) uint32
	getDetailedParticipantsHandler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool) uint32
	getParticipantsURLsHandler     func(err error, client *nex.Client, callID uint32, lstGatherings []uint32) uint32
	getGatheringRelationsHandler   func(err error, client *nex.Client, callID uint32, id uint32, descr string) uint32
	deleteFromDeletionsHandler     func(err error, client *nex.Client, callID uint32, lstDeletions []uint32, pid uint32) uint32
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
		fmt.Printf("Unsupported MatchMakingExt method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Match Making Ext protocol
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}

	protocol.Setup()

	return protocol
}
