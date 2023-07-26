// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMatchmakeSession sets the UpdateMatchmakeSession handler function
func (protocol *Protocol) UpdateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder)) {
	protocol.updateMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleUpdateMatchmakeSession(packet nex.PacketInterface) {
	if protocol.updateMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.updateMatchmakeSessionHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateMatchmakeSessionHandler(nil, client, callID, anyGathering)
}
