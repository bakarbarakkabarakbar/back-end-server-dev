# README #

## API Contract
Kindly find the api documentation at the following 
postman [link](https://lunar-firefly-676833.postman.co/workspace/Bootcamp~d5198276-f65f-4967-8870-3b05048d857c/collection/12975315-49435fb3-6b29-4465-9fad-aaa3a47678c5?action=share&creator=12975315)

## How to Use API 
1. Open the link mentioned above, this should open 
the Postman Workspace after getting authorized 
2. Get the super admin JWT Bearer token by sending basic auth 
using provided request. This request should return JWT token 
in the "Authorization" header. Super admin credential mentioned below.
```
Id = "super-admin"
Password = "STRONG.password79"
```
3. Parse the JWT Bearer token via this [link](https://jwt.io/#debugger-io) and copy the header
and payload to folder "Auth Needed/Super Admin". 
```
- Algorithm HS256
- Secret for super admin = secret-key-super-admin 
- Payload = Copied from parsing in link above
Advance Configuration
- Request header prefix = leave blank
- JWT header = Copied from parsing in link above
```
4. You can access the request inside the folder since the request 
inherit the auth configuration from "Super Admin" Folder
5. You can create new admin and login with that credential with
repeating these process.
```
- Secret for admin = secret-key-admin
- Secret for customer = secret-key-customer
```

## Unsolved Issue/Algorithm/Improvement
1. Control authorization of admin that does not yet approved 
or active by admin.
2. Delete data from register approval when super admin delete 
actor data.
3. Flexible customer search based on the key given to 
the API. 
4. Parsing JWT Token with default prefix, current workaround is 
using an empty header
5. Relation between validation on table actors with
table registry approval. It should have mysql stored procedure
that can be automatically change the value of validation based on 
parameter on registry approval. 
6. Relation on activation on table actors, it should have some kind
of activity detection to determine, if this actor is still accessing 
the API.
7. Connection to External API when fetching customer data to mysql 
databases. There is noticeable performance impact regarding this import process, 
and it should be changed into scheduler or cronjob instead of each of API
hit.
8. Hardcoded env variables, dsn, secret-key, external API url
9. No docs for swagger API, if you want to support with swagger account, I could not
be any happier :)
10. Response with new JWT token when Authorized API hit
11. Isolate customer/actor data by their credentials
12. Timed/Flexible secret key generator for each role
13. Auto fetch JWT token in Postman using global variable
14. Memory optimization by using pointer

Feel free to contact me via issue or hit me via akbar.muhammadakbarmaulana@gmail.com

## Env

In Windows
```shell
setx MYSQL_SERVER_HOST localhost
```
```shell
setx MYSQL_SERVER_PORT 3306
```
```shell
setx MYSQL_SERVER_SCHEMA miniproject
```
```shell
setx MYSQL_SERVER_USER root
```
```shell
setx MYSQL_SERVER_PASSWORD 1234QWERasdf.
```
```shell
setx API_PORT 8081
```
```shell
setx SWAGGER_PORT 8082
```
```shell
setx APP_PORT 8080
```


## Setup

1. Install Go version 1.20.4
2. Install Gin-gonic
```shell
go get -u github.com/gin-gonic/gin
```
3. Install Gorm
```shell
go get -u gorm.io/gorm
```
4. Install Driver Mysql
```shell
go get -u gorm.io/driver/mysql
```
4. Install JWT
```shell
go get -u github.com/golang-jwt/jwt/v5
```

## Run

Use this command to run API app from root directory:

```shell
go run main.go
```

## Setup Mysql
Before creating mysql environment, then you can create mysql 
database instances and connect root to command line interface. 
And also to take a note about IP address of the mysql instances.
1. Create schema for Mysql databases
```
CREATE SCHEMA `miniproject` DEFAULT CHARACTER SET utf8mb4 ;
```
2. Create roles to give privileges on accessing database 
via MySQL Workbench or Dbeaver
```
CREATE USER 'super-admin' IDENTIFIED BY 'STRONG.password78';
GRANT ALL PRIVILEGES ON miniproject.* TO 'super-admin';
```
3. Create roles to give privileges golang service to access 
database
```
CREATE USER 'golang-service-account' IDENTIFIED BY 'STRONG.password79';
GRANT ALL PRIVILEGES ON miniproject.* TO 'golang-service-account';
```
4. Create table actors, customers, actor roles, register 
approvals, and actor sessions

```
CREATE TABLE actors(
	`id` BIGINT UNSIGNED,
    `username` VARCHAR(50),
    `password` VARCHAR(50),
    `role_id` INT UNSIGNED,
    `is_verified` ENUM('true','false'),
    `is_active` ENUM('true','false'),
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `actorsPK` PRIMARY KEY (`id`),
    CONSTRAINT `role_idFK` FOREIGN KEY (`role_id`) REFERENCES actor_roles(`id`)
);
```
```
CREATE TABLE customers(
	`id` BIGINT UNSIGNED,
    `first_name` VARCHAR(50),
    `last_name` VARCHAR(50),
    `email` VARCHAR(50),
    `avatar` VARCHAR(200),
	`created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT `customersPK` PRIMARY KEY (`id`)
);
```
```
CREATE TABLE actor_roles(
	`id` INT UNSIGNED,
    `role_name` VARCHAR(50),
    CONSTRAINT `actor_rolesPK` PRIMARY KEY (`id`)
);
```

```
CREATE TABLE register_approvals(
	`id` INT UNSIGNED,
    `admin_id` BIGINT UNSIGNED,
    `super_admin_id` BIGINT UNSIGNED,
	`status` VARCHAR(50),
    CONSTRAINT `register_approvalsPK` PRIMARY KEY (`id`)
);
```

```
CREATE TABLE actor_sessions (
  `id` INT UNSIGNED AUTO_INCREMENT,
  `user_id` INT UNSIGNED NOT NULL,
  `token` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `expires_at` TIMESTAMP,
  CONSTRAINT `register_approvalsPK` PRIMARY KEY (`id`));
```
5. Populate with initial data, this is required for 
authentication to API

```
INSERT INTO actor_roles(`id`,`role_name`) VALUES(1, 'super-admin'),(2, 'admin'), (3, 'customer');
```

```
INSERT INTO actors(`id`,`username`, `password`, `role_id`, `is_verified`, `is_active`) VALUES (1, "super-admin", "7fbe1732f8b44c15b88f0c1e4fe94fcd0c60ccec", 1, true, true);
```

---

## User Management Tech Spec

This README would normally document whatever steps are 
necessary to get your application up and running.

### Feature Description ###
there are four endpoint that could use for user management 
process, which are:
1. Customer Creation, Reading, Updating, and Deleting
2. Admin (With Auth) Creation, Reading, Updating, and Deleting
3. Super Admin (With Auth) Creation, Reading, Updating, and 
Deleting
4. Admin and Super Admin Authentication using Basic Auth

### Acceptance Criteria ###
1. Admins and super admin could register a customer at user 
management services
2. Admin could register as admin at user management services
3. Super admin could approve/reject admin registration at 
user management services
4. Super admin could see approval request at user management 
services
5. Admins and super admin  could login at user management 
services
6. Admins and super admin  could remove a customer data at 
user management services
7. Super admin  could remove a admin data at user management 
services
8. Super admin  could activate/deactivate a admin data at user 
management services
9. Admins and super admin could get all a customer data with 
parameter (search by name and email ) and pagination
10. Admins and super admin  could get all a admins data with 
parameter (search by username ) and pagination

### Further Behavior ###
1. Every time the admin gets a list of customers, service gets 
data from [link](https://reqres.in/api/users?page=2) and saves into 
the db if data does not exist.
2. Slow query is expected due to point no 1
3. For recreational purposes, this design implement EC2 AWS 
Instances, please aware that this server is not performance oriented
4. Due to EC2 Geographical Location, some of the internet provider 
is blocking connection/port to the EC2 AWS Instances, please aware 
that your internet is capable connecting to the instances.
5. Some of public Wi-Fi Connection is not able to reach EC2 AWS Instances, 
consider looking for other internet provider

### Architecture and Design ###
this service using onion architecture, there are 5 layers 
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:
1. **Entity**: this layer contains the domain model or entities
of the system. These are the core objects that 
represent the business concepts and rules.
2. **Repository**: This layer provides an interface for the 
application to access and manipulate the entities. 
It encapsulates the data access logic and provides
a way to abstract the database implementation details.
3. **Use case** : This layer contains the business logic 
or use cases of the system. It defines the operations 
that can be performed on the entities and orchestrates 
the interactions between the entities and the repository layer.
4. **Controller**: This layer handles the HTTP requests and
responses. It maps the incoming requests to the appropriate 
use case and returns the response to the client.
5. **Request handler**: This layer is responsible for handling 
the incoming HTTP requests and passing them on to 
the controller layer.

### Class State Diagram ###
this diagram will explain the relation in the class, available
interface, and connection between architecture and design

![class diagram](https://github.com/bakarbarakkabarakbar/back-end-server-dev/blob/6bc209b5dc2a6aa806a3491df071946ab696d60d/_diagram/class_diagram_backend.png)

There are three main endpoint that could be used in the API
1. Customer, this endpoint serve about the customer data by limited key value
2. Admin, this endpoint serve the user data management
3. Super Admin, this endpoint serve ultimate admin and user data management

### Customer Available Endpoint ###

![customer available endpoint](https://github.com/bakarbarakkabarakbar/back-end-server-dev/blob/1b238d2c409a4636746cf6d31a791ced45942915/_diagram/customer/activity.png)

### Admin Available Endpoint ###

![admin available endpoint](https://github.com/bakarbarakkabarakbar/back-end-server-dev/blob/1b238d2c409a4636746cf6d31a791ced45942915/_diagram/admin/activity.png)

###  Super Admin Available Endpoint ###

![super admin available endpoint](https://github.com/bakarbarakkabarakbar/back-end-server-dev/blob/1b238d2c409a4636746cf6d31a791ced45942915/_diagram/super-admin/activity.png)

###  Table Diagram ###
there are 5 table that dedicated on services which are `actors`
`actor_roles`, `customers`, `actor_sessions` and `register_approvals`. For
complete of database structure you can check image below.

![tables diagram](https://github.com/bakarbarakkabarakbar/back-end-server-dev/blob/1b238d2c409a4636746cf6d31a791ced45942915/_sql/mysql-tables.png)

