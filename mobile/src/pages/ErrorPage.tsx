import { useRouteError } from "react-router-dom";

interface ErrorMessage {
  message?: string;
}

export default function ErrorPage() {
  const error = useRouteError() as ErrorMessage;

  return (
    <div id="error-page">
      <h1>Oops!</h1>
      <p>Sorry, an unexpected error has occurred.</p>
      {error.message ? (
        <p>
          <i>{error.message}</i>
        </p>
      ) : (
        ""
      )}
    </div>
  );
}
