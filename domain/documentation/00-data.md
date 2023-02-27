# Data structures

## Board

```ts
interface Board {

    // The tiles that are placed on the board.
    // It is represented as an NxN matrix, where N is the number of rows and columns.
    tiles: BoardTile[][]
}
```

### Tile that are placed on the board

```ts
interface BoardTile {

    // The corresponding tile.
    tile: Tile

    // The tile rotation on the board.
    rotation: Rotation    
}
```

### Tile rotation

```ts
enum Rotation {
    ROTATION_000 = 0,
    ROTATION_090 = 90,
    ROTATION_180 = 180,
    ROTATION_270 = 270,
}
```

## Tiles

```ts
interface Tile {

    // The tile shape.
    shape: Shape

    // The tile treasure.
    treasure: string?
}
```

### Shape

```ts
enum Shape {
    I,
    T,
    V,
}
```
