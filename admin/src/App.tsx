import { Admin, Resource } from "react-admin";

import postgrestRestProvider from "@promitheus/ra-data-postgrest";

import { UserList } from "./user";

function App() {
  return (
    <Admin dataProvider={postgrestRestProvider("/admin/api/v1/")}>
      <Resource name="user" list={UserList}></Resource>
    </Admin>
  );
}

export default App;
