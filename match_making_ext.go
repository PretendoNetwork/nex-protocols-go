package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// MatchMakingExtProtocolID is the protocol ID for the MatchMakingExt protocol
	MatchMakingExtProtocolID = 0x32

	// MatchMakingExtMethodEndParticipation is the method ID for the method EndParticipation
	MatchMakingExtMethodEndParticipation = 0x1

	// MatchMakingExtMethodGetParticipants is the method ID for the method GetParticipants
	MatchMakingExtMethodGetParticipants = 0x2

	// MatchMakingExtMethodGetDetailedParticipants is the method ID for the method GetDetailedParticipants
	MatchMakingExtMethodGetDetailedParticipants = 0x3

	// MatchMakingExtMethodGetParticipantsURLs is the method ID for the method GetParticipantsURLs
	MatchMakingExtMethodGetParticipantsURLs = 0x4

	// MatchMakingExtMethodGetGatheringRelations is the method ID for the method GetGatheringRelations
	MatchMakingExtMethodGetGatheringRelations = 0x5

	// MatchMakingExtMethodDeleteFromDeletions is the method ID for the method DeleteFromDeletions
	MatchMakingExtMethodDeleteFromDeletions = 0x6
)

// MatchMakingExtProtocol handles the MatchMakingExt protocol
type MatchMakingExtProtocol struct {
	server                         *nex.Server
	EndParticipationHandler        func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)
	GetParticipantsHandler         func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)
	GetDetailedParticipantsHandler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)
	GetParticipantsURLsHandler     func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)
	GetGatheringRelationsHandler   func(err error, client *nex.Client, callID uint32, id uint32, descr string)
	DeleteFromDeletionsHandler     func(err error, client *nex.Client, callID uint32, lstDeletions []uint32, pid uint32)
}

// Setup initializes the protocol
func (matchMakingExtProtocol *MatchMakingExtProtocol) Setup() {
	nexServer := matchMakingExtProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if MatchMakingExtProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case MatchMakingExtMethodEndParticipation:
				go matchMakingExtProtocol.handleEndParticipation(packet)
				break
			case MatchMakingExtMethodGetParticipants:
				go matchMakingExtProtocol.handleGetParticipants(packet)
				break
			case MatchMakingExtMethodGetDetailedParticipants:
				go matchMakingExtProtocol.handleGetDetailedParticipants(packet)
				break
			case MatchMakingExtMethodGetParticipantsURLs:
				go matchMakingExtProtocol.handleGetParticipantsURLs(packet)
				break
			case MatchMakingExtMethodGetGatheringRelations:
				go matchMakingExtProtocol.handleGetGatheringRelations(packet)
				break
			case MatchMakingExtMethodDeleteFromDeletions:
				go matchMakingExtProtocol.handleDeleteFromDeletions(packet)
				break
			default:
				fmt.Printf("Unsupported MatchMakingExt method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleEndParticipation(packet nex.PacketInterface) {
	if matchMakingExtProtocol.EndParticipationHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::EndParticipation not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	idGathering := parametersStream.ReadUInt32LE()

	var err error
	var strMessage string
	strMessage, err = parametersStream.ReadString()
	if err != nil {
		go matchMakingExtProtocol.EndParticipationHandler(err, client, callID, 0, "")
	}

	go matchMakingExtProtocol.EndParticipationHandler(nil, client, callID, idGathering, strMessage)
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleGetParticipants(packet nex.PacketInterface) {
	if matchMakingExtProtocol.GetParticipantsHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::GetParticipants not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	idGathering := parametersStream.ReadUInt32LE()

	bOnlyActive := parametersStream.ReadUInt8() == 1

	go matchMakingExtProtocol.GetParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleGetDetailedParticipants(packet nex.PacketInterface) {
	if matchMakingExtProtocol.GetDetailedParticipantsHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::GetDetailedParticipants not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	idGathering := parametersStream.ReadUInt32LE()

	bOnlyActive := parametersStream.ReadUInt8() == 1

	go matchMakingExtProtocol.GetDetailedParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	if matchMakingExtProtocol.GetParticipantsURLsHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::GetParticipantsURLs not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	lstGatheringsCount := parametersStream.ReadUInt32LE()
	lstGatherings := make([]uint32, lstGatheringsCount)
	for i := 0; uint32(i) < lstGatheringsCount; i++ {
		lstGatherings[i] = parametersStream.ReadUInt32LE()
	}

	go matchMakingExtProtocol.GetParticipantsURLsHandler(nil, client, callID, lstGatherings)
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleGetGatheringRelations(packet nex.PacketInterface) {
	if matchMakingExtProtocol.GetGatheringRelationsHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::GetGatheringRelations not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	id := parametersStream.ReadUInt32LE()

	var err error
	var descr string
	descr, err = parametersStream.ReadString()
	if err != nil {
		go matchMakingExtProtocol.GetGatheringRelationsHandler(err, client, callID, 0, "")
	}

	go matchMakingExtProtocol.GetGatheringRelationsHandler(nil, client, callID, id, descr)
}

func (matchMakingExtProtocol *MatchMakingExtProtocol) handleDeleteFromDeletions(packet nex.PacketInterface) {
	if matchMakingExtProtocol.DeleteFromDeletionsHandler == nil {
		fmt.Println("[Warning] MatchMakingExtProtocol::DeleteFromDeletions not implemented")
		go respondNotImplemented(packet, MatchMakingExtProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, matchMakingExtProtocol.server)

	lstDeletionsCount := parametersStream.ReadUInt32LE()
	lstDeletions := make([]uint32, lstDeletionsCount)
	for i := 0; uint32(i) < lstDeletionsCount; i++ {
		lstDeletions[i] = parametersStream.ReadUInt32LE()
	}

	pid := parametersStream.ReadUInt32LE()

	go matchMakingExtProtocol.DeleteFromDeletionsHandler(nil, client, callID, lstDeletions, pid)
}

// EndParticipation sets the EndParticipation handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) EndParticipation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)) {
	matchMakingExtProtocol.EndParticipationHandler = handler
}

// GetParticipants sets the GetParticipants handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) GetParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	matchMakingExtProtocol.GetParticipantsHandler = handler
}

// GetDetailedParticipants sets the GetDetailedParticipants handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) GetDetailedParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	matchMakingExtProtocol.GetDetailedParticipantsHandler = handler
}

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) GetParticipantsURLs(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	matchMakingExtProtocol.GetParticipantsURLsHandler = handler
}

// GetGatheringRelations sets the GetGatheringRelations handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) GetGatheringRelations(handler func(err error, client *nex.Client, callID uint32, id uint32, descr string)) {
	matchMakingExtProtocol.GetGatheringRelationsHandler = handler
}

// DeleteFromDeletions sets the DeleteFromDeletions handler function
func (matchMakingExtProtocol *MatchMakingExtProtocol) DeleteFromDeletions(handler func(err error, client *nex.Client, callID uint32, lstDeletions []uint32, pid uint32)) {
	matchMakingExtProtocol.DeleteFromDeletionsHandler = handler
}

// NewMatchMakingExtProtocol returns a new MatchMakingExtProtocol
func NewMatchMakingExtProtocol(server *nex.Server) *MatchMakingExtProtocol {
	matchMakingExtProtocol := &MatchMakingExtProtocol{server: server}

	matchMakingExtProtocol.Setup()

	return matchMakingExtProtocol
}
