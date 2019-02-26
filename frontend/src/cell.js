import React from 'react';

export function Cell(props) {
  return (
    <div 
      className="cellframe"
      onClick={props.onClick}
      onContextMenu={props.onContextMenu}
      onMouseDown={props.onMouseDown}
      onMouseUp={props.onMouseUp}
      onDoubleClick={(e)=>{}}
      unselectable="on"
      // onSelectStart="return false;" 
    >
      <button
        className={props.className}
      >
        {props.value}
      </button>
    </div>
  );
}