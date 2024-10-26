### Requirements

To run this program, you need to have the following installed on your system:

1. **Docker**: Docker is a platform for developing, shipping, and running applications inside containers. You can install Docker by following the instructions on the [official Docker website](https://docs.docker.com/get-docker/).

2. **Docker Compose**: Docker Compose is a tool for defining and running multi-container Docker applications. You can install Docker Compose by following the instructions on the [official Docker Compose website](https://docs.docker.com/compose/install/).

3. **Go**: Go (also known as Golang) is a programming language designed for simplicity and efficiency. You can install Go by following the instructions on the [official Go website](https://golang.org/doc/install).

### Running the Program

#### Step 1: Clone the Repository

First, clone the repository containing the program:

```sh
git clone https://github.com/Gibsonguy911/gator.git
cd gator
```

#### Step 2: Connect to the Database

The program requires a PostgreSQL database to be running. You can either run a PostgreSQL database locally or use a Docker container to run the database.

In either instance, a config file is required wit the connection string to the database. An example is available in the root of the project and can be a skeleton for your own configuration.

```sh
cp config.example.json ~/.gatorconfig.json
```

#### Step 2a: Connect to a Local Database

Update ~/.gatorconfig.json with the connection string to your local PostgreSQL database.

Ensure the database is running and accessible on the specified host and port.

#### Step 2b: Build and Run with Docker Compose

To build and run the database using Docker Compose, follow these steps:

1. Ensure you are in the root directory of the repository where the `docker-compose.yaml` file is located.

2. Build and start the containers:

```sh
docker-compose up -d
```

This command start the containers in the background as defined in the `docker-compose.yaml` file.

3. To stop the containers, in the same directory, run:

```sh
docker-compose down
```

~/.gatorconfig.json is already configured to work with the Docker Compose setup.

#### Step 3: Install Go Dependencies and Build the Program

If you prefer to run the program directly on your host machine, you can install the Go dependencies and build the program:

1. Ensure you have Go installed and your `GOPATH` is set up correctly.

2. Navigate to the directory containing the Go source code.

3. Install the Go dependencies:

```sh
go mod tidy
```

4. Build the program:

```sh
go install
```

This will compile the program and place the executable in your `GOPATH/bin` directory.

5. Run the program:

```sh
gator help
```

### Conclusion

By following these steps, you should be able to set up and run the program using Docker, Docker Compose, and Go. If you encounter any issues, refer to the official documentation for Docker, Docker Compose, and Go for troubleshooting tips.
