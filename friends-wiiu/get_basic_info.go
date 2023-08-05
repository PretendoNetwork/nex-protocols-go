// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBasicInfo sets the GetBasicInfo handler function
func (protocol *Protocol) GetBasicInfo(handler func(err error, client *nex.Client, callID uint32, pids []uint32) uint32) {
	protocol.getBasicInfoHandler = handler
}

func (protocol *Protocol) handleGetBasicInfo(packet nex.PacketInterface) {
	if protocol.getBasicInfoHandler == nil {
		globals.Logger.Warning("FriendsWiiU::GetBasicInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.getBasicInfoHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getBasicInfoHandler(nil, client, callID, pids)
}
