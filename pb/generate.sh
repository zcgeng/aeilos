protoc aeilos.proto --go_out=./;
protoc aeilos.proto --js_out=import_style=commonjs,binary:../frontend/src;
{ echo "/* eslint-disable */"; cat ../frontend/src/aeilos_pb.js; } > ../frontend/src/aeilos_pb.js.new;
mv ../frontend/src/aeilos_pb.js.new ../frontend/src/aeilos_pb.js;