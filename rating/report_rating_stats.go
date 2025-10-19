// Package protocol implements the Rating protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"

	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	rating_types "github.com/PretendoNetwork/nex-protocols-go/v2/rating/types"
)

func (protocol *Protocol) handleReportRatingStats(packet nex.PacketInterface) {
	if protocol.ReportRatingStats == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Rating::ReportRatingStats not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var sessionToken rating_types.RatingSessionToken
	var stats types.List[rating_types.RatingStats]

	var err error

	err = sessionToken.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportRatingStats(fmt.Errorf("Failed to read sessionToken from parameters. %s", err.Error()), packet, callID, sessionToken, stats)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = stats.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportRatingStats(fmt.Errorf("Failed to read stats from parameters. %s", err.Error()), packet, callID, sessionToken, stats)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportRatingStats(nil, packet, callID, sessionToken, stats)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
