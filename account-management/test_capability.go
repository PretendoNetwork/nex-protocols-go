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
	var errorCode uint32

	if protocol.TestCapability == nil {
		globals.Logger.Warning("AccountManagement::TestCapability not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiCapability := types.NewPrimitiveU32(0)
	err = uiCapability.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.TestCapability(fmt.Errorf("Failed to read uiCapability from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.TestCapability(nil, packet, callID, uiCapability)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
