// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateGathering sets the UpdateGathering handler function
func (protocol *Protocol) UpdateGathering(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder)) {
	protocol.updateGatheringHandler = handler
}

func (protocol *Protocol) handleUpdateGathering(packet nex.PacketInterface) {
	if protocol.updateGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::UpdateGathering not implemented")
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
		go protocol.updateGatheringHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.updateGatheringHandler(nil, client, callID, anyGathering)
}
