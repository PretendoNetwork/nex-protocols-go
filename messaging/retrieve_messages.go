// Package messaging implements the Messaging protocol
package messaging

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/messaging/types"
)

// RetrieveMessages sets the RetrieveMessages handler function
func (protocol *MessagingProtocol) RetrieveMessages(handler func(err error, client *nex.Client, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs []uint32, bLeaveOnServer bool)) {
	protocol.retrieveMessagesHandler = handler
}

func (protocol *MessagingProtocol) handleRetrieveMessages(packet nex.PacketInterface) {
	if protocol.retrieveMessagesHandler == nil {
		globals.Logger.Warning("Messaging::RetrieveMessages not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	recipient, err := parametersStream.ReadStructure(messaging_types.NewMessageRecipient())
	if err != nil {
		go protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	lstMsgIDs, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read lstMsgIDs from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	bLeaveOnServer, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.retrieveMessagesHandler(fmt.Errorf("Failed to read bLeaveOnServer from parameters. %s", err.Error()), client, callID, nil, nil, false)
		return
	}

	go protocol.retrieveMessagesHandler(nil, client, callID, recipient.(*messaging_types.MessageRecipient), lstMsgIDs, bLeaveOnServer)
}
