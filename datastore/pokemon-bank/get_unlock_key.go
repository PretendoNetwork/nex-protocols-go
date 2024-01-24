// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetUnlockKey(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetUnlockKey == nil {
		globals.Logger.Warning("DataStorePokemonBank::GetUnlockKey not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	challengeValue := types.NewPrimitiveU32(0)
	err = challengeValue.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetUnlockKey(fmt.Errorf("Failed to read challengeValue from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetUnlockKey(nil, packet, callID, challengeValue)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
