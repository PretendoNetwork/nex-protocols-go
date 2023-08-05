// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCustomItem sets the GetCustomItem handler function
func (protocol *Protocol) GetCustomItem(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string) uint32) {
	protocol.getCustomItemHandler = handler
}

func (protocol *Protocol) handleGetCustomItem(packet nex.PacketInterface) {
	if protocol.getCustomItemHandler == nil {
		globals.Logger.Warning("PersistentStore::GetCustomItem not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getCustomItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		go protocol.getCustomItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	go protocol.getCustomItemHandler(nil, client, callID, uiGroup, strTag)
}
