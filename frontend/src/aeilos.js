import React from 'react';
import ReconnectingWebsocket from 'reconnecting-websocket';
import { Area } from './area.js'
import { ChatBox } from './chatbox.js'
import {InsideArea, ROW_LENGTH} from './utils';
import './index.css';
const pb = require('./aeilos_pb');

class ScoreBoard extends React.Component {
  render() {
    return (
      <div>
        <div className="scoreboard">
          {this.props.score}
        </div>
      </div>
    );
  }
}

export class Aeilos extends React.Component {
  constructor(props) {
    super(props);
    const socket = new ReconnectingWebsocket('wss://changgeng.me/ws/');
    // const socket = new ReconnectingWebsocket('ws://localhost:8000/ws/');
    this.state = {
      socket: socket,
      x: Math.floor(Math.random() * 200)-100,
      y: Math.floor(Math.random() * 200)-100,
      score: 0,
      curArea: [],
      lastWheel: new Date().getTime(),
      chatData: [],
      chatMsg: '',
      chatUserName: 'somebody',
    };

    socket.addEventListener('open', (event)=>{
      let msg = new pb.ClientToServer();
      let xy = new pb.XY();
      xy.setX(this.state.x);
      xy.setY(this.state.y);
      msg.setGetarea(xy);
      socket.send(msg.serializeBinary());

      let msgGetStats = new pb.ClientToServer();
      let getStats = new pb.GetStats();
      getStats.setUsername("user1");
      msgGetStats.setGetstats(getStats);
      socket.send(msgGetStats.serializeBinary());
    });

    var that = this;

    socket.addEventListener('message', function (event) {
      var blob = event.data;
      var fileReader     = new FileReader();
      fileReader.onload = function(event) {
          let response = pb.ServerToClient.deserializeBinary(event.target.result);
          switch(response.getResponseCase()){

            case pb.ServerToClient.ResponseCase.TOUCH:
              let cell = response.getTouch().getCell();
              if(!InsideArea(cell.getX(), cell.getY(), that.state.x, that.state.y)){
                break;
              }
              let newArea = that.state.curArea.map((arr)=>{return arr.slice();});
              let {x, y} = that.glob2local(cell.getX(), cell.getY())
              newArea[x][y] = cell;
              that.setState({
                curArea: newArea,
                score: that.state.score + response.getTouch().getScore()
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
                x: response.getArea().getX(),
                y: response.getArea().getY(),
              })

              break;

            case pb.ServerToClient.ResponseCase.MSG:
              let chatData = that.state.chatData.slice();
              that.setState({
                chatData: chatData.concat([response.getMsg()]),
              });
              break;

            case pb.ServerToClient.ResponseCase.UPDATE:
              let cell1 = response.getUpdate();
              if(!InsideArea(cell1.getX(), cell1.getY(), that.state.x, that.state.y)){
                break;
              }
              
              // let newArea1 = that.state.curArea.map((arr)=>{return arr.slice();});
              let xy1 = that.glob2local(cell1.getX(), cell1.getY());
              that.state.curArea[xy1.x][xy1.y] = cell1;

              // var t0 = performance.now();
              /* TOOOOOO SLOW!!!! */
              // that.setState({
              //   curArea: newArea1,
              // })
              // var t1 = performance.now();
              // console.log("update takes: ", t1-t0);
              break;

            case pb.ServerToClient.ResponseCase.STATS:
              let stats = response.getStats();
              that.setState({
                score: stats.getScore(),
                // userName: stats.getUsername(),
              })
              break;
            default:
              alert('error: response no type')
          } 
      };
      fileReader.readAsArrayBuffer(blob);
    });
  }

  glob2local(x, y) {
    return {x: x - this.state.x, y: y - this.state.y};
  }

  renderArea() {
    return (<Area 
      baseXY={{x:this.state.x, y:this.state.y}}
      socket={this.state.socket}
      curArea={this.state.curArea}
    />)
  }

  handleWheel(e) {

    const thisWheel = new Date().getTime();
    if(thisWheel-this.state.lastWheel < 50) {
      return
    }
    this.setState({
      lastWheel: thisWheel,
    });
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

    if(xMove ===0 && yMove === 0) {
      return;
    }

    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(this.state.x + xMove);
    xy.setY(this.state.y + yMove);
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    // this.setState({
    //   x: this.state.x + xMove,
    //   y: this.state.y + yMove,
    // });
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
      x: this.state.x + xmoves[direction],
      y: this.state.y + ymoves[direction],
    });
  }

  handleCoordX(event) {
    if(event.target.value === "") {
      return
    }
    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(parseInt(event.target.value, 10));
    xy.setY(this.state.y);
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      x: parseInt(event.target.value, 10),
    });
  }

  handleCoordY(event) {
    if(event.target.value === "") {
      return
    }
    let msg = new pb.ClientToServer();
    let xy = new pb.XY();
    xy.setX(this.state.x);
    xy.setY(parseInt(event.target.value, 10));
    msg.setGetarea(xy);
    this.state.socket.send(msg.serializeBinary());
    this.setState({
      y: parseInt(event.target.value, 10),
    });
  }

  handleChatSend(event) {
    if(this.state.chatMsg === "") {
      return;
    }
    // send the message
    let msg = new pb.ClientToServer();
    let chatMsg = new pb.ChatMsg();
    chatMsg.setUsername(this.state.chatUserName);
    chatMsg.setMsg(this.state.chatMsg);
    chatMsg.setTime(new Date().getTime());
    msg.setChatmsg(chatMsg);
    this.state.socket.send(msg.serializeBinary());

    // after sending
    this.setState({chatMsg:""})
  }

  render() {
    return (
      <div className="aeilos">
        <div className="area-container">
        <div className="area" onWheel={(e)=>{e.preventDefault();this.handleWheel(e);}}>
          {this.renderArea()}
        </div>
        </div>

        <div className="controlplane">
          <ScoreBoard score={this.state.score}/>

          <div className="navbuttons">
            <div>
            Current location: ({this.state.x}, {this.state.y})
            </div>
            <button onClick={()=>{
              this.handleMoveMap(0);
            }}>
              ⬆
            </button>
            <button onClick={()=>{
              this.handleMoveMap(1);
            }}>
              ⬇
            </button>
            <button onClick={()=>{
              this.handleMoveMap(2);
            }}>
              ⬅
            </button>
            <button onClick={()=>{
              this.handleMoveMap(3);
            }}>
              ➡
            </button>
            <div>
              Jump to:
              <input className="inputcoord" type="number" value={this.state.x} onChange={this.handleCoordX.bind(this)} />
              <input className="inputcoord" type="number" value={this.state.y} onChange={this.handleCoordY.bind(this)} />
            </div>
          </div>
          <hr/>
          <ChatBox 
            onClick={()=>{this.handleChatSend();}}
            chatData={this.state.chatData}
            value={this.state.chatMsg}
            handleMsgChange={(e)=>{this.setState({chatMsg: e.target.value})}}
            handleUserName={(e)=>{this.setState({chatUserName: e.target.value})}}
          />
        </div>

        
      </div>
    );
  }
}