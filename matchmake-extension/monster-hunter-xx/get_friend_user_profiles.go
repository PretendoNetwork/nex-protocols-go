// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendUserProfiles sets the GetFriendUserProfiles handler function
func (protocol *Protocol) GetFriendUserProfiles(handler func(err error, packet nex.PacketInterface, callID uint32, pids []uint64) uint32) {
	protocol.getFriendUserProfilesHandler = handler
}

func (protocol *Protocol) handleGetFriendUserProfiles(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getFriendUserProfilesHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::GetFriendUserProfiles not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		errorCode = protocol.getFriendUserProfilesHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getFriendUserProfilesHandler(nil, packet, callID, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
