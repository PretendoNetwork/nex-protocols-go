// Package persistent_store implements the NAT Traversal NEX protocol
package persistent_store

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InsertCustomItem sets the InsertCustomItem handler function
func (protocol *PersistentStoreProtocol) InsertCustomItem(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool)) {
	protocol.insertCustomItemHandler = handler
}

func (protocol *PersistentStoreProtocol) handleInsertCustomItem(packet nex.PacketInterface) {
	if protocol.insertCustomItemHandler == nil {
		globals.Logger.Warning("PersistentStore::InsertCustomItem not implemented")
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
		go protocol.insertCustomItemHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0, "", nil, false)
		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		go protocol.insertCustomItemHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), client, callID, 0, "", nil, false)
		return
	}

	hData, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.insertCustomItemHandler(fmt.Errorf("Failed to read hData from parameters. %s", err.Error()), client, callID, 0, "", nil, false)
		return
	}

	bReplace, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.insertCustomItemHandler(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), client, callID, 0, "", nil, false)
		return
	}

	go protocol.insertCustomItemHandler(nil, client, callID, uiGroup, strTag, hData, bReplace)
}
