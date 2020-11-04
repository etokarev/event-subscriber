package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/Bestowinc/protoss/gen/go/proto/core"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	projectName string
	subName string
)


func init() {
	flag.StringVar(&projectName, "project", "", "")
	flag.StringVar(&subName, "sub", "", "")
}

func main() {
	flag.Parse()

	fmt.Printf("Connecting to [%v] pubsub project\n", projectName)

	client, err := pubsub.NewClient(context.Background(), projectName)
	if err != nil {
		panic(err)
	}

	subscription := client.Subscription(subName)
	exists, err := subscription.Exists(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		panic(fmt.Sprintf("subscription %v does not exist in %v project", subName, projectName))
	}

	unpacker, err := NewUnpacker(PolicyProtoEvents)
	if err != nil {
		panic(err)
	}

	err = subscription.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		envelope := &core.EventEnvelope{}
		if err := proto.Unmarshal(msg.Data, envelope); err != nil {
			panic(err)
		}
		fmt.Printf("[%v]: TypeUrl: %v\n", msg.ID, envelope.TypeUrl)

		event, err := unpacker.Unpack(envelope)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[%v]: Event: %v", msg.ID, prettyProtoMarshal(event))
		msg.Ack()
	})
	if err != nil {
		panic(err)
	}

}

// prettyProtoMarshal convert a proto message to JSON representation
func prettyProtoMarshal(pb proto.Message) string {
	m := &jsonpb.Marshaler{Indent: "\t"}
	var buf bytes.Buffer
	if err := m.Marshal(&buf, pb); err != nil {
		panic(err)
	}

	return string(buf.Bytes())
}
