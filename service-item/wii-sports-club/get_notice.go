// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// GetNotice sets the GetNotice handler function
func (protocol *Protocol) GetNotice(handler func(err error, packet nex.PacketInterface, callID uint32, getNoticeParam *service_item_wii_sports_club_types.ServiceItemGetNoticeParam) uint32) {
	protocol.getNoticeHandler = handler
}

func (protocol *Protocol) handleGetNotice(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNoticeHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetNotice not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getNoticeParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemGetNoticeParam())
	if err != nil {
		errorCode = protocol.getNoticeHandler(fmt.Errorf("Failed to read getNoticeParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getNoticeHandler(nil, packet, callID, getNoticeParam.(*service_item_wii_sports_club_types.ServiceItemGetNoticeParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
