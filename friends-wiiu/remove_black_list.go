// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveBlackList sets the RemoveBlackList handler function
func (protocol *Protocol) RemoveBlackList(handler func(err error, client *nex.Client, callID uint32, pid uint32) uint32) {
	protocol.removeBlackListHandler = handler
}

func (protocol *Protocol) handleRemoveBlackList(packet nex.PacketInterface) {
	if protocol.removeBlackListHandler == nil {
		globals.Logger.Warning("FriendsWiiU::RemoveBlackList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.removeBlackListHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.removeBlackListHandler(nil, client, callID, pid)
}
