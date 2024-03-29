// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriends sets the AddFriends handler function
func (protocol *Protocol) AddFriends(handler func(err error, packet nex.PacketInterface, callID uint32, pids []uint64) uint32) {
	protocol.addFriendsHandler = handler
}

func (protocol *Protocol) handleAddFriends(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::AddFriends not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.addFriendsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendsHandler(nil, packet, callID, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
