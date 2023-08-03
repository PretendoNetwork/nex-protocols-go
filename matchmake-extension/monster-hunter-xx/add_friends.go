// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriends sets the AddFriends handler function
func (protocol *Protocol) AddFriends(handler func(err error, client *nex.Client, callID uint32, pids []uint64)) {
	protocol.addFriendsHandler = handler
}

func (protocol *Protocol) handleAddFriends(packet nex.PacketInterface) {
	if protocol.addFriendsHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::AddFriends not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.addFriendsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.addFriendsHandler(nil, client, callID, pids)
}
