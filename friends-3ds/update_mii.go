// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateMii sets the UpdateMii handler function
func (protocol *Friends3DSProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *friends_3ds_types.Mii)) {
	protocol.updateMiiHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdateMii(packet nex.PacketInterface) {
	if protocol.updateMiiHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateMii not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
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
