# My Project Training
# For check website

1. Use command "cd src/healthcheck/"
2. Use command "go get ./..."
3. Use command "go run main.go test.csv"



A test.csv file is in a folder public/testfile

If you want to test another file, just drop your file on folder public/testfile,
And use command "go run main.go 'your file name'"

# If recieve Message "Fail to send the report via Healtcheck Report API"

1. Open new terminal on src/healthcheck/
2. Use command "go run main.go"
3. Use command "curl -X POST http://localhost:8080/api/line/authorization" on new terminal
4. Login by your line id
5. now, back to check website

# For test

1. Use command "go test -v healthcheck/test"

# Don't care a dockerfile because, i can't open host's browser from a docker container.
