import {
  List,
  Datagrid,
  TextField,
  EmailField,
  TextInput,
  ReferenceManyCount,
} from "react-admin";

const listFilters = [
  <TextInput source="username@ilike" label="Username" alwaysOn />,
  <TextInput source="email@ilike" label="Email" alwaysOn />,
];

export const UserList = () => (
  <List filters={listFilters}>
    <Datagrid>
      <TextField source="id" />
      <TextField source="username" />
      <EmailField source="email" />
      <ReferenceManyCount
        label="Games"
        reference="player"
        target="attendee_id"
      />
    </Datagrid>
  </List>
);
