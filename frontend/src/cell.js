import React from 'react';

export function Cell(props) {
  return (
    <div className="cellframe">
      <button
        className={props.className}
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
    </div>
  );
}