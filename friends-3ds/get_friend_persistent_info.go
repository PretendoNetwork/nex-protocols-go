// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetFriendPersistentInfo sets the GetFriendPersistentInfo handler function
func (protocol *Protocol) GetFriendPersistentInfo(handler func(err error, client *nex.Client, callID uint32, pidList []uint32) uint32) {
	protocol.getFriendPersistentInfoHandler = handler
}

func (protocol *Protocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if protocol.getFriendPersistentInfoHandler == nil {
		globals.Logger.Warning("Friends3DS::GetFriendPersistentInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidList, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getFriendPersistentInfoHandler(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getFriendPersistentInfoHandler(nil, client, callID, pidList)
}
