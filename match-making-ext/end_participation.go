package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// EndParticipation sets the EndParticipation handler function
func (protocol *MatchMakingExtProtocol) EndParticipation(handler func(err error, client *nex.Client, callID uint32, idGathering uint32, strMessage string)) {
	protocol.EndParticipationHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleEndParticipation(packet nex.PacketInterface) {
	if protocol.EndParticipationHandler == nil {
		globals.Logger.Warning("MatchMakingExt::EndParticipation not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering := parametersStream.ReadUInt32LE()

	var err error
	var strMessage string
	strMessage, err = parametersStream.ReadString()
	if err != nil {
		go protocol.EndParticipationHandler(err, client, callID, 0, "")
	}

	go protocol.EndParticipationHandler(nil, client, callID, idGathering, strMessage)
}
