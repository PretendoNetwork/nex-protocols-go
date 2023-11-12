// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Protocol) UpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_wiiu_types.MiiV2) uint32) {
	protocol.updateMiiHandler = handler
}

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateMiiHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateMii not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	miiV2, err := parametersStream.ReadStructure(friends_wiiu_types.NewMiiV2())
	if err != nil {
		errorCode = protocol.updateMiiHandler(fmt.Errorf("Failed to read miiV2 from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateMiiHandler(nil, packet, callID, miiV2.(*friends_wiiu_types.MiiV2))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
