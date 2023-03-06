import "./App.css";
import Board from "./components/Board";

import { Board as State } from "./model/Board";

function App() {
  const board: State = {
    tiles: [
      [
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "A",
          },
          rotation: 180,
        },
        {
          tile: {
            shape: 1,
            treasure: "N",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 1,
            treasure: "B",
          },
          rotation: 180,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 0,
        },
      ],
      [
        {
          tile: {
            shape: 2,
            treasure: "X",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: "U",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 1,
            treasure: "Q",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 270,
        },
      ],
      [
        {
          tile: {
            shape: 1,
            treasure: "C",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "D",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 1,
            treasure: "E",
          },
          rotation: 180,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "F",
          },
          rotation: 270,
        },
      ],
      [
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 1,
            treasure: "M",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 1,
            treasure: "P",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 90,
        },
      ],
      [
        {
          tile: {
            shape: 1,
            treasure: "G",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 2,
            treasure: "V",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 1,
            treasure: "H",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 1,
            treasure: "I",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "J",
          },
          rotation: 270,
        },
      ],
      [
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 2,
            treasure: "T",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 1,
            treasure: "O",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 270,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 0,
        },
      ],
      [
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 180,
        },
        {
          tile: {
            shape: 1,
            treasure: "R",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "K",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 2,
            treasure: "S",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 1,
            treasure: "L",
          },
          rotation: 0,
        },
        {
          tile: {
            shape: 0,
            treasure: ".",
          },
          rotation: 90,
        },
        {
          tile: {
            shape: 2,
            treasure: ".",
          },
          rotation: 90,
        },
      ],
    ],
    remainingTile: {
      tile: {
        shape: 2,
        treasure: "W",
      },
      rotation: 0,
    },
    players: [
      {
        color: 0,
        position: {
          line: 0,
          row: 0,
        },
        targets: ["J", "R", "U", "W", "N", "A"],
        score: 0,
      },
      {
        color: 1,
        position: {
          line: 6,
          row: 6,
        },
        targets: ["C", "T", "L", "H", "D", "P"],
        score: 0,
      },
      {
        color: 2,
        position: {
          line: 0,
          row: 6,
        },
        targets: ["B", "M", "X", "S", "E", "O"],
        score: 0,
      },
      {
        color: 3,
        position: {
          line: 6,
          row: 0,
        },
        targets: ["G", "K", "Q", "I", "F", "V"],
        score: 0,
      },
    ],
    remainingPlayers: [0, 1, 2, 3],
    currentPlayerIndex: 0,
    gameState: 0,
  };

  return (
    <main>
      <Board state={board} />
    </main>
  );
}

export default App;
