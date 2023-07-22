// Package friends_3ds implements the Friends 3DS NEX protocol
package friends_3ds

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateBlackList sets the UpdateBlackList handler function
func (protocol *Friends3DSProtocol) UpdateBlackList(handler func(err error, client *nex.Client, callID uint32, unknown []uint32)) {
	protocol.updateBlackListHandler = handler
}

func (protocol *Friends3DSProtocol) handleUpdateBlackList(packet nex.PacketInterface) {
	if protocol.updateBlackListHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateBlackList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.updateBlackListHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateBlackListHandler(nil, client, callID, unknown)
}
