import { createBrowserRouter, RouterProvider } from "react-router-dom";

import ErrorPage from "./pages/ErrorPage";

import boardRoutes from "./pages/board";

const routes = [...boardRoutes];
const router = createBrowserRouter(
  routes.map((route) => {
    return {
      ...route,
      errorElement: <ErrorPage />,
    };
  })
);

function App() {
  return (
    <main>
      <RouterProvider router={router} />
    </main>
  );
}

export default App;
