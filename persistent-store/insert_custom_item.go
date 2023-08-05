// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// InsertCustomItem sets the InsertCustomItem handler function
func (protocol *Protocol) InsertCustomItem(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, hData *nex.DataHolder, bReplace bool) uint32) {
	protocol.insertCustomItemHandler = handler
}

func (protocol *Protocol) handleInsertCustomItem(packet nex.PacketInterface) {
	if protocol.insertCustomItemHandler == nil {
		globals.Logger.Warning("PersistentStore::InsertCustomItem not implemented")
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
