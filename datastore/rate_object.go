// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRateObject(packet nex.PacketInterface) {
	if protocol.RateObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::RateObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	datastoreVersion := endpoint.LibraryVersions().DataStore
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	target := datastore_types.NewDataStoreRatingTarget()
	param := datastore_types.NewDataStoreRateObjectParam()
	var fetchRatings types.Bool

	var err error

	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObject(fmt.Errorf("failed to read target from parameters. %s", err.Error()), packet, callID, target, param, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObject(fmt.Errorf("failed to read param from parameters. %s", err.Error()), packet, callID, target, param, fetchRatings)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	// * When scanning all the DDLs for Wii U games, the following results were found:
	// *
	// * - The latest NEX version without this field is NEX 3.4.0 (or NEX 3.4.7 if including custom builds)
	// * - The earliest NEX version with this field is NEX 3.5.1
	// * - JOYSOUND for Wii U is a NEX 3.4.0 game with this field, but it also has custom modifications so likely an outlier
	// *
	// * So this is a best guess. The exact version of this addition is unknown, but this seems to be the most likely
	if datastoreVersion.GreaterOrEqual("3.5.0") {
		err = fetchRatings.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.RateObject(fmt.Errorf("failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, target, param, fetchRatings)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	} else {
		// * Prior to this field being added, ratings were always returned no matter what
		fetchRatings = true
	}

	rmcMessage, rmcError := protocol.RateObject(nil, packet, callID, target, param, fetchRatings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
