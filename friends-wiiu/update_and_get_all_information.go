// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	var err error

	if protocol.UpdateAndGetAllInformation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::UpdateAndGetAllInformation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	nnaInfo := friends_wiiu_types.NewNNAInfo()
	err = nnaInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read nnaInfo from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	presence := friends_wiiu_types.NewNintendoPresenceV2()
	err = presence.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read presence from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	birthday := types.NewDateTime(0)
	err = birthday.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read birthday from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAndGetAllInformation(nil, packet, callID, nnaInfo, presence, birthday)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
