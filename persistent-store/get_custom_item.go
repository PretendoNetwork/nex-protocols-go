// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetCustomItem(packet nex.PacketInterface) {
	if protocol.GetCustomItem == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "PersistentStore::GetCustomItem not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroup := types.NewPrimitiveU32(0)
	strTag := types.NewString("")

	var err error

	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCustomItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCustomItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCustomItem(nil, packet, callID, uiGroup, strTag)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
