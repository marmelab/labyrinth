import {
  Show,
  SimpleShowLayout,
  TextField,
  ReferenceOneField,
  ReferenceManyField,
  Datagrid,
  FunctionField,
  type RowClickFunction,
} from "react-admin";

import { renderGameState, renderScore } from "../board/commons";

const gameRowClick: RowClickFunction = (_id, _resource, record) =>
  `/board/${record.board_id}/show`;

export const UserShow = () => {
  return (
    <Show>
      <SimpleShowLayout>
        <TextField source="username" />
        <TextField source="email" />
        <ReferenceManyField
          label="Games"
          reference="player"
          target="attendee_id"
        >
          <Datagrid rowClick={gameRowClick} bulkActionButtons={false}>
            <TextField source="board_id" />
            <ReferenceOneField
              label="Game State"
              reference="board"
              source="board_id"
              target="id"
            >
              <FunctionField render={renderGameState} />
            </ReferenceOneField>
            <FunctionField label="Score" render={renderScore} />
          </Datagrid>
        </ReferenceManyField>
      </SimpleShowLayout>
    </Show>
  );
};
