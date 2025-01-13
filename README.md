# Fetch Receipt Processer Challange

A webservice that allows users to post receipts and get the points awarded for validated reciepts.

## ⚙️ Installation

```bash
git clone https://github.com/thecommercialguy/FetchExcercise.git
```

## 🚀 Run Server Locally

#### Prerequisites:
* Go (download from https://go.dev/dl/)

#### Steps:

1. Navigate to project directory:
```bash
cd fetchServer/
```

2. Start the server:
```bash
go run .
```

## 🚀 Run with docker

#### Steps:

1. Navigate to project directory:
```bash
cd fetchServer/
```

2. Create an image with "Dockerfile.multistage":
```bash
docker build -t fetch-server:multistage -f Dockerfile.multistage .
```

3. Run image in container: 
```bash
docker run fetch-server:multistage
```

## Run Tests Locally

#### Prerequisites:
* Go (download from https://go.dev/dl/)

#### Steps:

1. Navigate to project directory:
```bash
cd fetchServer/
```

2. Run unit tests:
```bash
go test ./...
```

## Run Tests with Docker

#### Steps:

1. Navigate to project directory:
```bash
cd fetchServer/
```

2. Run unit tests:
```bash
docker build -f Dockerfile.multistage -t fetch-server-test --progress plain --no-cache --target run-test-stage .
```

## Endpoints

### GET /receipts/{id}/points

```json
{
    "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```

Response body:

```json
{
    "points": 99
}
```

### POST /receipts/process

```json
{
  "retailer": "RetailerName",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "12:00",
  "items": [
    {
      "shortDescription": "item name ",
      "price": "10.99"
    }
  ],
  "total": "10.99"
}
```

Response body:

```json
{
    "id": "7fb1377b-b223-49d9-a31a-5a02701dd310"
}
```
