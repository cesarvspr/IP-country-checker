Avoxi Challenge

## Description

This project a microservice that is daily update with the GeoLite country dataset with an api endpoint to authorize login based on the informed ip address.

### The request
- Success example
```curl
curl --location --request GET 'localhost:8013/avoxi/login?ip=207.250.234.100&whitelist=US'
```
- Failed example
```curl
curl --location --request GET 'localhost:8013/avoxi/login?ip=32.199.77.157&whitelist=BR'
```
- Endpoint: /login GET

The ip should be a string sent as request params.
The white list should be an string containing ISO codes separated by comma, application isn't case sensitive.
#### Example
```
- whitelist=EUA,BR
- ip=35.199.77.157
```

### The response

If the ip address country location matches the given whitelist countries the response will be a 200 status. Otherwise a status 417 should be expected.

### How it works

At startup the application will setup a server application on port :8013 (can be easily changed). Should be ready in less than 5 seconds.
Then dataset is scraped from web and loaded inside memory at runtime. 
A cron job will run everyday at 00:30 to sync with the latest dataset available. 

### Why store in memory during runtime?

This application has response times between 4~3ms, due to not been necessary to access any external database.