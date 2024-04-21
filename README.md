<img src="logo.gif" width="125" height="125">

# Ticket Cook (tck) 

## Summary
Ticket Cook (tck) is a command line utility for managing ticket-based work. It is designed with support engineers in mind. Its goal is to help manage the complexity of working across multiple tickets at the same time.

Tck creates and manages three distinc directories that simplify the workflow for managing ticket work:

- `$TCK_HOME_DIR/`: This is the root directory for the application. The default location for this directory is `$HOME/tck`. This location can be changed by setting the environment variables `TCK_HOME_DIR`.

- `tickets/`: This directory stores open tickets that are being actively worked on. Newly initialized ticket directories are saved as sub-directories this directory.

- `.closed/`: This directory stores closed tickets that are no longer being actively worked on. Running the command `tck close <ticket>` will move a ticket from the `tickets/` directory to the `.closed/` directory. Conversely, running the command `tck reopen <ticket>` will move the ticket from the `closed/` directory back to the `tickets/` directory.

- `recipes/`: This directory stores "recipe" directories. Recipe directories can be used as templates when initializing new ticket directories. It copies over the files within the recipe directory to the new ticket directory. This allows users to manage their often-used scripts, notes, "recipes" etc.. for support work in one place. A `recipe.json` file within the directory allows users to create new files as well.

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
            recipe.json/
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

## Quick Start
```
Choose location for home directory

tck bootstrap

tck init ticket

tck ls

tck create template

# populate template with example files

tck init ticket

tck close
```

### Full List of Commands and Flags
tck bootstrap

tck init
    --description, -d
    --url-format, -u
    --recipe, -r

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

## Environment Variables
`TCK_ID`
- **Description**: Defines the Ticket ID instead of manually specifying it as an argument to the CLI.
- **Example Usage**: `TCK_ID=1234 tck get`
- **Available for the following commands:**
    - `tck init`
    - `tck get`
    - `tck close`
    - `tck reopen`
    - `tck remove`

`TCK_RECIPE`
- `tck init`

`TCK_HOME_DIR`
- **All commands**

`TCK_EDITOR`
- `tck get`
- `tck recipe get`

`TCK_URL_FORMAT`
- `tck init`

 
