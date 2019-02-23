import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
const pb = require('./aeilos_pb');


class WStest extends React.Component {
  constructor(props) {
    super(props);
    const socket = new WebSocket('ws://localhost:8000/ws');
    socket.addEventListener('open', (event)=>{
      let msg = new pb.ClientToServer();
      let touch = new pb.TouchRequest();
      touch.setX(10);
      touch.setY(10);
      touch.setTouchtype(pb.TouchType.FLIP)
      msg.setTouch(touch);
      socket.send(msg.serializeBinary());
    });

    var that = this;
    socket.addEventListener('message', function (event) {
      var blob = event.data;
      var fileReader     = new FileReader();
      fileReader.onload  = function(event) {
          let msg = pb.ServerToClient.deserializeBinary(event.target.result);
          that.setState({
            hello: that.state.hello + ' ' + msg.getMsg(),
          });
      };
      fileReader.readAsArrayBuffer(blob);
    });



    this.state = {
      hello: "hello",
    };
  }

  render() {
    return (
      <div>
        {this.state.hello}
      </div>
    );
  }
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
  <WStest />,
  document.getElementById('root')
);
