# aeilos -- an online multi player mineswipper!

## Tasks

- Feature
	- User register/login
	- Score of each user
	- Punishment of minus-zero
	- DIY flag
	- Own an area
	- Chat
- Front
	- MiniMap
	- Frontend beautify
- Back
	- Refine updateZeros() // because it might be slow
	- Record user current area and Stop global broadcasting




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