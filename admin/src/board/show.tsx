import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  ReferenceManyCount,
  ReferenceOneField,
  ReferenceManyField,
  Datagrid,
  FunctionField,
  type RowClickFunction,
} from "react-admin";

import { renderGameState, renderPlayerColor, renderScore } from "./commons";

const playerRowClick: RowClickFunction = (_id, _resource, record) =>
  `/user/${record.attendee_id}/show`;

export const BoardShow = () => {
  return (
    <Show>
      <SimpleShowLayout>
        <TextField source="id" />
        <ReferenceManyCount
          label="Players"
          reference="player"
          target="board_id"
        />
        <FunctionField label="Game State" render={renderGameState} />
        <DateField source="updated_at" />
        <DateField source="created_at" />
        <ReferenceManyField
          label="Players"
          reference="player"
          target="board_id"
        >
          <Datagrid rowClick={playerRowClick} bulkActionButtons={false}>
            <TextField source="board_id" />
            <ReferenceOneField
              label="Username"
              reference="user"
              source="attendee_id"
              target="id"
            >
              <TextField source="username" />
            </ReferenceOneField>
            <FunctionField label="Color" render={renderPlayerColor} />
            <FunctionField label="Score" render={renderScore} />
            <TextField source="win_order" />
          </Datagrid>
        </ReferenceManyField>
      </SimpleShowLayout>
    </Show>
  );
};
