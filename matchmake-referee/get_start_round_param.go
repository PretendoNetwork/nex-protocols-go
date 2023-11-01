// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetStartRoundParam sets the GetStartRoundParam handler function
func (protocol *Protocol) GetStartRoundParam(handler func(err error, packet nex.PacketInterface, callID uint32, roundID uint64) uint32) {
	protocol.getStartRoundParamHandler = handler
}

func (protocol *Protocol) handleGetStartRoundParam(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getStartRoundParamHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetStartRoundParam not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	roundID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.getStartRoundParamHandler(fmt.Errorf("Failed to read roundID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getStartRoundParamHandler(nil, packet, callID, roundID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
