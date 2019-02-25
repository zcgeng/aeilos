import React from 'react';

export function Cell(props) {
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