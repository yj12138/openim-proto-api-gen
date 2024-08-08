package imapigen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func getCodeGeneratorRequestByFile(protoFile string) (*descriptorpb.FileDescriptorSet, error) {
	tmpFile, err := os.CreateTemp("", "descriptor*.pb")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	cmd := exec.Command("protoc",
		"--descriptor_set_out="+tmpFile.Name(),
		"--include_source_info",
		"--include_imports",
		"--proto_path=.",
		protoFile)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Failed to run protoc: %v", err)
	}
	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}
	fileSet := &descriptorpb.FileDescriptorSet{}
	if err := proto.Unmarshal(data, fileSet); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal descriptor set: %v", err)
	}
	return fileSet, nil
}

func TestGen(t *testing.T) {
	log.SetOutput(os.Stdout)
	fileDescSet, err := getCodeGeneratorRequestByFile("test.proto")
	if err != nil {
		t.Error(err)
	}
	for _, fileDesc := range fileDescSet.File {
		err = Gen(fileDesc)
		if err != nil {
			t.Error(err)
		}
	}
}
