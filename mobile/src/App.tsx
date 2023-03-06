import "./App.css";

import Tile from "./components/Tile";
import { Rotation, Shape } from "./model/Tile";

function App() {
  return (
    <Tile
      boardTile={{
        tile: { treasure: "A", shape: Shape.ShapeT },
        rotation: Rotation.Rotation90,
      }}
      disabled={true}
      onClick={(e) => {
        console.log(e);
      }}
    />
  );
}

export default App;
