# 
Fetch Reciept Processer Challange

## ⚙️ Installation

Code editor / CLI:

```bash
git clone https://github.com/thecommercialguy/FetchExcercise.git
```

## 🚀 Run Server

Ensure "\fetchServer" is your PWD:

```bash
go run .
```

## 🚀 Run with docker

Create an image with "Dockerfile":

```bash
docker build . -t <image-name>:latest
```

Run image in container: 

```bash
docker run -p 8080:8080 <image-name>:latest
```

## Unit testing
Ensure "\fetchServer" is your PWD:

```bash
go test ./...
```

## Endpoints

### GET /reciepts/{id}/points

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

### POST /reciepts/process

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


