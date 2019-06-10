# aeilos -- an online multi player minesweeper!

## Tasks

- Bugs
  - Cannot Save Login
  - LeaderBoard css style
- Feature
  - Punishment of minus-zero
  - DIY flag
  - Own an area
- Front
  - area size
  - MiniMap
  - Adjust onWheel
  - Frontend beautify
  - Cache areas
- Back
  - Refine updateZeros() // because it might be slow
  - Record user current area and Stop global broadcasting

## How to Build

### Build protobuf

First install protobuf compiler from google. Then generate the grpc source code.

```sh
cd pb; ./generate.sh;
```

### Build frontend

```sh
cd frontend; ./deploy.sh;
```

### Build backend

```sh
# under root directory of this project
go build
# start the service
./aeilos
```
