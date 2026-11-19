// Package protocol implements the Uplay Win protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetRewards(packet nex.PacketInterface) {
	if protocol.GetRewards == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "UplayWin::GetRewards not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var startRowIndex types.Int32
	var maximumRows types.Int32
	var sortExpression types.String
	var cultureName types.String
	var platformCode types.String
	var gameCode types.String

	if err := startRowIndex.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read startRowIndex: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := maximumRows.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read maximumRows: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := sortExpression.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read sortExpression: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := cultureName.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read cultureName: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := platformCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read platformCode: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	if err := gameCode.ExtractFrom(parametersStream); err != nil {
		_, rmcError := protocol.GetRewards(fmt.Errorf("failed to read gameCode: %s", err.Error()), packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}
		return
	}

	rmcMessage, rmcError := protocol.GetRewards(nil, packet, callID, startRowIndex, maximumRows, sortExpression, cultureName, platformCode, gameCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
