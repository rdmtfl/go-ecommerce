# go-ecommerce
An e-commerce REST API written in Go and MongoDB. 


## Installation & Usage
### Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

### Usage
#### 1. Clone this repository
```bash
git clone https://github.com/rdmtfl/go-ecommerce
cd go-ecommerce/project
```

#### 2. Start the application
```bash
make build
```
or manually
```bash
docker-compose up --build -d
```

This API is available at: `http://localhost:8080/`.

## TODO

### Missing Endpoints
#### Order Endpoints
- `GET /orders`

#### Client Endpoints
- `POST /clients`
- `GET /clients`
- `GET /clients/{id}`
- `PUT /api/clients/{id}`

### Additional Improvements
- Add request validation
- Implement error handling
- Write integration tests
- Add API documentation

## License
This project is licensed under the MIT License.