// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Protocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *friends_3ds_types.Mii) uint32) {
	protocol.updateMiiHandler = handler
}

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	if protocol.updateMiiHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMii not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mii, err := parametersStream.ReadStructure(friends_3ds_types.NewMii())
	if err != nil {
		go protocol.updateMiiHandler(fmt.Errorf("Failed to read mii from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateMiiHandler(nil, client, callID, mii.(*friends_3ds_types.Mii))
}
