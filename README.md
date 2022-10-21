# Go Database Samples

Various examples of using Go (1.19) to connect to databases.

Steps to run:
- Use the `databases` *.sql files to help with seeding data (if needed)
- Create the required database
- Create an App Service on Linux - Go - and add the required environment variables (in `/routes/getdata/getdata.go`, `/config/db/db.go`)

On local
- Point this to a local instance of the type of database you'd like to use, or a remote one
- Run `go run main.go`