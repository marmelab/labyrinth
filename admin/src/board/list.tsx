import {
  List,
  Datagrid,
  TextField,
  DeleteButton,
  FunctionField,
  ReferenceManyCount,
  SelectInput,
} from "react-admin";

import { renderGameState } from "./commons";

const listFilters = [
  <SelectInput
    label="Game State"
    source="game_state@in"
    choices={[
      { id: "(0, 1)", name: "On Going" },
      { id: "(2)", name: "Completed" },
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
