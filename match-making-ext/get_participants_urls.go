package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (protocol *MatchMakingExtProtocol) GetParticipantsURLs(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	protocol.GetParticipantsURLsHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleGetParticipantsURLs(packet nex.PacketInterface) {
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

	lstGatheringsCount := parametersStream.ReadUInt32LE()
	lstGatherings := make([]uint32, lstGatheringsCount)
	for i := 0; uint32(i) < lstGatheringsCount; i++ {
		lstGatherings[i] = parametersStream.ReadUInt32LE()
	}

	go protocol.GetParticipantsURLsHandler(nil, client, callID, lstGatherings)
}
