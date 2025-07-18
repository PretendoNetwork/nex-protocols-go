// Package protocol implements the StorageManager protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleAcquireNexUniqueID(packet nex.PacketInterface) {
	if protocol.AcquireNexUniqueID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "StorageManager::AcquireNexUniqueID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var slot types.UInt8

	var err error

	err = slot.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AcquireNexUniqueID(fmt.Errorf("Failed to read slot from parameters. %s", err.Error()), packet, callID, slot)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AcquireNexUniqueID(nil, packet, callID, slot)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
