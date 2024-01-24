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
	var errorCode uint32

	if protocol.InsertCustomItem == nil {
		globals.Logger.Warning("PersistentStore::InsertCustomItem not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroup := types.NewPrimitiveU32(0)
	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag := types.NewString("")
	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	hData := types.NewAnyDataHolder()
	err = hData.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read hData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	bReplace := types.NewPrimitiveBool(false)
	err = bReplace.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.InsertCustomItem(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.InsertCustomItem(nil, packet, callID, uiGroup, strTag, hData, bReplace)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
