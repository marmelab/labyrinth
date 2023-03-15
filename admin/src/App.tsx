import { Admin, Resource } from "react-admin";

import postgrestRestProvider from "@promitheus/ra-data-postgrest";

import { UserList } from "./user";
import { authProvider } from "./auth";

function App() {
  return (
    <Admin
      authProvider={authProvider}
      dataProvider={postgrestRestProvider("/admin/api/v1/")}
    >
      <Resource name="user" list={UserList}></Resource>
    </Admin>
  );
}

export default App;
