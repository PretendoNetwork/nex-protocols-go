// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGatherings sets the UnregisterGatherings handler function
func (protocol *Protocol) UnregisterGatherings(handler func(err error, client *nex.Client, callID uint32, lstGatherings []uint32) uint32) {
	protocol.unregisterGatheringsHandler = handler
}

func (protocol *Protocol) handleUnregisterGatherings(packet nex.PacketInterface) {
	if protocol.unregisterGatheringsHandler == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGatherings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.unregisterGatheringsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.unregisterGatheringsHandler(nil, client, callID, lstGatherings)
}
