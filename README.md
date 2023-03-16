# labyrinth

To run this program, you will need `go` installed. You can download it from [https://go.dev/dl/](https://go.dev/dl/).

## CLI

You can run the CLI program using:

```sh
# The following options are available :
# - SAVE_ID=<save-id>       To set the save ID name.
# - BOARD_SIZE=<size>       To set the board size.
# - PLAYER_COUNT=<count>    To set the numbe rof players.
make [options] cli-run
```

Note: You can clean existing saves using the following command:

```sh
make cli-clean
```

## Web application

/!\ Warning: Running webapplication requires `php8.1` and `composer` installed globally.

/!\ Warning: You will also need the following PHP extensions installed: `php8.1-xml`, `php8.1-mbstring`, `php8.1-zip`,
`php8.1-pgsql` and `php8.1-curl`.

To run the web application, you will need to create the environment files, and then set the corresponding environment
variables:

```sh
make setup-env

# This will create the following files:
webapp/.env
```

Once the env file has been set, you will need to create the `webapp/config/jwt/jwk.pub` file using [this tool](https://russelldavies.github.io/jwk-creator/) and the `webapp/config/jwt/public.pem` key.

You can now either run the app in development or production mode.

/!\ Note: The server automatically detects mobile devices and serves the mobile app automatically. You can test mobile mode using the developper tools from your browser.

### Development

```sh
make run
```

This will start the following applications:

```
webapp              https://localhost:9443
admin               https://localhost:9443/admin/
swagger             https://localhost:9443/admin/swagger/
```

### Production

```sh
make production
```

This will start the following applications:

```
webapp              http://0.0.0.0:80
webapp              https://0.0.0.0:443
admin               https://0.0.0.0:443/admin/
```

## Deploy to production

To deploy to production, you will need to:

- Add the SSH key to `.secrets/labyrinth-ed25519.pem` ;
- Create the `.secrets/.env` file:

```sh
export SERVER_USER=<server-user>
export SERVER_HOSTNAME=<server-domain>
```

You can the run the deploy command:

```sh
make production-deploy
```

## Test

You can run the unit tests for both domain and webapp using the following command:

```sh
make test
```

/!\ Warning: Running tests requires `go`, `node`, `npm`, `php8.1` and `composer` installed globally.

/!\ Warning: You will also need the following PHP extensions installed: `php8.1-xml`, `php8.1-mbstring`, `php8.1-zip`,
`php8.1-pgsql` and `php8.1-curl`.

## End to End testing

To run End-to-End tests, you will need to install Cypress dependencies, please see the [system requirements](https://docs.cypress.io/guides/getting-started/installing-cypress#System-requirements).

You will also need to setup the environment variables used by the test runner using an environment file and configure the user and password accordingly to your local configuration.

```sh
cp -n cypress.env.dist cypress.env
```

You then can the run end to end tests using:

```sh
make test-e2e
```
