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
  useRecordContext,
} from "react-admin";

import Typography from "@mui/material/Typography";

import { renderGameState, renderPlayerColor, renderScore } from "./commons";

const playerRowClick: RowClickFunction = (_id, _resource, record) =>
  record.is_bot ? false : `/user/${record.attendee_id}/show`;

const PlayerName = () => {
  const record = useRecordContext();

  return record.is_bot ? (
    <Typography component="span" variant="body2">
      [🤖] Bot #{record.id}
    </Typography>
  ) : (
    <ReferenceOneField
      label="Username"
      reference="user"
      source="attendee_id"
      target="id"
    >
      [🤺] <TextField source="username" />
    </ReferenceOneField>
  );
};

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
            <PlayerName />
            <FunctionField label="Color" render={renderPlayerColor} />
            <FunctionField label="Score" render={renderScore} />
            <TextField source="win_order" />
          </Datagrid>
        </ReferenceManyField>
      </SimpleShowLayout>
    </Show>
  );
};
