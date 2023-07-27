// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendRelationships sets the GetFriendRelationships handler function
func (protocol *Protocol) GetFriendRelationships(handler func(err error, client *nex.Client, callID uint32, pids []uint32)) {
	protocol.getFriendRelationshipsHandler = handler
}

func (protocol *Protocol) handleGetFriendRelationships(packet nex.PacketInterface) {
	if protocol.getFriendRelationshipsHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendRelationships not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getFriendRelationshipsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendRelationshipsHandler(nil, client, callID, pids)
}
