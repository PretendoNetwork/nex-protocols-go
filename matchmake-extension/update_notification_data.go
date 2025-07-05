// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	if protocol.UpdateNotificationData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateNotificationData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	libraryVersion := endpoint.LibraryVersions().Main

	var uiType types.UInt32
	var uiParam1 types.UInt64
	var uiParam2 types.UInt64
	var strParam types.String

	var err error

	err = uiType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, uiType, uiParam1, uiParam2, strParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = uiParam1.ExtractFrom(parametersStream)
	} else {
		var Param1 types.UInt32
		err = Param1.ExtractFrom(parametersStream)
		uiParam1 = types.NewUInt64(uint64(Param1))
	}

	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam1 from parameters. %s", err.Error()), packet, callID, uiType, uiParam1, uiParam2, strParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = uiParam2.ExtractFrom(parametersStream)
	} else {
		var Param2 types.UInt32
		err = Param2.ExtractFrom(parametersStream)
		uiParam2 = types.NewUInt64(uint64(Param2))
	}

	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam2 from parameters. %s", err.Error()), packet, callID, uiType, uiParam1, uiParam2, strParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read strParam from parameters. %s", err.Error()), packet, callID, uiType, uiParam1, uiParam2, strParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateNotificationData(nil, packet, callID, uiType, uiParam1, uiParam2, strParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
