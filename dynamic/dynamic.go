package dynamic

import (
	"embed"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

/*func ParseAs(fqdn string, message_bytes []byte) (*protoreflect.ProtoMessage, error) {
	messageName := protoreflect.FullName(fqdn)
	pbDesc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(messageName))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	msg := pbDesc.(protoreflect.MessageDescriptor)
	dynamicpb.NewMessage()
	err = proto.Unmarshal(message_bytes, msg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return (&msg), nil
}*/

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

func recoverableRegistration(f protoreflect.FileDescriptor) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	protoregistry.GlobalFiles.RegisterFile(f)
	return
}

func Register(pbfile string, efs *embed.FS) {
	fileDescriptorSet := pbfile

	var data []byte
	var err error
	// Read descriptors from file
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

	// Process descriptors from Protobuf into their runtime representation
	//var registry protoregistry.Files

	for _, file := range files.File {
		_, err := protoregistry.GlobalFiles.FindFileByPath(file.GetName())
		if err == nil {
			fmt.Printf("Not registering: %q\n", file.GetName())
			continue
		}
		fileDescriptor, err := protodesc.NewFile(file, protoregistry.GlobalFiles) //&registry)
		if err != nil {
			log.Fatalf("failed to process %q: %v\n", file.GetName(), err)
		}
		fmt.Printf("Processed %q as descriptor\n", file.GetName())

		//err = recoverableRegistration(fileDescriptor)
		if err := protoregistry.GlobalFiles.RegisterFile(fileDescriptor); err != nil {
			//fmt.Println("meh")
			log.Fatalf("failed to process %q: %v\n", file.GetName(), err)
		}
		fmt.Printf("Registered %q\n", file.GetName())
	}

	// Get descriptor for message type
	/* descriptor, err := protoregistry.GlobalFiles.FindDescriptorByName(messageType)
	   if err != nil {
	           log.Fatalf("failed to find message type %q in given descriptors: %v\n", messageType, err)
	   }
	   messageDescriptor, ok := descriptor.(protoreflect.MessageDescriptor)
	   if !ok {
	           log.Fatalf("element named %q is not a message (%T)\n", messageType, descriptor)
	   }

	   // Now we can create a dynamic message and use that to read the binary format from stdin
	   messageData, err := io.ReadAll(os.Stdin)
	   if err != nil {
	           log.Fatalf("failed to read message data from stdin: %v\n", err)
	   }
	   message := dynamicpb.NewMessage(messageDescriptor)
	   if err := proto.Unmarshal(messageData, message); err != nil {
	           log.Fatalf("failed to process input data for message type %q: %v\n", messageType, err)
	   }

	   // And write text format to stdout
	   _, _ = fmt.Print(prototext.Format(message))
	*/
}
