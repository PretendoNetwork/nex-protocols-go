// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleTestCapability(packet nex.PacketInterface) {
	if protocol.TestCapability == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::TestCapability not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uiCapability types.UInt32

	err := uiCapability.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.TestCapability(fmt.Errorf("Failed to read uiCapability from parameters. %s", err.Error()), packet, callID, uiCapability)
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
