import { Admin, Resource } from "react-admin";

import postgrestRestProvider from "@promitheus/ra-data-postgrest";

import { BoardList } from "./board";
import { UserList } from "./user";
import { authProvider } from "./auth";

function App() {
  return (
    <Admin
      authProvider={authProvider}
      dataProvider={postgrestRestProvider("/admin/api/v1/")}
    >
      <Resource name="board" list={BoardList}></Resource>
      <Resource name="user" list={UserList}></Resource>
    </Admin>
  );
}

export default App;
