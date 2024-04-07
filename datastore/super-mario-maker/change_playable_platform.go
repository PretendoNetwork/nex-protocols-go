// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleChangePlayablePlatform(packet nex.PacketInterface) {
	if protocol.ChangePlayablePlatform == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::ChangePlayablePlatform not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	params := types.NewList[*datastore_super_mario_maker_types.DataStoreChangePlayablePlatformParam]()
	params.Type = datastore_super_mario_maker_types.NewDataStoreChangePlayablePlatformParam()

	err := params.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ChangePlayablePlatform(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ChangePlayablePlatform(nil, packet, callID, params)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
