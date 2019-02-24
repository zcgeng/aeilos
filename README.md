# aeilos -- an online multi player mineswipper!

## Tasks

- Left+right click
- Persist data
- Load data
- User register/login
- Score board
- DIY flag
- Frontend beautify
- Refine updateZeros

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
$ go biuld
# start the service
$ ./aeilos
```