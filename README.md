# FINAL PROJECT DIGITALENT X HACKTIV8
it's a final project for digitalent x hacktiv8. Here's a [Final Project Requirement](https://drive.google.com/file/d/1Jk3AINh3o9d8ukqaqjTWKSWzFSOt_Nl2/view?usp=sharing) task details:

```text
Pada final project ini, kalian akan diminta untuk membuat suatu aplikasi bernama MyGram, yang  dimana pada aplikasi ini kalian dapat menyimpan foto maupun membuat comment untuk foto orang  lain. Aplikasi ini akan dilengkapi dengan proses CRUD dengan table dan alur yang akan dijelaskan  berikut ini:
1.Project ini bebas dikerjakan dengan library apapun. Namun agar proses pengerjaannya lebih cepat dan mudah, disarankan untuk menggunakan framework Gin Gonic dan orm Gorm.
2.Berikut merupakan library/package yang wajib digunakan
- github.com/dgrijalva/jwt-go
- golang.org/x/crypto
3.Dalam project ini akan memerlukan 4 table. 
- Tabel User
- Tabel Photo
- Tabel Comment
- Tabel Social Media
```

# Documentation

### `POST` users/register
it's a register user endpoint. This endpoint can be accessed by everyone.

**Request Body** :
```json
{
    "username" : "string", # must be a unique
    "email" : "string", # must be a valid email, and unique
    "password" : "string", #must be longer than or equal to 6 char  
	"age" : "int" #must be greater than or equal to 8 years
}
```
**Notes** : On this endpoint, `password` will be hashed using `crypto` library 

**Response Body** :
- success response:

- failed response:


### `POST` users/login
it's a login user endpoint. This endpoint can be accessed by everyone.

**Request Body** :
```json
{
    "email" : "string", # must be a valid email
    "password" : "string",
}
```
**Notes** : On this endpoint, `email` can be replace with `username` 

**Response Body** :
- success response:

- failed response:

## photos endpoint

### `POST` photos
it's a create photo endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "title" : "string",
    "caption" : "string",
    "photo_url" : "string",
}
```

**Response Body** :
- success response:

- failed response:


### `PUT` photos/:photoId
it's a update photo endpoint. This endpoint only can be accessed by `user` that created this photo. 

**Params**
- photoId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "title" : "string",
    "caption" : "string",
    "photo_url" : "string",
}
```

**Response Body** :
- success response:

- failed response:


### `DELETE` photos/:photoId
it's a delete photo endpoint. This endpoint only can be accessed by `user` that created this photo.

**Params**
- photoId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` photos
it's a get all photo endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:


### `GET` photos/me
it's a get all photo {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` photos/:photoId
it's a get photo by photoId endpoint. This endpoint only can be accessed by `user` that logged in.

**Params**
- photoId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:


## comments endpoint

### `POST` photos/:photoId/comment
it's a create comment on photo endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "message" : "string",
}
```

**Response Body** :
- success response:

- failed response:


### `PUT` comments/:commentId
it's a update comment endpoint. This endpoint only can be accessed by `user` that created this comment. 

**Params**
- commentId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "message" : "string",
}
```

**Response Body** :
- success response:

- failed response:


### `DELETE` comments/:commentId
it's a delete comment endpoint. This endpoint only can be accessed by `user` that created this comment.

**Params**
- commentId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` comments
it's a get all comment endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:


### `GET` comments/me
it's a get all comment {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` comments/:commentId
it's a get comment by commentId endpoint. This endpoint only can be accessed by `user` that logged in.

**Params**
- commentId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:


## socialmedia endpoint

### `POST` socialmedia
it's a create social media endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "name" : "string",
    "social_media_url" : "string",
}
```
**Notes:** `user` only can create 1 social media account.

**Response Body** :
- success response:

- failed response:


### `PUT` socialmedia/:sosmedId
it's a update social media endpoint. This endpoint only can be accessed by `user` that created this social media. 

**Params**
- sosmedId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Request Body** :
```json
{
    "name" : "string",
    "social_media_url" : "string",
}
```

**Response Body** :
- success response:

- failed response:


### `DELETE` socialmedia/:sosmedId
it's a delete social media endpoint. This endpoint only can be accessed by `user` that created this social media.

**Params**
- sosmedId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` socialmedia
it's a get all social media endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:


### `GET` socialmedia/me
it's a get social media {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response:

### `GET` socialmedia/:sosmedId
it's a get social media by sosmedId endpoint. This endpoint only can be accessed by `user` that logged in.

**Params**
- sosmedId : `int` | required

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:

- failed response: