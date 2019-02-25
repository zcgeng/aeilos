# aeilos -- an online multi player mineswipper!

## Tasks

- Front
	- MiniMap
	- DIY flag
	- Frontend beautify
- Back
	- Refine updateZeros
- All
	- User register/login
	- Score board




## How to Build

### Build protobuf
First install protobuf compiler from google. Then generate the grpc source code.
```sh
$ cd pb; ./generate.sh;
```

### Build frontend
```sh
$ cd frontend; npm run build;
```

### Build backend
```sh
# under root directory of this project
$ go build
# start the service
$ ./aeilos
```