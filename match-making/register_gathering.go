// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterGathering sets the RegisterGathering handler function
func (protocol *Protocol) RegisterGathering(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder)) {
	protocol.registerGatheringHandler = handler
}

func (protocol *Protocol) handleRegisterGathering(packet nex.PacketInterface) {
	if protocol.registerGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterGathering not implemented")
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
		go protocol.registerGatheringHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil)
	}

	go protocol.registerGatheringHandler(nil, client, callID, anyGathering)
}
