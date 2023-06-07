# README #

## API Contract
find the api documentation at the following postman [link](https://lunar-firefly-676833.postman.co/workspace/Bootcamp~d5198276-f65f-4967-8870-3b05048d857c/collection/12975315-49435fb3-6b29-4465-9fad-aaa3a47678c5?action=share&creator=12975315)
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

## Coach Appointment Tech Spec

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
2. the db if data does not exist.
2. Slow query is expected due to point no 1
3. For recreational purposes, this design implement EC2 AWS 
Instances, please aware that this server is not performance oriented
4. Due to EC2 Geographical Location, some of the internet provider 
is blocking connection/port to the EC2 AWS Instances, please aware 
that your internet is capable connecting to the instances.
5. BRIvolution Wi-Fi Connection is not able to reach EC2 AWS Instances, 
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

### Service State Diagram ###
this diagram will explain the state of appointment through 
th every process at system to achieve the output state.

![state diagram](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/main_state.png)

As explain at state diagram. there are 3 action could be performed
within the system which are create appointment, approval of appointment,
and reschedule appointment. So to cover the business logic of each
action. there i provide the activity diagram for them.

### Create Appointment Activity Diagram ###
The `create appointment` action is the entry point of this system. 
where user create the appointment request, and before the appointment 
create there are several requirment to pass which are:
1. appointment schedule should match the schedule of coach.
2. appointment schedule should not intercept with another user
appointment date with status booked, waiting for coach approval, 
or waiting reschedule approval from user.

Here is the complete diagram for this endpoint.

![create appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/create-appointment/activity.png)

### Approval Of Appointment Activity Diagram ###
The `approval of appointment` action is used to approve any appointment
that on state of rescheduling and waiting for coach approval. at this
action, appointment could be rejected and approved. and before reschedule
the appointment, the schedule should fulfill the requirement same as
`create appointment` endpoint.

Here is the complete diagram for this endpoint.

![approval of appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/approval-appointment/activity.png)

###  Reschedule Appointment Activity Diagram ###
The `reschedule appointment` action is used by coach to propose a new
schedule to customer. 

Here is the complete diagram for this endpoint.

![reschedule appointment activity](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/reschedule-appointment/activity.png)

###  Data Flow Diagram ###
there are 5 table that dedicated on this services which are `main_users`
`main_roles`, `appointments`, `time_zone_name` and `schedules`. For
complete explaination of database structure you can follow this ddf.


![data flow diagram](https://gitlab.com/Nacute/fita-be-test/-/raw/main/_diagrams/dfd.finish.png)


![data flow diagram](https://gitlab.com/Nacute/fita-be-test/-/blob/main/_diagrams/dfd.finish.png)
