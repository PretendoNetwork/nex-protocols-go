// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterGathering sets the RegisterGathering handler function
func (protocol *Protocol) RegisterGathering(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32) {
	protocol.registerGatheringHandler = handler
}

func (protocol *Protocol) handleRegisterGathering(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::RegisterGathering not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.registerGatheringHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.registerGatheringHandler(nil, packet, callID, anyGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
