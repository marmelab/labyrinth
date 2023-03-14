import {
  List,
  Datagrid,
  TextField,
  EmailField,
  TextInput,
  ReferenceInput,
} from "react-admin";

const listFilters = [
  <TextInput source="username@ilike" label="Username" alwaysOn />,
  <TextInput source="email@ilike" label="Email" alwaysOn />,
];

export const UserList = () => (
  <List filters={listFilters}>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="username" />
      <EmailField source="email" />
    </Datagrid>
  </List>
);
