// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetState sets the SetState handler function
func (protocol *Protocol) SetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32, uiNewState uint32) uint32) {
	protocol.setStateHandler = handler
}

func (protocol *Protocol) handleSetState(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setStateHandler == nil {
		globals.Logger.Warning("MatchMaking::SetState not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setStateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiNewState, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setStateHandler(fmt.Errorf("Failed to read uiNewState from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.setStateHandler(nil, packet, callID, idGathering, uiNewState)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
