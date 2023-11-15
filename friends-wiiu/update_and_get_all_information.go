// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateAndGetAllInformation == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateAndGetAllInformation not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nnaInfo, err := parametersStream.ReadStructure(friends_wiiu_types.NewNNAInfo())
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read nnaInfo from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	presence, err := parametersStream.ReadStructure(friends_wiiu_types.NewNintendoPresenceV2())
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read presence from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	birthday, err := parametersStream.ReadDateTime()
	if err != nil {
		_, errorCode = protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read birthday from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateAndGetAllInformation(nil, packet, callID, nnaInfo.(*friends_wiiu_types.NNAInfo), presence.(*friends_wiiu_types.NintendoPresenceV2), birthday)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
