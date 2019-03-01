import React from 'react';

export function Cell(props) {
  return (
    <div 
      className="cellframe"
      onClick={props.onClick}
      onDoubleClick={props.onDoubleClick}
      onContextMenu={props.onContextMenu}
      onMouseDown={props.onMouseDown}
      onMouseUp={props.onMouseUp}
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