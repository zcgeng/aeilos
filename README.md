# aeilos -- an online multi player mineswipper!

## Tasks

- Feature
	- User register/login
	- Score board
	- DIY flag
- Front
	- Frontend beautify
	- MiniMap
- Back
	- Refine updateZeros
	- Record user current area
	- Stop global broadcasting




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