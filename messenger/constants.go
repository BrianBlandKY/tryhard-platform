package messenger

// Messenger Status
// Not to be used in Commands
const (
	DISCONNECTED = Status(iota)
	CONNECTED
	CLOSED
	RECONNECTING
	CONNECTING
)

// PlatformServices Struct
type PlatformServices struct {
	Connection string
	Party      string
}

// Services - Defined on initiation.
var Services PlatformServices

// Party States
const (
	// OPENED State
	OPENED = "OPENED"

	// SUSPENDED State
	SUSPENDED = "SUSPENDED"

	// STALE State
	STALE = "STALE"

	// INGAME State
	INGAME = "IN-GAME"

	// JOINED State
	JOINED = "PLAYER_JOINED"

	// DISBANDED State
	DISBANDED = "PLAYER_DISBANDED"

	// NOTAVAILABLE State
	NOTAVAILABLE = "NOTAVAILABLE"
)

// Actions
const (
	// CONNECT State
	CONNECT = "CONNECT"

	// DISCONENECT State
	DISCONNECT = "DISCONNECT"

	// TIMEOUT State
	TIMEOUT = "TIMEOUT"

	// HEARTBEAT State
	HEARTBEAT = "HEARTBEAT"

	// JOIN State
	JOIN = "JOIN"

	// DISBAND State
	DISBAND = "DISBAND"

	// ERROR State
	ERROR = "ERROR"
)
