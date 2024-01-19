// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRequestMigration(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RequestMigration == nil {
		globals.Logger.Warning("DataStorePokemonBank::RequestMigration not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	oneTimePassword := types.NewString("")
	err = oneTimePassword.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestMigration(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	boxes := types.NewList[*types.PrimitiveU32]()
	boxes.Type = types.NewPrimitiveU32(0)
	err = boxes.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestMigration(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestMigration(nil, packet, callID, oneTimePassword, boxes)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
