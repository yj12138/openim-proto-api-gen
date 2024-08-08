package imapigen

import (
	"log"

	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

func Gen(file *descriptorpb.FileDescriptorProto) error {

	for i, message := range file.MessageType {
		log.Println("message:", *message.Name)
		for _, location := range file.GetSourceCodeInfo().GetLocation() {
			path := location.GetPath()
			// 检查是否是消息的注释
			if len(path) >= 2 && path[0] == 4 && path[1] == int32(i) {
				// 获取前缀注释
				leadingComments := location.GetLeadingComments()
				if strings.TrimSpace(leadingComments) != "" {
					log.Printf("  Leading Comments: %s \n", leadingComments)
				}

				// 获取后缀注释
				trailingComments := location.GetTrailingComments()
				if strings.TrimSpace(trailingComments) != "" {
					log.Printf("  Trailing Comments: %s \n", trailingComments)
				}
			}
		}
	}
	return nil
}
