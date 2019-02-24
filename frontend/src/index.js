import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
const pb = require('./aeilos_pb');

function InsideArea(x, y, ax, ay) {
  return (x >= ax) && (y >= ay) && (x < ax+10) && (y < ay+10);
}

function getCellDesc(pbcell) {
  switch(pbcell.getCelltypeCase()) {
    case pb.Cell.CelltypeCase.BOMBS:
      if(pbcell.getBombs() === 0) return '0';
      if(pbcell.getBombs() === 9) return '*';
      if(pbcell.getBombs() === 11) return '??';
      return pbcell.getBombs();
    case pb.Cell.CelltypeCase.UNTOUCHED:
      return ' '
    case pb.Cell.CelltypeCase.FLAGURL:
      return 'P'
    default:
      alert('error: cell no type')
      return ' '
  }
}

class Aeilos extends React.Component {
  constructor(props) {
    super(props);
    const socket = new WebSocket('ws://localhost:8000/ws');
    socket.addEventListener('open', (event)=>{
      let msg = new pb.ClientToServer();
      let xy = new pb.XY();
      xy.setX(0);
      xy.setY(0);
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
              let newArea = that.state.curArea.map((arr)=>{return arr.slice();});
              newArea[cell.getX()][cell.getY()] = cell;
              that.setState({
                curArea: newArea,
                baseXY: that.state.baseXY,
                socket: that.state.socket,
              })
              break;
            case pb.ServerToClient.ResponseCase.AREA:
              let cellsList = response.getArea().getCellsList();
              // reshape the cellsList[100] to [10][10]
              let cells2d = [];
              while(cellsList.length) 
                cells2d.push(cellsList.splice(0,10))

              that.setState({
                curArea: cells2d,
                baseXY: {x: response.getArea().getX(), y: response.getArea().getY()},
                socket: that.state.socket,
              });
              break;
            case pb.ServerToClient.ResponseCase.MSG:
              console.log(response.getMsg());
              break;
            case pb.ServerToClient.ResponseCase.UPDATE:
              let cell1 = response.getUpdate();
              console.log("auto explore zeros, ",cell1.getX(), cell1.getY())
              if(!InsideArea(cell1.getX(), cell1.getY(), that.state.baseXY.x, that.state.baseXY.y)){
                break;
              }
              let newArea1 = that.state.curArea.map((arr)=>{return arr.slice();});
              newArea1[cell1.getX()][cell1.getY()] = cell1;
              that.setState({
                curArea: newArea1,
                baseXY: that.state.baseXY,
                socket: that.state.socket,
              })
              break;
            default:
              alert('error: response no type')
          } 
      };
      fileReader.readAsArrayBuffer(blob);
    });

    this.state = {
      socket: socket,
      baseXY: {x: 0, y:0},
      curArea: [],
    };
  }

  handleClick(globX, globY) {
    // global x and global y
    let msg = new pb.ClientToServer();
    let touch = new pb.TouchRequest();
    touch.setX(globX);
    touch.setY(globY);
    touch.setTouchtype(pb.TouchType.FLIP);
    msg.setTouch(touch);
    this.state.socket.send(msg.serializeBinary());
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
            ()=>{
              this.handleClick(this.state.baseXY.x+i, this.state.baseXY.y+j)
            }
          }
        />
      });
      return  <div key={i} className="board-row">{cellRow}</div>
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
    >
      {props.value}
    </button>
  );
}


function Square(props) {
  return (
    <button
      className="square"
      onClick={
        props.onClick
      }
    >
      {props.value}
    </button>
  );
}

class Board extends React.Component {
  renderSquare(i) {
    return <Square 
      value={this.props.squares[i]}
      onClick={()=>{this.props.onClick(i)}}
    />;
  }

  render() {
    return (
      <div>
        <div className="board-row">
          {this.renderSquare(0)}
          {this.renderSquare(1)}
          {this.renderSquare(2)}
        </div>
        <div className="board-row">
          {this.renderSquare(3)}
          {this.renderSquare(4)}
          {this.renderSquare(5)}
        </div>
        <div className="board-row">
          {this.renderSquare(6)}
          {this.renderSquare(7)}
          {this.renderSquare(8)}
        </div>
      </div>
    );
  }
}

function calculateWinner(squares) {
  const lines = [
    [0, 1, 2],
    [3, 4, 5],
    [6, 7, 8],
    [0, 3, 6],
    [1, 4, 7],
    [2, 5, 8],
    [0, 4, 8],
    [2, 4, 6],
  ];
  for (var i = lines.length - 1; i >= 0; i--) {
    const [a, b, c] = lines[i];
    if(squares[a] && squares[a] === squares[b] && squares[a] === squares[c]) {
      return squares[a];
    }
  }
  return null;
}

class Game extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      history: [{
        squares: Array(9).fill(null),
      }],
      xIsNext: true,
      stepNumber: 0,
    };
  }

  handleClick(i) {
    const history = this.state.history.slice(0, this.state.stepNumber + 1);
    const current = history[history.length - 1];
    const squares = current.squares.slice();
    if (calculateWinner(squares) || squares[i]) {
      return;
    }
    squares[i] = this.state.xIsNext ? 'X' : 'O';
    this.setState({
        history: history.concat([{
        squares: squares,
      }]),
      xIsNext: !this.state.xIsNext,
      stepNumber: history.length,
    })
  }

  jumpTo(step) {
    this.setState({
      stepNumber: step,
      xIsNext: (step % 2) === 0,
    });
  }

  render() {
    const history = this.state.history;
    const current = history[this.state.stepNumber];
    const winner = calculateWinner(current.squares);

    const moves = history.map((step, move) => {
      const desc = move ?
        'Go to move #' + move :
        'Go to game start';
      return (
        <li key={move}>
          <button
            onClick={() => this.jumpTo(move)}
          >
            {desc}
          </button>
        </li>
      );
    })

    let status;
    if (winner) {
      status = 'Player ' + winner + ' won';
    } else {
      status = 'Next player: ' + (this.state.xIsNext ? 'X' : 'O');
    }

    return (
      <div className="game">
        <div className="game-board">
          <Board
            squares={current.squares}
            onClick={(i) => {this.handleClick(i)}}
          />
        </div>
        <div className="game-info">
          <div>{status}</div>
          <ol>{moves}</ol>
        </div>
      </div>
    );
  }
}

// ========================================

ReactDOM.render(
  // <Game />,
  <Aeilos />,
  document.getElementById('root')
);
