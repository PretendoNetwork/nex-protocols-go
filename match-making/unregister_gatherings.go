// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterGatherings sets the UnregisterGatherings handler function
func (protocol *Protocol) UnregisterGatherings(handler func(err error, packet nex.PacketInterface, callID uint32, lstGatherings []uint32) uint32) {
	protocol.unregisterGatheringsHandler = handler
}

func (protocol *Protocol) handleUnregisterGatherings(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.unregisterGatheringsHandler == nil {
		globals.Logger.Warning("MatchMaking::UnregisterGatherings not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstGatherings, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.unregisterGatheringsHandler(fmt.Errorf("Failed to read lstGatherings from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.unregisterGatheringsHandler(nil, packet, callID, lstGatherings)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
