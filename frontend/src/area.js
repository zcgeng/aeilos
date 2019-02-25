import React from 'react';
import {cellIsBombOrFlag,ROW_HEIGHT,ROW_LENGTH,getCellNumber,getCellDesc} from './utils'
import {Cell} from './cell'
const pb = require('./aeilos_pb');

export class Area extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      mouseDown: {l:false, r:false},
    };
  }

  glob2local(x, y) {
    return {x: x - this.props.baseXY.x, y: y - this.props.baseXY.y};
  }

  handleMouseDown(e) {
    if(e.nativeEvent.which === 1){
      this.setState({
        mouseDown: {
          l:true,
          r:this.state.mouseDown.r,
        },
      });
    }
    if(e.nativeEvent.which === 3){
      this.setState({
        mouseDown: {
          l:this.state.mouseDown.l,
          r:true,
        },
      });
    }
  }

  // get the bomb number shown on the map
  getNeighbourBombs(locX, locY) {
    let bombCount = 0;
    for(let i = -1; i < 2; i++){
      for(let j = -1; j < 2; j++){
        if(cellIsBombOrFlag(this.props.curArea[locX+i][locY+j])){
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
    let cell = this.props.curArea[x][y]
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
      this.setState({
        mouseDown: {
          l:false,
          r:this.state.mouseDown.r,
        },
      });
    }
    if(e.nativeEvent.which === 3){
      this.setState({
        mouseDown: {
          l:this.state.mouseDown.l,
          r:false,
        },
      });
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
    if(getCellDesc(this.props.curArea[x][y]) !== ' '){
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
    if(getCellDesc(this.props.curArea[x][y]) !== ' '){
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
    if(getCellDesc(this.props.curArea[x][y]) !== ' '){
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
    const mmap = this.props.curArea.map((row)=>{
      return row.map((cell)=>{
        return getCellDesc(cell);
      });
    });

    const cellBoard = mmap.map((row, i) => {
      const cellRow = row.map((val, j) => {
        return <Cell 
          key={j} 
          value={val} 
          x={this.props.baseXY.x+i} 
          y={this.props.baseXY.y+j}
          onClick={
            (event)=>{
              event.preventDefault();
              this.handleClick(this.props.baseXY.x+i, this.props.baseXY.y+j, event);
            }
          }
          onContextMenu={
            (event)=>{
              event.preventDefault();
              this.handleClick(this.props.baseXY.x+i, this.props.baseXY.y+j, event);
            }
          }
          onMouseDown={this.handleMouseDown.bind(this)}
          onMouseUp={
            (event)=>{
              this.handleMouseUp(this.props.baseXY.x+i, this.props.baseXY.y+j, event);
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