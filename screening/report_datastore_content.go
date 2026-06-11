// Package protocol implements the Screening protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	screening_types "github.com/PretendoNetwork/nex-protocols-go/v2/screening/types"
)

func (protocol *Protocol) handleReportDataStoreContent(packet nex.PacketInterface) {
	if protocol.ReportDataStoreContent == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Screening::ReportDataStoreContent not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var pContentParam screening_types.ScreeningDataStoreContentParam
	var pViolationParam screening_types.ScreeningUGCViolationParam

	err := pContentParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportDataStoreContent(fmt.Errorf("failed to read pContentParam from parameters. %s", err.Error()), packet, callID, pContentParam, pViolationParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = pViolationParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportDataStoreContent(fmt.Errorf("failed to read pViolationParam from parameters. %s", err.Error()), packet, callID, pContentParam, pViolationParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportDataStoreContent(nil, packet, callID, pContentParam, pViolationParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
