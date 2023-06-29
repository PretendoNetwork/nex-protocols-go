package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendReport sets the SendReport handler function
func (protocol *SecureConnectionProtocol) SendReport(handler func(err error, client *nex.Client, callID uint32, reportID uint32, reportData []byte)) {
	protocol.SendReportHandler = handler
}

func (protocol *SecureConnectionProtocol) handleSendReport(packet nex.PacketInterface) {
	if protocol.SendReportHandler == nil {
		globals.Logger.Warning("SecureConnection::SendReport not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	reportID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.SendReportHandler(fmt.Errorf("Failed to read reportID from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	reportData, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.SendReportHandler(fmt.Errorf("Failed to read reportData from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.SendReportHandler(nil, client, callID, reportID, reportData)
}
