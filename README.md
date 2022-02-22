# minitwit
This is the minitwit project for the DevOps course. 

# README IS TBD


# Initializing PostgreSQL DB
With a PostgreSQL DB running in the background, initialize your env by sourcing with .env found in root.

1. Navigate into minitwit/migrations and run the following commands
2. Windows: 'go run . init' | Linux: 'go run *.go init'
3. Windows: 'go run .' | Linux: 'go run *.go'


Output should be the console returning that the DB is on migration 3.

# Building static assets
Navigate to minitwit/assets and run 'npm run build'

# Run the API
Navigate to minitwit/cmd/minitwit and run 'go run main.go', now the API should be running and you can navigate to the host supplied in the npm step.

# TODO: DOCKER AND VAGRANT STEPS WHEN THEY ARE DONE :)
