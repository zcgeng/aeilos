# aeilos -- an online multi player minesweeper!

## Tasks

- Bugs
  - Cannot Save Login
- Feature
  - Punishment of minus-zero
  - DIY flag
  - Own an area
  - MiniMap
- Frontend
  - Frontend redesign
    - LeaderBoard css style
    - hide chatbox by animation
  - area size
  - Adjust onWheel
  - Cache areas
- Backend
  - Refine updateZeros() // because it might be slow
  - Record user current area and Stop global broadcasting
- Done
  - Left+right click
  - Persist & load data
  - User register/login
  - Score board
  - Chatbox: send and receive, get history messages
  - Ranking: get my ranking, leaderboard

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
