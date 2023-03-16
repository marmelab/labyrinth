import {
  List,
  Datagrid,
  TextField,
  DeleteButton,
  FunctionField,
  ReferenceManyCount,
  SelectInput,
} from "react-admin";

import { gameState, renderGameState } from "./commons";

const listFilters = [
  <SelectInput
    label="Game State"
    source="game_state@in"
    choices={[
      {
        id: `(${gameState.placeTile}, ${gameState.movePawn})`,
        name: "On Going",
      },
      { id: `(${gameState.completed})`, name: "Completed" },
    ]}
    alwaysOn
  />,
];

export const BoardList = () => (
  <List filters={listFilters}>
    <Datagrid rowClick="show">
      <TextField source="id" />
      <ReferenceManyCount
        label="Players"
        reference="player"
        target="board_id"
      />
      <FunctionField label="Game State" render={renderGameState} />
      <DeleteButton />
    </Datagrid>
  </List>
);
