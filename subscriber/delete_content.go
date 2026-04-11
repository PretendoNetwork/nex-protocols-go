// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDeleteContent(packet nex.PacketInterface) {
	if protocol.DeleteContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::DeleteContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var topics types.List[types.String]
	var contentID types.UInt64

	var err error

	err = topics.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteContent(fmt.Errorf("Failed to read topics from parameters. %s", err.Error()), packet, callID, topics, contentID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = contentID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteContent(fmt.Errorf("Failed to read contentID from parameters. %s", err.Error()), packet, callID, topics, contentID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteContent(nil, packet, callID, topics, contentID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
