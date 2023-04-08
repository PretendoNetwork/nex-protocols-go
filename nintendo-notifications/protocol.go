package nintendo_notifications

import nex "github.com/PretendoNetwork/nex-go"

const (
	// ProtocolID is the protocol ID for the Nintendo Notifications protocol
	ProtocolID = 0x64

	// MethodProcessNintendoNotificationEvent1 is the method ID for the method ProcessNintendoNotificationEvent (1)
	MethodProcessNintendoNotificationEvent1 = 0x1

	// MethodProcessNintendoNotificationEvent2 is the method ID for the method ProcessNintendoNotificationEvent (2)
	MethodProcessNintendoNotificationEvent2 = 0x2
)

// NintendoNotificationsProtocol handles the NintendoNotifications protocol
type NintendoNotificationsProtocol struct {
	Server *nex.Server
}

// Setup initializes the protocol
func (protocol *NintendoNotificationsProtocol) Setup() {
	// TODO: Do something
	// This protocol doesn't seem to get requests from the client, it only sends them
	// So no handling is done for in-coming requests at the moment
}

// NewNintendoNotificationsProtocol returns a new NintendoNotificationsProtocol
func NewNintendoNotificationsProtocol(server *nex.Server) *NintendoNotificationsProtocol {
	nintendoNotificationsProtocol := &NintendoNotificationsProtocol{Server: server}

	nintendoNotificationsProtocol.Setup()

	return nintendoNotificationsProtocol
}
