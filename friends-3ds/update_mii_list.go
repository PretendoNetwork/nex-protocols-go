// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMiiList sets the UpdateMiiList handler function
func (protocol *Protocol) UpdateMiiList(handler func(err error, client *nex.Client, callID uint32, miiList *friends_3ds_types.MiiList)) {
	protocol.updateMiiListHandler = handler
}

func (protocol *Protocol) handleUpdateMiiList(packet nex.PacketInterface) {
	if protocol.updateMiiListHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMiiList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	miiList, err := parametersStream.ReadStructure(friends_3ds_types.NewMiiList())
	if err != nil {
		go protocol.updateMiiListHandler(fmt.Errorf("Failed to read miiList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateMiiListHandler(nil, client, callID, miiList.(*friends_3ds_types.MiiList))
}
