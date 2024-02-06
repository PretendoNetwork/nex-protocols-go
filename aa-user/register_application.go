// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRegisterApplication(packet nex.PacketInterface) {
	if protocol.RegisterApplication == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AAUser::RegisterApplication not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	titleID := types.NewPrimitiveU64(0)

	err := titleID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RegisterApplication(fmt.Errorf("Failed to read titleID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RegisterApplication(nil, packet, callID, titleID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
