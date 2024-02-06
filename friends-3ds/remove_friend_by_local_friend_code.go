// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleRemoveFriendByLocalFriendCode(packet nex.PacketInterface) {
	if protocol.RemoveFriendByLocalFriendCode == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::RemoveFriendByLocalFriendCode not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lfc := types.NewPrimitiveU64(0)

	err := lfc.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RemoveFriendByLocalFriendCode(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RemoveFriendByLocalFriendCode(nil, packet, callID, lfc)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
