package main

import (
	"io"
	"os"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/go-multierror"

	"code.vegaprotocol.io/vega/cmd/scenariorunner/core"
)

const indent string = "  "

// ProcessFiles takes an array of paths to files, reads them in and returns their contents as an array of instruction sets (set per file)
func ProcessFiles(filesWithPath []string) ([]*core.InstructionSet, error) {
	contents, err := openFiles(filesWithPath)
	if err != nil {
		return nil, err
	}

	var errs *multierror.Error
	instructionSets := make([]*core.InstructionSet, len(contents))

	for i, fileContents := range contents {
		instrSet := &core.InstructionSet{}
		marshallErr := unmarshal(fileContents, instrSet)
		if marshallErr != nil {
			errs = multierror.Append(errs, marshallErr)
		}
		instructionSets[i] = instrSet
	}

	return instructionSets, errs.ErrorOrNil()
}

// Output writes results to the specified file.
func Output(result proto.Message, outputFileWithPath string) error {
	f, err := os.Create(outputFileWithPath)
	if err != nil {
		return err
	}
	return marshal(result, f)
}

func openFiles(filesWithPath []string) ([]*os.File, error) {
	var n = len(filesWithPath)
	readers := make([]*os.File, n)
	var errs *multierror.Error
	var err error

	for i := 0; i < n; i++ {
		readers[i], err = os.Open(filesWithPath[i])
		if err != nil {
			errs = multierror.Append(errs, err)
		}
	}
	return readers, errs.ErrorOrNil()
}

func unmarshal(r io.Reader, msg proto.Message) error {
	return jsonpb.Unmarshal(r, msg)
}

func marshal(result proto.Message, out io.Writer) error {
	m := jsonpb.Marshaler{Indent: indent, EmitDefaults: true}
	return m.Marshal(out, result)
}
