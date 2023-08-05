// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriendByLocalFriendCode sets the RemoveFriendByLocalFriendCode handler function
func (protocol *Protocol) RemoveFriendByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64) uint32) {
	protocol.removeFriendByLocalFriendCodeHandler = handler
}

func (protocol *Protocol) handleRemoveFriendByLocalFriendCode(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.removeFriendByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByLocalFriendCode not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.removeFriendByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.removeFriendByLocalFriendCodeHandler(nil, client, callID, lfc)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
