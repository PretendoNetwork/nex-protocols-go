// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Protocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *friends_wiiu_types.MiiV2) uint32) {
	protocol.updateMiiHandler = handler
}

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	if protocol.updateMiiHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdateMii not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	miiV2, err := parametersStream.ReadStructure(friends_wiiu_types.NewMiiV2())
	if err != nil {
		go protocol.updateMiiHandler(fmt.Errorf("Failed to read miiV2 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateMiiHandler(nil, client, callID, miiV2.(*friends_wiiu_types.MiiV2))
}
