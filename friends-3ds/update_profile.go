// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProfile sets the UpdateProfile handler function
func (protocol *Protocol) UpdateProfile(handler func(err error, packet nex.PacketInterface, callID uint32, profileData *friends_3ds_types.MyProfile) uint32) {
	protocol.updateProfileHandler = handler
}

func (protocol *Protocol) handleUpdateProfile(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateProfileHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateProfile not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	profileData, err := parametersStream.ReadStructure(friends_3ds_types.NewMyProfile())
	if err != nil {
		errorCode = protocol.updateProfileHandler(fmt.Errorf("Failed to read showGame from profileData. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateProfileHandler(nil, packet, callID, profileData.(*friends_3ds_types.MyProfile))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
