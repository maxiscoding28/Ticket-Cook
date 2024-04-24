<img src="logo.gif" width="125" height="125">

# Ticket Cook (tck) 

## Summary
Ticket Cook (tck) is a command line utility for managing ticket-based work. 

It is designed with support engineers in mind. 

Its goal is to help manage the complexity of working across multiple tickets at the same time.

Tck creates and manages three distinc directories that simplify the workflow for managing ticket work:

- `$TCK_HOME_DIR/`
    - This is the root directory for the application.
    - The default location for this directory is `$HOME/tck`.
    - This location can be changed by setting the environment variables `TCK_HOME_DIR`.

- `tickets/`
    - This directory stores open tickets that are being actively worked on.
    - Newly initialized ticket directories are saved as sub-directories this directory.

- `.closed/`
    - This directory stores closed tickets that are no longer being actively worked on.
    - Running the command `tck close <ticket>` will move a ticket from the `tickets/` directory to the `.closed/` directory.
    - Conversely, running the command `tck reopen <ticket>` will move the ticket from the `closed/` directory back to the `tickets/` directory.

- `recipes/`
    - This directory stores "recipe" directories.
    - Recipe directories can be used as templates when initializing new ticket directories.
    - When a ticket is initialized with a recipe, all files within the specific recipe directory are copied over to the newly created ticket folder.

**Example Directory structure:**
```
$TCK_HOME_DIR/
    tickets/
        /ticket1234
            customer-logs
            shell-script.sh
            customer-config.json
    .closed/
        /ticket98765
            resolution.txt
            customer-logs.json
            workaround.yaml
    recipes/
        default/
        sev-template/
            sev-cheatsheet.md
            outage-definitions.txt
            meeting-notes.txt
        dev-repro/
            quick-setup.sh
            k8-repro.yaml
        tls-scripts/
            demo-CA.pem
            demo-crt.pem
            demo-key.pem
            openssl-commands.sh
```
## Installation
```
curl
```

## Commands

### Bootstrap
`tck bootstrap`

- **Relevant environment variables**
    - `TCK_HOME_DIR` to your desired root directory. By default this will be `$HOME/tck`

### Initialize Ticket Directory
`tck init`

- **Relevant environment variables**
    - `TCK_HOME_DIR` to your current working tck home directory. Also consider setting this as a variable in you `.bashrc` or `.zshrc`. By default this will be `$HOME/tck`.

    - `TCK_RECIPE` to reference the recipe to be used in the ticket initialization. This value is overriden by the `--recipe` flag.

    - `TCK_URL_FORMAT` to define the format for the url. This must be an `http` or `https` value. It must contain a `@` character to denote the location that the ticket ID should be interpolated. By default this will be `https://hashicorp.zendesk.com/agent/tickets/@`. Thefore, for a ticket with ID `1234`. The url will navigate to `https://hashicorp.zendesk.com/agent/tickets/1234`.

    - `TCK_ID` to define a ticket ID to be initialized. This value is overriden by the passed argument. For example, in the following command `TCK_ID=9876 tck init 1234`, the ticket ID `1234` will be used.
- **Available flags**
```
-d, --description string   Provide a short description for the ticket you are creating.
-h, --help                 help for init
-r, --recipe string        Use the provided recipe file to initialize the ticket directory.
-u, --url-format string    Format for generating a URL with the ticket ID. Use @ for the ticket ID location
``` 

tck ls
    --closed, -c

tck get
    --nav, -n
    --all, -a
    --closed, -c

tck close

tck reopen

tck rm
    <!-- Need to add -->
    --closed, -c 

tck recipe init

tck recipe ls

tck recipe get
    --all, -a

tck recipe rm
