import { List, Datagrid, TextField, DeleteButton } from "react-admin";

export const BoardList = () => (
  <List>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <DeleteButton />
    </Datagrid>
  </List>
);
