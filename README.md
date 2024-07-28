# Todo App API with Microservices
This project implements a simple Todo app API using Node.js and PostgreSQL, employing a microservice architecture. The application is divided into three separate services and A Go-based load balancer is used to distribute requests to the appropriate service:

- **Add and Select Task Service:** Handles creating new tasks and get all tasks.
- **Delete Task Service:** Handles deleting existing tasks.
- **Complete Task Service:** Handles marking tasks as completed.

### Load Balancer (Go): 
The load balancer is responsible for distributing the incoming requests to the appropriate service based on the endpoint.

- **Endpoint Get /:** Routes to Service One(Select).
- **Endpoint POST /add:** Routes to Service One(Add).
- **Endpoint DELETE /delete/:id:** Routes to Service Two.
- **Endpoint PUT /complete/:id:** Routes to Service Three.

### Diagram:
```

                +--------+
                | Client |
                +--------+
                    |
                    v
  +---------------------------------------+
  |         Load Balancer (Go)            |
  +---------------------------------------+
      |                |            |
      v                v            v
+-----------+      +--------+ +----------+
|Service    |      |Service | |Service   |
|One        |      |Two     | |Three     |
|(Add/Select|      |(Delete)| |(Complete)|
|Task)      |      |Task    | |Task      | 
+-----------+      +--------+ +----------+
          \           |         /
           \          |        /
            \         v       /
            +-------------+  /
            | PostgreSQL  | /
            | Database    |
            +-------------+
````