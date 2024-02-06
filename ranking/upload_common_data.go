// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUploadCommonData(packet nex.PacketInterface) {
	if protocol.UploadCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::UploadCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	commonData := types.NewBuffer(nil)
	uniqueID := types.NewPrimitiveU64(0)

	var err error

	err = commonData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadCommonData(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadCommonData(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadCommonData(nil, packet, callID, commonData, uniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
