// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetUserStatuses sets the GetUserStatuses handler function
func (protocol *Protocol) GetUserStatuses(handler func(err error, client *nex.Client, callID uint32, pids []uint32, unknown []uint8) uint32) {
	protocol.getUserStatusesHandler = handler
}

func (protocol *Protocol) handleGetUserStatuses(packet nex.PacketInterface) {
	if protocol.getUserStatusesHandler == nil {
		globals.Logger.Warning("Subscriber::GetUserStatuses not implemented")
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
		go protocol.getUserStatusesHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	unknown, err := parametersStream.ReadListUInt8()
	if err != nil {
		go protocol.getUserStatusesHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.getUserStatusesHandler(nil, client, callID, pids, unknown)
}
