# Todo App API with Microservices
This project implements a simple Todo app API using Node.js and PostgreSQL, employing a microservice architecture. The application is divided into three separate services and A Go-based load balancer is used to distribute requests to the appropriate service:

- **Add and Select Task Service:** Handles creating new tasks and get all tasks.
- **Delete Task Service:** Handles deleting existing tasks.
- **Complete Task Service:** Handles marking tasks as completed.

### API Gateway: 
Handles routing requests to the appropriate backend service. Contains a Load Balancer for routing and health checks.

### Load Balancer (Go): 
The load balancer is responsible for distributing the incoming requests to the appropriate service based on the endpoint.

- **Endpoint Get /:** Routes to Service One(Select).
- **Endpoint POST /add:** Routes to Service One(Add).
- **Endpoint DELETE /delete/:id:** Routes to Service Two.
- **Endpoint PUT /complete/:id:** Routes to Service Three.

### Diagram:
````
                                            +--------+
                                            | Client |
                                            +--------+
                                                |
                                                v
                                    +-------------------------+
                                    |    Load Balancer (Go)   |
                                    |       API Gateway       |
                                    | +---------------------+ |
                                    | | Health Check Thread | |
                                    | +---------------------+ |
                                    | +---------------------+ |
                                    | | /                     |
                                    | | /add                  |
                                    | | /delete/              |
                                    | | /complete/            |
                                    | +---------------------+ |
                                    +----------|--------------+
                                               |
                                               |
               +-------------------------------+-------------------------------+
               |                               |                               |
               v                               v                               v
   +---------------------+           +---------------------+           +---------------------+
   | Add/Select Server   |           | Delete Server       |           | Complete Server     |
   |                     |           |                     |           |                     |
   | +--------------+    |           | +--------------+    |           | +--------------+    |
   | | Server 1     |    |           | | Server 1     |    |           | | Server 1     |    |
   | +--------------+    |           | +--------------+    |           | +--------------+    |
   | +--------------+    |           | +--------------+    |           | +--------------+    |
   | | Server 2     |    |           | | Server 2     |    |           | | Server 2     |    |
   | +--------------+    |           | +--------------+    |           | +--------------+    |
   +---------------------+           +---------------------+           +---------------------+
                    \                             |                       /
                     \                            |                      /
                      \                           |                     /
                       \                          |                    /
                        \                         |                   /
                         \                        v                  /
                    +----------------------------------------------------------+
                    |                   PostgreSQL Database                    |
                    +----------------------------------------------------------+


````