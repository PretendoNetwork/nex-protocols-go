// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMiiList sets the UpdateMiiList handler function
func (protocol *Protocol) UpdateMiiList(handler func(err error, packet nex.PacketInterface, callID uint32, miiList *friends_3ds_types.MiiList) uint32) {
	protocol.updateMiiListHandler = handler
}

func (protocol *Protocol) handleUpdateMiiList(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateMiiListHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMiiList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	miiList, err := parametersStream.ReadStructure(friends_3ds_types.NewMiiList())
	if err != nil {
		errorCode = protocol.updateMiiListHandler(fmt.Errorf("Failed to read miiList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateMiiListHandler(nil, packet, callID, miiList.(*friends_3ds_types.MiiList))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
