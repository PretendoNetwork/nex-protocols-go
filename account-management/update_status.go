// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateStatus sets the UpdateStatus handler function
func (protocol *Protocol) UpdateStatus(handler func(err error, client *nex.Client, callID uint32, strStatus string)) {
	protocol.updateStatusHandler = handler
}

func (protocol *Protocol) handleUpdateStatus(packet nex.PacketInterface) {
	if protocol.updateStatusHandler == nil {
		globals.Logger.Warning("AccountManagement::UpdateStatus not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strStatus, err := parametersStream.ReadString()
	if err != nil {
		go protocol.updateStatusHandler(fmt.Errorf("Failed to read strStatus from parameters. %s", err.Error()), client, callID, "")
		return
	}

	go protocol.updateStatusHandler(nil, client, callID, strStatus)
}
