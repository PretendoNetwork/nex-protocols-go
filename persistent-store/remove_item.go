// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRemoveItem(packet nex.PacketInterface) {
	if protocol.RemoveItem == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "PersistentStore::RemoveItem not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uiGroup types.UInt32
	var strTag types.String

	var err error

	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RemoveItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, uiGroup, strTag)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RemoveItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, uiGroup, strTag)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveItem(nil, packet, callID, uiGroup, strTag)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
