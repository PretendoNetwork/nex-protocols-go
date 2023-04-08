package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipants sets the GetParticipants handler function
func (protocol *MatchMakingExtProtocol) GetParticipants(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, bOnlyActive bool)) {
	protocol.GetParticipantsHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleGetParticipants(packet nex.PacketInterface) {
	if protocol.GetParticipantsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipants not implemented")
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

	go protocol.GetParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}
