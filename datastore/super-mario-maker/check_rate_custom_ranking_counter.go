// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCheckRateCustomRankingCounter(packet nex.PacketInterface) {
	if protocol.CheckRateCustomRankingCounter == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::CheckRateCustomRankingCounter not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var applicationID types.UInt32

	err := applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CheckRateCustomRankingCounter(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, applicationID)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CheckRateCustomRankingCounter(nil, packet, callID, applicationID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
