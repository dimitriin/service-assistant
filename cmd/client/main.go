package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/dimitriin/service-assistant/pkg/protocol/models"
	"github.com/golang/protobuf/proto"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:1053")

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, 1)

	m := &models.CounterIncCmd{
		Name: "some_counter",
		Labels: map[string]string{
			"some_counter_label": "some_counter_label_value",
		},
	}

	payload, _ := proto.Marshal(m)

	fmt.Fprintf(conn, "%s%s", string(buf), string(payload))
	conn.Close()
}
