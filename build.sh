protoc ./protobufStruct.proto --go_out ./
protoc ./message/message.proto --go_out ./
protoc ./message/messageHead.proto --go_out ./
protoc ./message/messageBody.proto --go_out ./
cp ./github.com/littlefish12345/go-qq-protobuf/* .
rm -r github.com