package secure_connection

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendReport sets the SendReport handler function
func (protocol *SecureConnectionProtocol) SendReport(handler func(err error, client *nex.Client, callID uint32, reportID uint32, report []byte)) {
	protocol.SendReportHandler = handler
}

func (protocol *SecureConnectionProtocol) HandleSendReport(packet nex.PacketInterface) {
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

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 4 {
		err := errors.New("[SecureConnection::SendReport] Data missing report ID")
		go protocol.SendReportHandler(err, client, callID, 0, []byte{})
		return
	}

	reportID := parametersStream.ReadUInt32LE()
	report, err := parametersStream.ReadQBuffer()

	if err != nil {
		go protocol.SendReportHandler(err, client, callID, 0, []byte{})
		return
	}

	go protocol.SendReportHandler(nil, client, callID, reportID, report)
}
