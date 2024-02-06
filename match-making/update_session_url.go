// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateSessionURL(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateSessionURL == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::UpdateSessionURL not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateSessionURL(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strURL := types.NewString("")
	err = strURL.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateSessionURL(fmt.Errorf("Failed to read strURL from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateSessionURL(nil, packet, callID, idGathering, strURL)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
