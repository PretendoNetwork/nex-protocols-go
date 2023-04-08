package match_making

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGatherings sets the UnregisterGatherings handler function
func (protocol *MatchMakingProtocol) UnregisterGatherings(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	protocol.UnregisterGatheringsHandler = handler
}

func (protocol *MatchMakingProtocol) HandleUnregisterGatherings(packet nex.PacketInterface) {
	if protocol.UnregisterGatheringsHandler == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGatherings not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings := parametersStream.ReadListUInt32LE()

	go protocol.UnregisterGatheringsHandler(nil, client, callID, lstGatherings)
}
