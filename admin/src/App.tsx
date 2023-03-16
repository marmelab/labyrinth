import { Admin, Resource } from "react-admin";

import postgrestRestProvider from "@promitheus/ra-data-postgrest";

import { BoardList, BoardShow } from "./board";
import { UserList, UserShow } from "./user";
import { authProvider } from "./auth";

function App() {
  return (
    <Admin
      authProvider={authProvider}
      dataProvider={postgrestRestProvider("/admin/api/v1/")}
    >
      <Resource name="board" list={BoardList} show={BoardShow}></Resource>
      <Resource name="user" list={UserList} show={UserShow}></Resource>
    </Admin>
  );
}

export default App;
