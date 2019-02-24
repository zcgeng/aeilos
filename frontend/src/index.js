import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
const pb = require('./aeilos_pb');

const ROW_LENGTH = 40;
const ROW_HEIGHT = 20;

function InsideArea(x, y, ax, ay) {
  return (x >= ax) && (y >= ay) && (x < ax+ROW_HEIGHT) && (y < ay+ROW_LENGTH);
}

function getCellDesc(pbcell) {
  switch(pbcell.getCelltypeCase()) {
    case pb.Cell.CelltypeCase.BOMBS:
      if(pbcell.getBombs() === 0) return '0';
      if(pbcell.getBombs() === 9) return '💣';
      if(pbcell.getBombs() === 11) return '??';
      return pbcell.getBombs();
    case pb.Cell.CelltypeCase.UNTOUCHED:
      return ' '
    case pb.Cell.CelltypeCase.FLAGURL:
      return '🚩'
    default:
      alert('error: cell no type')
      return ' '
  }
}

function cellIsBomb(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.BOMBS) && (pbcell.getBombs() === 9);
}

function cellIsNumber(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.BOMBS) && (pbcell.getBombs() !== 9);
}

function cellIsFlag(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.FLAGURL);
}

function cellIsBombOrFlag(pbcell) {
  return cellIsBomb(pbcell) || cellIsFlag(pbcell);
}

function getCellNumber(pbcell) {
  if(!cellIsNumber(pbcell)) {
    return 0;
  }
  return(pbcell.getBombs());
}

class Aeilos extends React.Component {
  constructor(props) {
    super(props);
    const socket = new WebSocket('wss://changgeng.me/ws/');
    // const socket = new WebSocket('ws://localhost:8000/ws/');
    this.state = {
      socket: socket,
      x: 0,
      y: 0,
    };
  }

  renderArea(x, y) {
    return (<Area x={x} y={y} socket={this.state.socket}/>)
  }

  handleWheel(e) {
    let xMove = 0;
    let yMove = 0;
    if(e.deltaY < 10 && e.deltaY > 0) {
      xMove = 1;
    }
    else if(e.deltaY < 0 && e.deltaY > -10) {
      xMove = -1;
    }
    else {
      xMove = Math.round(e.deltaY/100)
    }

    if(e.deltaX < 50 && e.deltaX > 0) {
      yMove = 1;
    }
    else if(e.deltaX < 0 && e.deltaX > -50) {
      yMove = -1;
    }
    else {
      yMove = Math.round(e.deltaX/80)
    }

    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(this.state.x + xMove);
    xy.setY(this.state.y + yMove);
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      socket: this.state.socket,
      x: this.state.x + xMove,
      y: this.state.y + yMove,
    });
  }

  handleMoveMap(direction) {
    // up down left right none
    let xmoves = [-3, 3, 0, 0]
    let ymoves = [0, 0, -3, 3]
    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(this.state.x + xmoves[direction]);
    xy.setY(this.state.y + ymoves[direction]);
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      socket: this.state.socket,
      x: this.state.x + xmoves[direction],
      y: this.state.y + ymoves[direction],
    });
  }

  handleCoordX(event) {
    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(parseInt(event.target.value, 10));
    xy.setY(this.state.y);
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      socket: this.state.socket,
      x: parseInt(event.target.value, 10),
      y: this.state.y,
    });
  }

  handleCoordY(event) {
    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(this.state.x);
    xy.setY(parseInt(event.target.value, 10));
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      socket: this.state.socket,
      x: this.state.x,
      y: parseInt(event.target.value, 10),
    });
  }

  render() {
    return (
      <div onWheel={(e)=>{this.handleWheel(e)}}>
      <div>
        {this.renderArea(this.state.x, this.state.y)}
      </div>

      <div>
        <button onClick={()=>{
          this.handleMoveMap(0);
        }}>
          Go UP
        </button>
        <button onClick={()=>{
          this.handleMoveMap(1);
        }}>
          Go DOWN
        </button>
        <button onClick={()=>{
          this.handleMoveMap(2);
        }}>
          Go LEFT
        </button>
        <button onClick={()=>{
          this.handleMoveMap(3);
        }}>
          Go RIGHT
        </button>
      </div>

      <div>
        Current location: ({this.state.x}, {this.state.y})
      </div>

      <div>
        Jump to coordinate:
        <input type="number" value={this.state.x} onChange={this.handleCoordX.bind(this)} />
        <input type="number" value={this.state.y} onChange={this.handleCoordY.bind(this)} />
      </div>
      <div>You can also use [scroll wheel] or [shift]+[scroll wheel] to move</div>
      </div>
    );
  }
}

class Area extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      baseXY: {x: props.x, y:props.y},
      mouseDown: {l:false, r:false},
      curArea: [],
    };
    const socket = props.socket;

    socket.addEventListener('open', (event)=>{
      let msg = new pb.ClientToServer();
      let xy = new pb.XY();
      xy.setX(props.x);
      xy.setY(props.y);
      msg.setGetarea(xy);
      socket.send(msg.serializeBinary());
    });

    var that = this;
    socket.addEventListener('message', function (event) {
      var blob = event.data;
      var fileReader     = new FileReader();
      fileReader.onload = function(event) {
          let response = pb.ServerToClient.deserializeBinary(event.target.result);
          switch(response.getResponseCase()){

            case pb.ServerToClient.ResponseCase.TOUCH:
              console.log("you got",response.getTouch().getScore()+" scores")
              let cell = response.getTouch().getCell();
              if(!InsideArea(cell.getX(), cell.getY(), that.state.baseXY.x, that.state.baseXY.y)){
                break;
              }
              let newArea = that.state.curArea.map((arr)=>{return arr.slice();});
              let {x, y} = that.glob2local(cell.getX(), cell.getY())
              newArea[x][y] = cell;
              that.setState({
                curArea: newArea,
                mouseDown: that.state.mouseDown,
                baseXY: that.state.baseXY,
              })
              break;

            case pb.ServerToClient.ResponseCase.AREA:
              let cellsList = response.getArea().getCellsList();
              // reshape the cellsList[1500] to [ROW_LENGTH][30]
              let cells2d = [];
              while(cellsList.length) 
                cells2d.push(cellsList.splice(0,ROW_LENGTH))

              that.setState({
                curArea: cells2d,
                mouseDown: that.state.mouseDown,
                baseXY: {x: response.getArea().getX(), y: response.getArea().getY()},
              });
              break;

            case pb.ServerToClient.ResponseCase.MSG:
              console.log(response.getMsg());
              break;

            case pb.ServerToClient.ResponseCase.UPDATE:
              let cell1 = response.getUpdate();
              if(!InsideArea(cell1.getX(), cell1.getY(), that.state.baseXY.x, that.state.baseXY.y)){
                break;
              }
              
              // let newArea1 = that.state.curArea.map((arr)=>{return arr.slice();});
              let xy1 = that.glob2local(cell1.getX(), cell1.getY());
              that.state.curArea[xy1.x][xy1.y] = cell1;

              // var t0 = performance.now();
              /* TOOOOOO SLOW!!!! */
              // that.setState({
              //   curArea: newArea1,
              //   baseXY: that.state.baseXY,
              //   socket: that.state.socket,
              // })
              // var t1 = performance.now();
              // console.log("update takes: ", t1-t0);
              break;
            default:
              alert('error: response no type')
          } 
      };
      fileReader.readAsArrayBuffer(blob);
    });

  }

  glob2local(x, y) {
    return {x: x - this.state.baseXY.x, y: y - this.state.baseXY.y};
  }

  handleMouseDown(e) {
    if(e.nativeEvent.which === 1){
      this.state.mouseDown.l = true
    }
    if(e.nativeEvent.which === 3){
      this.state.mouseDown.r = true
    }
  }

  // get the bomb number shown on the map
  getNeighbourBombs(locX, locY) {
    let bombCount = 0;
    for(let i = -1; i < 2; i++){
      for(let j = -1; j < 2; j++){
        if(cellIsBombOrFlag(this.state.curArea[locX+i][locY+j])){
          bombCount++;
        }
      }
    }
    return bombCount
  }

  handleSimulClick(globX, globY) {
    let {x, y} = this.glob2local(globX, globY)
    if(x === 0 || x === ROW_HEIGHT-1 || y === 0 || y === ROW_LENGTH-1){
      return // we don't handle corner cases
    }
    let cell = this.state.curArea[x][y]
    if(cell.getCelltypeCase() !== pb.Cell.CelltypeCase.BOMBS 
      || cell.getBombs() === 9|| cell.getBombs() === 0){
      return
    }
    // console.log("neighbour:" ,this.getNeighbourBombs(x, y))
    // console.log("number:" ,getCellNumber(cell));
    if(this.getNeighbourBombs(x, y) === getCellNumber(cell)){
      this.flipNeighbours(globX, globY);
    }
  }

  handleMouseUp(globX, globY, e) {
    if(this.state.mouseDown.l && this.state.mouseDown.r){
      this.handleSimulClick(globX, globY);
    }
    if(e.nativeEvent.which === 1){
      this.state.mouseDown.l = false;
    }
    if(e.nativeEvent.which === 3){
      this.state.mouseDown.r = false;
    }
  }

  flipNeighbours(globX, globY) {
    for(let i = -1; i < 2; i++){
      for(let j = -1; j < 2; j++){
        this.flipCell(globX+i, globY+j);
      }
    }
  }

  flipCell(globX, globY) {
    // global x and global y
    let {x, y} = this.glob2local(globX, globY)
    if(getCellDesc(this.state.curArea[x][y]) !== ' '){
      return
    }
    let msg = new pb.ClientToServer();
    let touch = new pb.TouchRequest();
    touch.setX(globX);
    touch.setY(globY);
    touch.setTouchtype(pb.TouchType.FLIP);
    msg.setTouch(touch);
    this.props.socket.send(msg.serializeBinary());
  }

  flagCell(globX, globY) {
    // global x and global y
    let {x, y} = this.glob2local(globX, globY)
    if(getCellDesc(this.state.curArea[x][y]) !== ' '){
      return
    }
    let msg = new pb.ClientToServer();
    let touch = new pb.TouchRequest();
    touch.setX(globX);
    touch.setY(globY);
    touch.setTouchtype(pb.TouchType.FLAG);
    msg.setTouch(touch);
    this.props.socket.send(msg.serializeBinary());
  }

  handleClick(globX, globY, e) {
    // global x and global y
    let {x, y} = this.glob2local(globX, globY)
    if(getCellDesc(this.state.curArea[x][y]) !== ' '){
      return
    }
    if(e.type === 'click'){
      this.flipCell(globX, globY);
    } else if(e.type === 'contextmenu') {
      e.preventDefault();
      this.flagCell(globX, globY);
    }
  }

  render() {
    const mmap = this.state.curArea.map((row)=>{
      return row.map((cell)=>{
        return getCellDesc(cell);
      });
    });

    const cellBoard = mmap.map((row, i) => {
      const cellRow = row.map((val, j) => {
        return <Cell 
          key={j} 
          value={val} 
          x={this.state.baseXY.x+i} 
          y={this.state.baseXY.y+j}
          onClick={
            (event)=>{
              event.preventDefault();
              this.handleClick(this.state.baseXY.x+i, this.state.baseXY.y+j, event);
            }
          }
          onContextMenu={
            (event)=>{
              event.preventDefault();
              this.handleClick(this.state.baseXY.x+i, this.state.baseXY.y+j, event);
            }
          }
          onMouseDown={this.handleMouseDown.bind(this)}
          onMouseUp={
            (event)=>{
              this.handleMouseUp(this.state.baseXY.x+i, this.state.baseXY.y+j, event);
            }
          }
        />
      });
      return (<div key={i} className="board-row">{cellRow}</div>);
    });
    return (
      <div>
        {cellBoard}
      </div>
    );
  }
}

function Cell(props) {
  return (
    <button
      className="square"
      onClick={props.onClick}
      onContextMenu={props.onContextMenu}
      onMouseDown={props.onMouseDown}
      onMouseUp={props.onMouseUp}
      onDoubleClick={(e)=>{}}
      unselectable="on"
      // onSelectStart="return false;" 
    >
      {props.value}
    </button>
  );
}

// ========================================

ReactDOM.render(
  // <Game />,
  <Aeilos />,
  document.getElementById('root')
);
