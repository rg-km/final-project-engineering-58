# Dokumen Rest API

**Daftar Isi**
- [Dokumen Rest API](#dokumen-rest-api)
	- [Category](#category)
		- [Create Category](#create-category)
		- [Get All Category](#get-all-category)
		- [Update Category](#update-category)
		- [Delete Category](#delete-category)

## Category
### Create Category
`
POST: /api/category
`

Body Json:

```json
{
    "name": "Backend Developer"
}
```

Response:
```json
{
	"code": 201,
	"message": "Success",
	"data": {
		"id": "6f4c3579-7399-41a7-9877-bc9fb2a0bfb1",
		"name": "Backend Developer",
		"created_at": "2022-06-15T07:45:50.9257492Z",
		"updated_at": "2022-06-15T07:45:50.9257513Z"
	}
}
```

### Get All Category
`
GET: /api/category
`

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": [
		{
			"id": "6f4c3579-7399-41a7-9877-bc9fb2a0bfb1",
			"name": "Backend Developer",
			"created_at": "2022-06-15T07:45:50.9257492Z",
			"updated_at": "2022-06-15T07:45:50.9257513Z"
		}
	]
}
```

### Update Category
`
PUT: /api/category/:id
`

Body Json:

```json
{
    "name": "Software Developer"
}
```

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": {
		"id": "6f4c3579-7399-41a7-9877-bc9fb2a0bfb1",
		"name": "Software Developer",
		"created_at": "2022-06-15T07:45:50.9257492Z",
		"updated_at": "2022-06-15T07:45:50.9257513Z"
	}
}
```

### Delete Category
`
DELETE: /api/category/:id
`

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": null
}
```

## Auth
### Register
`
POST: /api/auth/register
`

Body JSON:
```json
{
    "name": "Ruang",
    "email": "Ruang@ac.id",
    "password": "Ruang123"
}
```

Response:
```json
{
    "code": 201,
    "message": "Success",
    "data": {
        "id": "e6cdf8d6-d380-4880-8042-e12277d2e0da",
        "name": "Ruang",
        "email": "Ruang@.ac.id",
        "role": "member",
        "created_at": "2022-06-16 14:15:55.757433367 +0000 UTC",
        "updated_at": "2022-06-16 14:15:55.757433457 +0000 UTC"
    }
}
```

### Login
`
POST: /api/auth/login
`

Body JSON:
```json
{
    "email": "Ruang@ac.id",
    "password": "Ruang123"
}
```

Response:
```json
{
    "code": 201,
    "message": "Success",
    "data": {
        "id": "e6cdf8d6-d380-4880-8042-e12277d2e0da",
        "name": "Ruang",
        "email": "Ruang@.ac.id",
        "role": "member",
        "created_at": "2022-06-16 14:15:55.757433367 +0000 UTC",
        "updated_at": "2022-06-16 14:15:55.757433457 +0000 UTC"
    }
}
```






