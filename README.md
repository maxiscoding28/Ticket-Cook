<img src="logo.gif" width="125" height="125">

# Ticket Cook (tck) 

### Installation

### Quick Start

### Managing Tickets

### Managing Recipes

### Environment Variables
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

 
