package dynamic

import (
	"embed"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func ParseAs(messageType string, message_bytes []byte) (*protoreflect.ProtoMessage, error) {

	descriptor, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(messageType))
	if err != nil {
		log.Fatalf("failed to find message type %q in given descriptors: %v\n", messageType, err)
	}
	messageDescriptor, ok := descriptor.(protoreflect.MessageDescriptor)
	if !ok {
		log.Fatalf("element named %q is not a message (%T)\n", messageType, descriptor)
	}

	message := dynamicpb.NewMessage(messageDescriptor)
	if err := proto.Unmarshal(message_bytes, message); err != nil {
		log.Fatalf("failed to process input data for message type %q: %v\n", messageType, err)
	}
	msg := message.Interface()
	return &msg, nil
}

func Register(pbfile string, efs *embed.FS) {
	fileDescriptorSet := pbfile

	var data []byte
	var err error
	var files descriptorpb.FileDescriptorSet
	if efs != nil {
		data, err = efs.ReadFile(fileDescriptorSet)
	} else {
		data, err = os.ReadFile(fileDescriptorSet)
	}
	if err != nil {
		log.Fatalln(err)
	}
	if err := proto.Unmarshal(data, &files); err != nil {
		log.Fatalf("failed to process descriptors in %s: %v\n", fileDescriptorSet, err)
	}

	log.Printf("Registered at runtimes:  ")
	for _, file := range files.File {
		_, err := protoregistry.GlobalFiles.FindFileByPath(file.GetName())
		if err == nil {
			continue
		}
		fileDescriptor, err := protodesc.NewFile(file, protoregistry.GlobalFiles)
		if err != nil {
			log.Fatalf("failed to process %q: %v\n", file.GetName(), err)
		}
		if err := protoregistry.GlobalFiles.RegisterFile(fileDescriptor); err != nil {
			log.Fatalf("failed to process %q: %v\n", file.GetName(), err)
		}
		log.Printf("    %q", file.GetName())
		//fmt.Printf("%q, ", file.GetName())
	}
	//fmt.Printf("]\n")
}
