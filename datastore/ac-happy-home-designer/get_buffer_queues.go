// Package protocol implements the Animal Crossing: Happy Home Designer protocol
package protocol

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_ac_happy_home_designer_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/ac-happy-home-designer/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetBufferQueues(packet nex.PacketInterface) {
	if protocol.AddToBufferQueue == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreACHappyHomeDesigner::GetBufferQueues not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	params := types.NewList[*datastore_ac_happy_home_designer_types.BufferQueueParam]()
	params.Type = datastore_ac_happy_home_designer_types.NewBufferQueueParam()

	err := params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetBufferQueues(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetBufferQueues(nil, packet, callID, params)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
