import { List, Datagrid, TextField, EmailField, TextInput } from "react-admin";

import { CustomManyCount } from "../shared/CustomManyCount";

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
      <CustomManyCount
        label="Games"
        reference="board_user"
        target="user_id"
        sortBy="board_id"
        sortByOrder="ASC"
      />
    </Datagrid>
  </List>
);
