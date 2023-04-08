package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDetailedParticipants sets the GetDetailedParticipants handler function
func (protocol *MatchMakingExtProtocol) GetDetailedParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	protocol.GetDetailedParticipantsHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleGetDetailedParticipants(packet nex.PacketInterface) {
	if protocol.GetDetailedParticipantsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetDetailedParticipants not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering := parametersStream.ReadUInt32LE()

	bOnlyActive := parametersStream.ReadUInt8() == 1

	go protocol.GetDetailedParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}
