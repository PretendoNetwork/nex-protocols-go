// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleTestCapability(packet nex.PacketInterface) {
	var err error

	if protocol.TestCapability == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::TestCapability not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiCapability := types.NewPrimitiveU32(0)
	err = uiCapability.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.TestCapability(fmt.Errorf("Failed to read uiCapability from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.TestCapability(nil, packet, callID, uiCapability)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
