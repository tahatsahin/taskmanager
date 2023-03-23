<b>Task Manager App with Go</b>

- This is a CRUD practice project in Go.
- Used technologies: Go, JWT, HTTP-Negroni-DotEnv-Gorilla-Crypto Packages, MongoDB
- Endpoints:
  - **/users: DELETE** - Delete registered user
  - **/users/register: POST** - Register user
  - **/users/login: POST** - Login with registered user credentials
  - **/tasks: GET** - Get all created tasks, note that all /tasks endpoints are restricted for unknown users.
  - **/tasks: POST** - Create a new task
  - **/tasks/{id}: GET** - Get a specific task by given id
  - **/tasks/{id}: PUT** - Update a specific task by given id
  - **/tasks/{id}: DELETE** - Delete task by given id

- Gorilla Mux is used for routing
- Negroni is used for chaining routes
- JWT is used for authorizing users
- Standart HTTP package is used for creating servers and handling routes
- Crypto package is used to generate and read private and public RSA

To be able to run the program, clone the repo by using 
```
git clone https://github.com/tahatsahin/taskmanager.git
```

After cloning, go to the directory that the project is cloned and run the below command.
```
go install
```

This command will install the necessary dependencies on your machine.
After completing installation, go to the directory 
```
{InstallationDirectory}/taskmanager/common
```
and create a file named .env by using the given info in sample.env file.

After saving .env file, you can start the program by using
```
go run main.go
```

Feel free to contact me at: <b>tahatsahin@gmail.com</b>