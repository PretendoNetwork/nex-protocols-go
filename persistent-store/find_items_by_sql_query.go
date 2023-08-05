// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindItemsBySQLQuery sets the FindItemsBySQLQuery handler function
func (protocol *Protocol) FindItemsBySQLQuery(handler func(err error, client *nex.Client, callID uint32, uiGroup uint32, strTag string, strQuery string) uint32) {
	protocol.findItemsBySQLQueryHandler = handler
}

func (protocol *Protocol) handleFindItemsBySQLQuery(packet nex.PacketInterface) {
	if protocol.findItemsBySQLQueryHandler == nil {
		globals.Logger.Warning("PersistentStore::FindItemsBySQLQuery not implemented")
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
		go protocol.findItemsBySQLQueryHandler(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), client, callID, 0, "", "")
		return
	}

	strTag, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findItemsBySQLQueryHandler(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), client, callID, 0, "", "")
		return
	}

	strQuery, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findItemsBySQLQueryHandler(fmt.Errorf("Failed to read strQuery from parameters. %s", err.Error()), client, callID, 0, "", "")
		return
	}

	go protocol.findItemsBySQLQueryHandler(nil, client, callID, uiGroup, strTag, strQuery)
}
