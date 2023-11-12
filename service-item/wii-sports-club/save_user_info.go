// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// SaveUserInfo sets the SaveUserInfo handler function
func (protocol *Protocol) SaveUserInfo(handler func(err error, packet nex.PacketInterface, callID uint32, userInfo *service_item_wii_sports_club_types.ServiceItemUserInfo) uint32) {
	protocol.saveUserInfoHandler = handler
}

func (protocol *Protocol) handleSaveUserInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.saveUserInfoHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::SaveUserInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	userInfo, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemUserInfo())
	if err != nil {
		errorCode = protocol.saveUserInfoHandler(fmt.Errorf("Failed to read userInfo from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.saveUserInfoHandler(nil, packet, callID, userInfo.(*service_item_wii_sports_club_types.ServiceItemUserInfo))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
