module simple_crud

go 1.22.4 // Recommended to use Go 1.22.4 or newer for latest driver compatibility

require (
	github.com/julienschmidt/httprouter v1.3.0
	go.mongodb.org/mongo-driver v1.15.0 // This is the new official driver
)

require (
	// These are indirect dependencies of the mongo-driver,
	// they will be added automatically by 'go mod tidy'
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.16.0 // indirect
)

require golang.org/x/crypto v0.17.0 // indirect
