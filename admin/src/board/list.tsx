import {
  List,
  Datagrid,
  TextField,
  DeleteButton,
  FunctionField,
  ReferenceManyCount,
} from "react-admin";

import { renderGameState } from "./commons";

export const BoardList = () => (
  <List>
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
