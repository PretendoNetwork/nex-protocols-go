// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByName sets the AddFriendByName handler function
func (protocol *Protocol) AddFriendByName(handler func(err error, packet nex.PacketInterface, callID uint32, username string) uint32) {
	protocol.addFriendByNameHandler = handler
}

func (protocol *Protocol) handleAddFriendByName(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendByNameHandler == nil {
		globals.Logger.Warning("FriendsWiiU::AddFriendByName not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	username, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendByNameHandler(fmt.Errorf("Failed to read username from parameters. %s", err.Error()), packet, callID, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendByNameHandler(nil, packet, callID, username)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
