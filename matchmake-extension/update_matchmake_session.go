// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMatchmakeSession sets the UpdateMatchmakeSession handler function
func (protocol *Protocol) UpdateMatchmakeSession(handler func(err error, packet nex.PacketInterface, callID uint32, anyGathering *nex.DataHolder) uint32) {
	protocol.updateMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleUpdateMatchmakeSession(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateMatchmakeSession not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.updateMatchmakeSessionHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateMatchmakeSessionHandler(nil, packet, callID, anyGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
