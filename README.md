# devops-practice
This project is work in progress.

### Tasks accomplished: 
- [x] created a Go application which connects to a local running Dynamodb database and creates a `Movies` table in it. 
- [x] basic database operations have been written to get familiar with Dynamodb database fundamentals and how to use it with AWS SDK Go V2. 
- [x] created server and a couple of endpoints to show JSON http response in the form as requried by challenge. 
- [x] Dockerize the app and make it spin using `docker compose up`. 
- [ ] write tests for http handlers and Dynamdb operation functions using AWS SDK Go V2.
- [ ] Implement CI using Travis CI

### Steps to run the application
- After cloning the app, run `docker network create -d bridge mynet` to create a network on which the DynamoDB container and server container can connect upon.
- Run `docker compose up --build` to start the app. 
- Test the endpoints at `http://127.0.0.1/secret` and `http://127.0.0.1/health`. 
