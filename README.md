# go-start

## How to start
### Prerequisites
- Golang & PostgreSQL installed

Or
- Docker installed

### Docker
If you prefer to run with docker on your local, you can use this command
```sh
make docker-up
```
Your local is good to go (can skip the later part)

You can build production image (use production environment)
```sh
docker build -t {image_name} .
```

### Import Data Schema
Create a database, then import these to your database
- import Data Definition
```sh
psql [database_name] -f ./core/database/data/ddl.sql
```

### Set Environment
Put your database environment in here `./env/env.development.yaml`

### Install Dependencies
```sh
go get
```

### Run Apps
```sh
make run
```

### Project Structure
This project has Handler, Provider, Factory, Business Model, Repository, Entity, Data Source.

#### Overview Diagram
![GoStartOverview-Project Structure](https://user-images.githubusercontent.com/37319946/144997923-e1b6c456-fcaa-4f11-89e5-fab362c6136d.jpg)

- Handler

Handler is a layer that will convert the request to the object data of the Business Model. The Business Model receives data as Golang struct, and Handler will convert from either JSON, Proto buff, MQ Message, etc. depending on the request.

- Provider

Provider provides every dependency that we need. Because we use Dependency Injection, the dependency will be initialized in this layer then will be injected into the factory.

- Factory

Factory produces Business Models with the dependencies we need. Provider will inject the dependencies into Factory, then we forward the dependencies into our Business Model.

- Business Model

Business Model is the core / main logic of the business.

- Entity

Entity is the single object data of the business.

- Repository

Repository is a layer of common logic functionality to access the data source.

- Data Source

Data Source is where the data has been stored.

#### App Structure

This is how the application runs from the start until it shutdown.
![GoStartOverview-App Structure](https://user-images.githubusercontent.com/37319946/144998023-086ef2a4-e52c-4310-94cf-5632822454d6.jpg)
