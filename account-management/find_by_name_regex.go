// Package account_management implements the Account Management NEX protocol
package account_management

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// FindByNameRegex sets the FindByNameRegex handler function
func (protocol *AccountManagementProtocol) FindByNameRegex(handler func(err error, client *nex.Client, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange)) {
	protocol.findByNameRegexHandler = handler
}

func (protocol *AccountManagementProtocol) handleFindByNameRegex(packet nex.PacketInterface) {
	if protocol.findByNameRegexHandler == nil {
		globals.Logger.Warning("AccountManagement::FindByNameRegex not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiGroups, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.findByNameRegexHandler(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	strRegex, err := parametersStream.ReadString()
	if err != nil {
		go protocol.findByNameRegexHandler(fmt.Errorf("Failed to read strRegex from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	resultRange, err := parametersStream.ReadStructure(nex.NewResultRange())
	if err != nil {
		go protocol.findByNameRegexHandler(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), client, callID, 0, "", nil)
		return
	}

	go protocol.findByNameRegexHandler(nil, client, callID, uiGroups, strRegex, resultRange.(*nex.ResultRange))
}
