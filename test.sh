# Tickets

# bootstrap
export TCK_HOME_DIR="/Users/maxwinslow/dev/sandbox-go/tck/test"
rm -rf $TCK_HOME_DIR
tck bootstrap

## overwrite bootstrap
tck bootstrap
tck recipe init test

# init
tck init my-test-ticket

## overrwrite init
tck init my-test-ticket

## --url-format, -u
tck init my-test-ticket2 --url-format="https://google.com/@"
tck init my-test-ticket3 --url-format="http://google.com/@"
tck init my-invalid-url-ticket --url-format="google.com"
tck init my-invalid-url-ticket --url-format="http://google.com"

## --description, -d

## --recipe, -r

## Env vars

# close
tck ls
tck ls -c
tck init my-test-ticket
tck ls
tck ls -c
tck close invalid-ticket
tck close my-test-ticket
tck ls 
tck ls -c

tck reopen invalid-ticket
tck reopen my-test-ticket
tck ls 
tck ls -c

# rm
tck rm my-test-ticket
tck ls
## env vars

# get
tck init my-test-ticket
tck get invalid-ticket
tck get my-test-ticket
tck get -a
## --nav, -n
## env vars




