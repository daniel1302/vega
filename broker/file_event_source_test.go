package broker_test

import (
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"code.vegaprotocol.io/data-node/broker"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/types"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestReceiveEvents(t *testing.T) {
	f, err := os.CreateTemp("", "vega.evt")
	filePath, err := filepath.Abs(filepath.Dir(f.Name()))
	defer os.Remove(filePath)
	defer f.Close()
	if err != nil {
		t.Errorf("failed to create temporary file for events:%s", err)
	}

	a1 := events.NewAssetEvent(context.Background(), types.Asset{ID: "1"})
	a2 := events.NewAssetEvent(context.Background(), types.Asset{ID: "2"})
	a3 := events.NewAssetEvent(context.Background(), types.Asset{ID: "3"})

	evts := []*eventspb.BusEvent{
		a1.StreamMessage(), a2.StreamMessage(),
		a3.StreamMessage(),
	}

	file := &testEventFile{}

	writeEventsToFile(evts, file)
	f.Close()

	source, err := broker.NewFileEventSource(file, 0, 0)
	if err != nil {
		t.Errorf("failed to create file event source:%s", err)
	}

	evtCh, _ := source.Receive(context.Background())

	e1 := <-evtCh
	r1 := e1.(*events.Asset)
	e2 := <-evtCh
	r2 := e2.(*events.Asset)
	e3 := <-evtCh
	r3 := e3.(*events.Asset)

	assert.Equal(t, "1", r1.Asset().Id)
	assert.Equal(t, "2", r2.Asset().Id)
	assert.Equal(t, "3", r3.Asset().Id)
}

func writeEventsToFile(events []*eventspb.BusEvent, fi *testEventFile) {
	sizeBytes := make([]byte, 4)

	for _, e := range events {
		size := uint32(proto.Size(e))
		protoBytes, err := proto.Marshal(e)
		if err != nil {
			panic("failed to marshal bus event:" + e.String())
		}

		binary.BigEndian.PutUint32(sizeBytes, size)
		allBytes := append(sizeBytes, protoBytes...)
		fi.Write(allBytes)
	}
}

type testEventFile struct {
	bytes []byte
}

func (t *testEventFile) Open() error {
	return nil
}

func (t *testEventFile) Close() error {
	return nil
}

func (t *testEventFile) Write(b []byte) (n int, err error) {
	t.bytes = append(t.bytes, b...)
	return len(b), nil
}

func (t *testEventFile) ReadAt(b []byte, off int64) (n int, err error) {
	r := bytes.NewReader(t.bytes)
	return r.ReadAt(b, off)
}
