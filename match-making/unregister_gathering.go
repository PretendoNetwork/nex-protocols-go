// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGathering sets the UnregisterGathering handler function
func (protocol *MatchMakingProtocol) UnregisterGathering(handler func(err error, client *nex.Client, callID uint32, idGathering uint32)) {
	protocol.UnregisterGatheringHandler = handler
}

func (protocol *MatchMakingProtocol) handleUnregisterGathering(packet nex.PacketInterface) {
	if protocol.UnregisterGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGathering not implemented")
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
		go protocol.UnregisterGatheringHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.UnregisterGatheringHandler(nil, client, callID, idGathering)
}
