// Package persistent_store implements the NAT Traversal NEX protocol
package persistent_store

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetItem sets the GetItem handler function
func (protocol *PersistentStoreProtocol) GetItem(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string)) {
	protocol.getItemHandler = handler
}

func (protocol *PersistentStoreProtocol) handleGetItem(packet nex.PacketInterface) {
	if protocol.getItemHandler == nil {
		globals.Logger.Warning("PersistentStore::GetItem not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroup, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		go protocol.getItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	go protocol.getItemHandler(nil, client, callID, uiGroup, strTag)
}
