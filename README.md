# board-server
Server that responds with random position for a board of size 10x10

# How to run
Go to the folder containing this project
Then either `go run .`
Or
follow the instructions to make a build then run the build

You can then curl or go on browser to `http://localhost:8080/position` which should respond with a json containing the x and y values for the randomly generated position.
