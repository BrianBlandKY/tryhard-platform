package realm

import (
	"log"
	d "try-hard-platform/messenger"
)

type Realm struct {
	conn d.Connection
}

// Connect to server
func (r *Realm) Connect(url, id string) {
	conn, err := d.Connect(url, id)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}
	r.conn = conn

	r.connectionSubscription()
	r.connectedSubscription()
	r.heartbeatSubscription()
}

// Close server
func (r *Realm) Close() {
	r.conn.Close()
}

// subscribe to client connection messages
// - reply to client
// - publish/broadcast to other clients
func (r *Realm) connectionSubscription() {
	_, err := r.conn.Subscribe("connection", func(subject, reply string, cmd d.Command) {
		log.Println("SERVER - connection", cmd)

		// send reply to client
		r.conn.PublishRequestCommand(reply, d.Command{
			Data:    d.CONNECTED,
			Subject: subject,
		})

		// publish connection
		cmd.Subject = "connected"
		r.conn.PublishCommand(cmd)
	})
	if err != nil {
		log.Println("Failed to subscribe to [connection].")
		log.Panicln(err)
	}
}

func (r *Realm) connectedSubscription() {
	_, err := r.conn.Subscribe("connected", func(subject, reply string, cmd d.Command) {
		// log connected client
		log.Println("SERVER - connected", cmd)
	})
	if err != nil {
		log.Println("Failed to subscribe to [connected].")
		log.Panicln(err)
	}
}

// subscribe to heartbeat messages
// it is the responsibility of the client to manage the frequency.
// - reply to client
// - capture times difference?
// - auto-disconnect clients after x time? (this is the only way to capture disconnections)
func (r *Realm) heartbeatSubscription() {
	_, err := r.conn.Subscribe("heartbeat", func(subject, reply string, cmd d.Command) {
		log.Println("SERVER - heartbeat", cmd)

		// send reply to client instant
		cmd.Data = d.HEARTBEAT
		r.conn.PublishRequestCommand(reply, d.Command{
			Data:    d.HEARTBEAT,
			Subject: subject,
		})
	})
	if err != nil {
		log.Println("Failed to subscribe to [heartbeat].")
		log.Panicln(err)
	}
}

func main() {
	cm := Realm{}
	cm.Connect("", "")
}
