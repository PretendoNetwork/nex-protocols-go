// Package protocol implements the Miitopia DataStore protocol
package protocol

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2"
	datastore_miitopia_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/miitopia/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSearchMii(packet nex.PacketInterface) {
	if protocol.SearchMii == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreMiitopia::SearchMii not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	searchParam := datastore_miitopia_types.NewMiiTubeSearchParam()

	err := searchParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SearchMii(fmt.Errorf("Failed to read searchParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SearchMii(nil, packet, callID, searchParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
