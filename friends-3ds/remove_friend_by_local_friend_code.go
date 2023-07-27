// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriendByLocalFriendCode sets the RemoveFriendByLocalFriendCode handler function
func (protocol *Protocol) RemoveFriendByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64)) {
	protocol.removeFriendByLocalFriendCodeHandler = handler
}

func (protocol *Protocol) handleRemoveFriendByLocalFriendCode(packet nex.PacketInterface) {
	if protocol.removeFriendByLocalFriendCodeHandler == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByLocalFriendCode not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.removeFriendByLocalFriendCodeHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.removeFriendByLocalFriendCodeHandler(nil, client, callID, lfc)
}
