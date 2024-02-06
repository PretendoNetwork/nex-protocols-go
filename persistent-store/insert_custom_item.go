// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleInsertCustomItem(packet nex.PacketInterface) {
	var err error

	if protocol.InsertCustomItem == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "PersistentStore::InsertCustomItem not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroup := types.NewPrimitiveU32(0)
	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertCustomItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strTag := types.NewString("")
	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertCustomItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	hData := types.NewAnyDataHolder()
	err = hData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertCustomItem(fmt.Errorf("Failed to read hData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	bReplace := types.NewPrimitiveBool(false)
	err = bReplace.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertCustomItem(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.InsertCustomItem(nil, packet, callID, uiGroup, strTag, hData, bReplace)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
