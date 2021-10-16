package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// RemoteLogDeviceProtocolID is the protocol ID for the RemoteLogDevice protocol
	RemoteLogDeviceProtocolID = 0x1

	// RemoteLogDeviceMethodLog is the method ID for the method Log
	RemoteLogDeviceMethodLog = 0x1
)

// RemoteLogDeviceProtocol handles the RemoteLogDevice protocol
type RemoteLogDeviceProtocol struct {
	server     *nex.Server
	LogHandler func(err error, client *nex.Client, callID uint32, strMessage string)
}

// Setup initializes the protocol
func (remoteLogDeviceProtocol *RemoteLogDeviceProtocol) Setup() {
	nexServer := remoteLogDeviceProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if RemoteLogDeviceProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case RemoteLogDeviceMethodLog:
				go remoteLogDeviceProtocol.handleLog(packet)
				break
			default:
				fmt.Printf("Unsupported RemoteLogDevice method ID: %#v\n", request.MethodID())
				break
			}
		}
	})
}

func (remoteLogDeviceProtocol *RemoteLogDeviceProtocol) handleLog(packet nex.PacketInterface) {
	if remoteLogDeviceProtocol.LogHandler == nil {
		fmt.Println("[Warning] RemoteLogDeviceProtocol::Log not implemented")
		go respondNotImplemented(packet, RemoteLogDeviceProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, remoteLogDeviceProtocol.server)

	var err error
	var strMessage string
	strMessage, err = parametersStream.ReadString()
	if err != nil {
		go remoteLogDeviceProtocol.LogHandler(err, client, callID, "")
	}

	go remoteLogDeviceProtocol.LogHandler(nil, client, callID, strMessage)
}

// Log sets the Log handler function
func (remoteLogDeviceProtocol *RemoteLogDeviceProtocol) Log(handler func(err error, client *nex.Client, callID uint32, strMessage string)) {
	remoteLogDeviceProtocol.LogHandler = handler
}

// NewRemoteLogDeviceProtocol returns a new RemoteLogDeviceProtocol
func NewRemoteLogDeviceProtocol(server *nex.Server) *RemoteLogDeviceProtocol {
	remoteLogDeviceProtocol := &RemoteLogDeviceProtocol{server: server}

	remoteLogDeviceProtocol.Setup()

	return remoteLogDeviceProtocol
}
