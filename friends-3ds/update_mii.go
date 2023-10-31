// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Protocol) UpdateMii(handler func(err error, packet nex.PacketInterface, callID uint32, mii *friends_3ds_types.Mii) uint32) {
	protocol.updateMiiHandler = handler
}

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateMiiHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMii not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	mii, err := parametersStream.ReadStructure(friends_3ds_types.NewMii())
	if err != nil {
		errorCode = protocol.updateMiiHandler(fmt.Errorf("Failed to read mii from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateMiiHandler(nil, packet, callID, mii.(*friends_3ds_types.Mii))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
