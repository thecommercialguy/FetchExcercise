# 
Fetch Reciept Processer Challange

## ⚙️ Installation

Inside a Go module:

```bash
go get github.com/thecommercialguy/FetchExcercise.git
```

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