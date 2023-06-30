// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProfile sets the UpdateProfile handler function
func (protocol *Friends3DSProtocol) UpdateProfile(handler func(err error, client *nex.Client, callID uint32, profileData *friends_3ds_types.MyProfile)) {
	protocol.UpdateProfileHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdateProfile(packet nex.PacketInterface) {
	if protocol.UpdateProfileHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateProfile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	profileData, err := parametersStream.ReadStructure(friends_3ds_types.NewMyProfile())
	if err != nil {
		go protocol.UpdateProfileHandler(fmt.Errorf("Failed to read showGame from profileData. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdateProfileHandler(nil, client, callID, profileData.(*friends_3ds_types.MyProfile))
}
