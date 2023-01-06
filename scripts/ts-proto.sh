ls ./proto/checkers | xargs -I {} ./node_modules/protoc/protoc/bin/protoc \
  --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
  --ts_proto_out="./client/src/types/generated" \
  --proto_path="./proto" \
  --ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
  checkers/{}
