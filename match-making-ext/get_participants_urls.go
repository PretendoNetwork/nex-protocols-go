package match_making_ext

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (protocol *MatchMakingExtProtocol) GetParticipantsURLs(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	protocol.GetParticipantsURLsHandler = handler
}

func (protocol *MatchMakingExtProtocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	if protocol.GetParticipantsURLsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::GetParticipantsURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.GetParticipantsURLsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.GetParticipantsURLsHandler(nil, client, callID, lstGatherings)
}
