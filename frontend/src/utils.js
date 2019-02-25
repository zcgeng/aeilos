const pb = require('./aeilos_pb');

export const ROW_LENGTH = 40;
export const ROW_HEIGHT = 20;

export function InsideArea(x, y, ax, ay) {
  return (x >= ax) && (y >= ay) && (x < ax+ROW_HEIGHT) && (y < ay+ROW_LENGTH);
}

export function getCellDesc(pbcell) {
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

export function cellIsBomb(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.BOMBS) && (pbcell.getBombs() === 9);
}

export function cellIsNumber(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.BOMBS) && (pbcell.getBombs() !== 9);
}

export function cellIsFlag(pbcell) {
  return (pbcell.getCelltypeCase()===pb.Cell.CelltypeCase.FLAGURL);
}

export function cellIsBombOrFlag(pbcell) {
  return cellIsBomb(pbcell) || cellIsFlag(pbcell);
}

export function getCellNumber(pbcell) {
  if(!cellIsNumber(pbcell)) {
    return 0;
  }
  return(pbcell.getBombs());
}