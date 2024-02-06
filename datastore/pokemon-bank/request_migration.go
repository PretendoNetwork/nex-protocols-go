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

	if protocol.RequestMigration == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::RequestMigration not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	oneTimePassword := types.NewString("")
	err = oneTimePassword.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestMigration(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	boxes := types.NewList[*types.PrimitiveU32]()
	boxes.Type = types.NewPrimitiveU32(0)
	err = boxes.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestMigration(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestMigration(nil, packet, callID, oneTimePassword, boxes)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
