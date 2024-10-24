# Docs
## Running The Server
Before running the server you might want to clone or download the source code and create .env file from .env.template.  
The easiest way to run the server is by cloning or download the source code and run.
```
go run cmd/main.go # this is the case most of the time especially unix user
```

## Accessing the API endpoint docs
For now there is only one endpoint that is documented using openAPI spec, you can find it under the *api* directory.
### Publishing Messages
The service is a Restful API based so you can publish a message to a topic using http client like curl, postman, etc.
```
curl -X POST {BASE_URL}/sample-topic \
     -H "Content-Type: application/json" \
     -d '{"message": "hello notifiying"}'
```
### Subscribing Topic
Client or Subscriber can listen to a topic via SSE endpoint to get the latest message.
```
curl {BASE_URL}/sample-topic
```