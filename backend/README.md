# Heroes App

### Build and run

In order to run the application, you should first run mongo daemon process 
(that manages all MongoDB server tasks, e.g. accepting client connections and responding to them):

```bash
$ cd src

#creates the directory where the database will store our documents/collections
$ mkdir -p data/db

#create mongo daemon process
$ mongod --dbpath data/db
```
The mongo daemon process should be running in the current terminal.

Then, we can run our server by opening a new terminal and doing:
```bash
$ cd src

#Compiles the program
$ go build

#Runs the program
$ ./heroes
```

Now, we can do requests to our server :)

### Run unit tests
In order to run the unit tests, you should start a mongo daemon process and the server. Then run the file with the tests by doing:
```bash
$ cd src

#The -v flag produces detailed outputs about passed and failed tests
$ go test -v router/hero_test.go
```