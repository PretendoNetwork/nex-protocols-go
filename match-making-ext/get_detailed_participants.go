package match_making_ext

import (
	"fmt"

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

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.GetDetailedParticipantsHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0, false)
	}

	bOnlyActive, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.GetDetailedParticipantsHandler(fmt.Errorf("Failed to read bOnlyActive from parameters. %s", err.Error()), client, callID, 0, false)
	}

	go protocol.GetDetailedParticipantsHandler(nil, client, callID, idGathering, bOnlyActive)
}
