// Package match_making implements the Match Making NEX protocol
package match_making

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGatherings sets the UnregisterGatherings handler function
func (protocol *MatchMakingProtocol) UnregisterGatherings(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32)) {
	protocol.UnregisterGatheringsHandler = handler
}

func (protocol *MatchMakingProtocol) handleUnregisterGatherings(packet nex.PacketInterface) {
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

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.UnregisterGatheringsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.UnregisterGatheringsHandler(nil, client, callID, lstGatherings)
}
