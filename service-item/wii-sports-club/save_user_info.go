// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// SaveUserInfo sets the SaveUserInfo handler function
func (protocol *Protocol) SaveUserInfo(handler func(err error, client *nex.Client, callID uint32, userInfo *service_item_wii_sports_club_types.ServiceItemUserInfo) uint32) {
	protocol.saveUserInfoHandler = handler
}

func (protocol *Protocol) handleSaveUserInfo(packet nex.PacketInterface) {
	if protocol.saveUserInfoHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::SaveUserInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	userInfo, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemUserInfo())
	if err != nil {
		go protocol.saveUserInfoHandler(fmt.Errorf("Failed to read userInfo from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.saveUserInfoHandler(nil, client, callID, userInfo.(*service_item_wii_sports_club_types.ServiceItemUserInfo))
}
