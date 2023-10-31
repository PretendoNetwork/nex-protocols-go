// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateGathering sets the UpdateGathering handler function
func (protocol *Protocol) UpdateGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32) {
	protocol.updateGatheringHandler = handler
}

func (protocol *Protocol) handleUpdateGathering(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::UpdateGathering not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateGatheringHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateGatheringHandler(nil, packet, callID, anyGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
