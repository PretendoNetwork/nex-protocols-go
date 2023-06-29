package match_making_ext

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

// MatchMakingExtProtocol handles the MatchMakingExt protocol
type MatchMakingExtProtocol struct {
	Server                         *nex.Server
	EndParticipationHandler        func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)
	GetParticipantsHandler         func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)
	GetDetailedParticipantsHandler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)
	GetParticipantsURLsHandler     func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)
	GetGatheringRelationsHandler   func(err error, client *nex.Client, callID uint32, id uint32, descr string)
	DeleteFromDeletionsHandler     func(err error, client *nex.Client, callID uint32, lstDeletions []uint32, pid uint32)
}

// Setup initializes the protocol
func (protocol *MatchMakingExtProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

func (protocol *MatchMakingExtProtocol) HandlePacket(packet nex.PacketInterface) {
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

// NewMatchMakingExtProtocol returns a new MatchMakingExtProtocol
func NewMatchMakingExtProtocol(server *nex.Server) *MatchMakingExtProtocol {
	protocol := &MatchMakingExtProtocol{Server: server}

	protocol.Setup()

	return protocol
}
