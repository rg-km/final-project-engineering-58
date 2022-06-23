# Dokumen Rest API

**Token Admin**
```json
eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjk5OTk5OTk5OTksInJvbGUiOiJhZG1pbiIsInN1YiI6IjhhZjA0NmU0LWYzZjgtNDlhYi1hYjAxLWM1ODUxMGNhYWIwOSJ9.hm19mQnskQAWERIIZ5ayKU3W8yWzymoOqRpJwE9nwyU
```

**Token Constributor**
```json
eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjk5OTk5OTk5OTksInJvbGUiOiJjb25zdHJpYnV0b3IiLCJzdWIiOiI4YWYwNDZlNC1mM2Y4LTQ5YWItYWIwMS1jNTg1MTBjYWFiMDkifQ.W3tc4xgyQI10n2fAz_yHDj0pAzMWEfxQFPzxDFIVeI8
```

**Token Member**
```json
eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjk5OTk5OTk5OTksInJvbGUiOiJtZW1iZXIiLCJzdWIiOiI4YWYwNDZlNC1mM2Y4LTQ5YWItYWIwMS1jNTg1MTBjYWFiMDkifQ._7IVCyIGCZ8Z2vuMGqQlmutmp1ONS67p0oH5YR4cpic
```


**Daftar Isi**
- [Dokumen Rest API](#dokumen-rest-api)
	- [Category](#category)
		- [Create Category](#create-category)
		- [Get All Category](#get-all-category)
		- [Update Category](#update-category)
		- [Delete Category](#delete-category)
	- [Posts](#posts)
		- [Create Post](#create-post)
		- [Get All Post](#get-all-post)
		- [Update Post](#update-post)
		- [Delete Post](#delete-post)
	- [Users](#users)
		- [Create User](#create-user)

## Category

Memerlukan header *Authorization* dengan role `admin`

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

## Posts

Memerlukan header *Authorization* dengan role `admin` atau `constributor`

### Create Post
`
POST: /api/post
`

Body Json:

```json
{
	"title": "Belajar Golang Basic",
	"description": "belajar golang",
	"content": "golang",
	"url_video": "https://youtu.be/M8zFCByLGIg",
	"category_id": "8016c26d-554d-4f11-b0d0-46e9557c2673"
}
```

Response:
```json
{
	"code": 201,
	"message": "Success",
	"data": {
		"id": "c69e4486-dbae-4e8a-8976-1c349771d60a",
		"title": "Belajar Golang Basic",
		"description": "belajar golang",
		"content": "golang",
		"url_video": "https://youtu.be/M8zFCByLGIg",
		"category_id": "8016c26d-554d-4f11-b0d0-46e9557c2673",
		"user_id": "8af046e4-f3f8-49ab-ab01-c58510caab09",
		"parent_id": "",
		"created_at": "2022-06-23T09:24:19.068798Z",
		"updated_at": "2022-06-23T09:24:19.0687997Z"
	}
}
```

### Get All Post
`
GET: /api/posts
`

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": [
		{
			"id": "c69e4486-dbae-4e8a-8976-1c349771d60a",
			"title": "Belajar Golang Basic",
			"description": "belajar golang",
			"content": "golang",
			"url_video": "https://youtu.be/M8zFCByLGIg",
			"category_id": "8016c26d-554d-4f11-b0d0-46e9557c2673",
			"user_id": "8af046e4-f3f8-49ab-ab01-c58510caab09",
			"parent_id": "",
			"created_at": "2022-06-23T09:24:19.068798Z",
			"updated_at": "2022-06-23T09:24:19.0687997Z"
		}
	]
}
```

### Update Post
`
PUT: /api/post/:id
`

Body Json:

```json
{
	"title": "Belajar Golang Basic",
	"description": "belajar golang dari awal",
	"content": "golang",
	"url_video": "https://youtu.be/M8zFCByLGIg",
	"category_id": "8016c26d-554d-4f11-b0d0-46e9557c2673"
}
```

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": {
		"id": "c69e4486-dbae-4e8a-8976-1c349771d60a",
		"title": "Belajar Golang Basic",
		"description": "belajar golang",
		"content": "golang",
		"url_video": "https://youtu.be/M8zFCByLGIg",
		"category_id": "8016c26d-554d-4f11-b0d0-46e9557c2673",
		"user_id": "8af046e4-f3f8-49ab-ab01-c58510caab09",
		"parent_id": "",
		"created_at": "2022-06-23T09:24:19.068798Z",
		"updated_at": "2022-06-23T09:24:19.0687997Z"
	}
}
```

### Delete Post
`
DELETE: /api/post/:id
`

Response:
```json
{
	"code": 200,
	"message": "Success",
	"data": null
}
```

## Users

Memerlukan header *Authorization* dengan role `admin`

### Create User
`
POST: /api/user
`

Body Json:

```json
{
	"name": "Febri Hidayan",
	"email": "febri@app.com",
	"password": "password",
	"role": "admin"
}
```

>Note: role boleh dimasukan apa saja seperti role `admin`, `constributor`, dan `member`

Response:
```json
{
	"code": 201,
	"message": "Success",
	"data": {
		"id": "43664e69-8baf-4c81-81a6-8791896c46d1",
		"name": "Febri Hidayan",
		"email": "febri@app.com",
		"role": "admin",
		"created_at": "2022-06-23 09:33:39.9423329 +0000 UTC",
		"updated_at": "2022-06-23 09:33:39.9423345 +0000 UTC"
	}
}
```
