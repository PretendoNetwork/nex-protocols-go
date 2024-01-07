// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveItem(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RemoveItem == nil {
		globals.Logger.Warning("PersistentStore::RemoveItem not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiGroup := types.NewPrimitiveU32(0)
	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RemoveItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strTag := types.NewString("")
	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RemoveItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RemoveItem(nil, packet, callID, uiGroup, strTag)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
