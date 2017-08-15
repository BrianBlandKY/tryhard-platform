package messenger

// Constants
const (
	DISCONNECTED = Status(iota)
	CONNECTED
	CLOSED
	RECONNECTING
	CONNECTING
	HEARTBEAT
)

// Deprecate constants below

// Services
const (
	// PARTY
	PARTY = "PARTY"
)

// Commands
const (
	// JOIN
	JOIN = "JOIN"

	// DISBAND
	DISBAND = "DISBAND"
)

// States
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
	// ERROR STATE
	ERROR = "ERROR"

	// CONNECT State
	CONNECT = "CONNECT"

	// DISCONENECT State
	DISCONNECT = "DISCONNECT"
)
