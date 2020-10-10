package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// NotificationsProtocolID is the protocol ID for the Notifications protocol
	NotificationsProtocolID = 0x0E

	// NotificationsMethodProcessNotificationEvent is the method ID for the method ProcessNotificationEvent
	NotificationsMethodProcessNotificationEvent = 0x1
)

// NotificationsProtocol handles the Notifications nex protocol
type NotificationsProtocol struct {
	server                          *nex.Server
	ProcessNotificationEventHandler func(err error, client *nex.Client, callID uint32, oEvent *NotificationEvent)
}

// NotificationEvent holds information about a notification event
type NotificationEvent struct {
	pidSource uint32
	uiType    uint32
	uiParam1  uint32
	uiParam2  uint32
	strParam  string

	hierarchy []nex.StructureInterface
	*nex.NullData
}

// GetHierarchy returns the Structure hierarchy
func (notificationEvent *NotificationEvent) GetHierarchy() []nex.StructureInterface {
	return notificationEvent.hierarchy
}

// ExtractFromStream extracts a NotificationEvent structure from a stream
func (notificationEvent *NotificationEvent) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	if len(stream.Bytes()[stream.ByteOffset():]) < 18 {
		return errors.New("[NotificationEvent::ExtractFromStream] Data size too small")
	}

	notificationEvent.pidSource = stream.ReadUInt32LE()
	notificationEvent.uiType = stream.ReadUInt32LE()
	notificationEvent.uiParam1 = stream.ReadUInt32LE()
	notificationEvent.uiParam2 = stream.ReadUInt32LE()
	var param string
	param, err = stream.ReadString()
	if err != nil {
		return err
	}
	notificationEvent.strParam = param

	fmt.Printf("%+v\n", notificationEvent)

	return nil
}

// NewNotificationEvent returns a new NotificationEvent
func NewNotificationEvent() *NotificationEvent {
	notificationEvent := &NotificationEvent{}

	nullData := nex.NewNullData()

	notificationEvent.NullData = nullData

	notificationEvent.hierarchy = []nex.StructureInterface{
		nullData,
	}

	return notificationEvent
}

// Setup intitalizes the protocol
func (notificationsProtocol *NotificationsProtocol) Setup() {
	nexServer := notificationsProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if NotificationsProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case NotificationsMethodProcessNotificationEvent:
				go notificationsProtocol.handleProcessNotificationEvent(packet)
				break
			default:
				fmt.Printf("Unsupported Notifications method ID: %#v\n", request.MethodID())
				go notificationsProtocol.respondNotImplemented(packet)
				break
			}
		}
	})
}

func (notificationsProtocol *NotificationsProtocol) respondNotImplemented(packet nex.PacketInterface) {
	client := packet.Sender()
	request := packet.RMCRequest()

	rmcResponse := nex.NewRMCResponse(NotificationsProtocolID, request.CallID())
	rmcResponse.SetError(0x80010002)

	rmcResponseBytes := rmcResponse.Bytes()

	var responsePacket nex.PacketInterface
	if packet.Version() == 1 {
		responsePacket, _ = nex.NewPacketV1(client, nil)
	} else {
		responsePacket, _ = nex.NewPacketV0(client, nil)
	}

	responsePacket.SetVersion(packet.Version())
	responsePacket.SetSource(packet.Destination())
	responsePacket.SetDestination(packet.Source())
	responsePacket.SetType(nex.DataPacket)
	responsePacket.SetPayload(rmcResponseBytes)

	responsePacket.AddFlag(nex.FlagNeedsAck)
	responsePacket.AddFlag(nex.FlagReliable)
	notificationsProtocol.server.Send(responsePacket)
}

func (notificationsProtocol *NotificationsProtocol) handleProcessNotificationEvent(packet nex.PacketInterface) {
	if notificationsProtocol.ProcessNotificationEventHandler == nil {
		fmt.Println("[Warning] NotificationsProtocol::ProcessNotificationEvent not implemented")
		go notificationsProtocol.respondNotImplemented(packet)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, notificationsProtocol.server)

	dataHolderName, err := parametersStream.ReadString()

	if err != nil {
		go notificationsProtocol.ProcessNotificationEventHandler(err, client, callID, nil)
		return
	}

	if dataHolderName != "NotificationEvent" {
		err := errors.New("[NotificationProtocol::ProcessNotificationEvent] Data holder name does not match")
		go notificationsProtocol.ProcessNotificationEventHandler(err, client, callID, nil)
		return
	}

	_ = parametersStream.ReadUInt32LE() // length including this field

	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go notificationsProtocol.ProcessNotificationEventHandler(err, client, callID, nil)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, notificationsProtocol.server)

	notificationEvent, err := dataHolderContentStream.ReadStructure(NewAuthenticationInfo())

	if err != nil {
		go notificationsProtocol.ProcessNotificationEventHandler(err, client, callID, nil)
		return
	}

	go notificationsProtocol.ProcessNotificationEventHandler(nil, client, callID, notificationEvent.(*NotificationEvent)) // TODO: Finish this call and uncomment using parsed values!
}

// ProcessNotificationEvent sets the ProcessNotificationEvent handler function
func (notificationsProtocol *NotificationsProtocol) ProcessNotificationEvent(handler func(err error, client *nex.Client, callID uint32, oEvent *NotificationEvent)) {
	notificationsProtocol.ProcessNotificationEventHandler = handler
}

//NewNotificationsProtocol returns a new NotificationsProtocol
func NewNotificationsProtocol(server *nex.Server) *NotificationsProtocol {
	notificationsProtocol := &NotificationsProtocol{server: server}

	notificationsProtocol.Setup()

	return notificationsProtocol
}
