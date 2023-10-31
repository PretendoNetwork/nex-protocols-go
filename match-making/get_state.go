// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetState sets the GetState handler function
func (protocol *Protocol) GetState(handler func(err error, packet nex.PacketInterface, callID uint32, idGathering uint32) uint32) {
	protocol.getStateHandler = handler
}

func (protocol *Protocol) handleGetState(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStateHandler == nil {
		globals.Logger.Warning("MatchMaking::GetState not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getStateHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStateHandler(nil, packet, callID, idGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
