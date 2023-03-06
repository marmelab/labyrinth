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

To run the web application, you will need to create the environment files, and then set the corresponding environment
variables:

```sh
make setup-env

# This will create the following files:
webapp/.env
```

Once the env file has been set,you can either run the app in development or poduction mode.

/!\ Note: The server automatically detects mobile devices and serves the mobile app automatically. You can test mobile
mode using the developper tools from your browser.

### Development

```sh
make run
```

This will start the following applications:

```
webapp              127.0.0.1:9000
api-domain          127.0.0.1:9001
```

### Production

```sh
make production
```

This will start the following applications:

```
webapp              127.0.0.1:80
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

/!\ Warning: Running tests requires `go`, `php8.1` and `composer` installed globally.

/!\ Warning: You will also need the following PHP extensions installed: `php8.1-xml`, `php8.1-mbstring`, `php8.1-zip`,
`php8.1-pgsql` and `php8.1-curl`.

## End to End testing

You will need to have `Firefox` installed and add the following configuration lines to the `webapp/.env` file:

```sh
# The webapp URL
WEBAPP_URL="http://localhost:9000"
```

The you need to run the full app stack in development mod using:

```sh
make run
```

You can the run end to end tests using:

```sh
make test-e2e
```
