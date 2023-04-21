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
![register_success](https://user-images.githubusercontent.com/37493831/233639710-6e1fd28b-7a9e-4a5a-91b8-710de97ace64.png)

- failed response:
![register_failed_age](https://user-images.githubusercontent.com/37493831/233639700-d33f9823-c8b6-4b54-a485-6c5094308871.png)
![register_failed_duplicate](https://user-images.githubusercontent.com/37493831/233639703-fdf977a1-d7bc-440f-9396-a9dd08a07657.png)
![register_failed_pasw](https://user-images.githubusercontent.com/37493831/233639705-26fc6f6f-064a-4323-902c-6c9853405d08.png)

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
![login_success](https://user-images.githubusercontent.com/37493831/233639691-316e01d1-ad17-40dc-b7e5-cbd959bbbb2f.png)
![login_success_username](https://user-images.githubusercontent.com/37493831/233639696-2969af37-e1f1-4510-ab6c-400b0967eb71.png)

- failed response:
![login_failed_email](https://user-images.githubusercontent.com/37493831/233639675-1614f1cc-2a3b-45ad-b8a3-eb29f2674a2c.png)
![login_failed_pasw](https://user-images.githubusercontent.com/37493831/233639687-667586c8-e634-4b7c-b963-113e7feaab29.png)

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
![create_photo_success](https://user-images.githubusercontent.com/37493831/233639736-541486d0-8c84-40af-b3aa-082ec95412b5.png)

- failed response:
![create_photo_failed_auth](https://user-images.githubusercontent.com/37493831/233639733-be17e142-1ed5-4cd8-967d-9b47c7d632c1.png)

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
![update_photo_success](https://user-images.githubusercontent.com/37493831/233639729-13e0965c-e7b5-42ce-8662-74ba5f9c94be.png)

- failed response:
![update_photo_failed_auth](https://user-images.githubusercontent.com/37493831/233639714-f37a3347-aea2-4bd9-9350-9e0385a46870.png)
![update_photo_failed_not_found](https://user-images.githubusercontent.com/37493831/233639717-0fa58d38-25b3-4bd1-9cc2-f3a42f9a7e9f.png)
![update_photo_failed_unauth](https://user-images.githubusercontent.com/37493831/233639721-59da2a82-692c-4295-bc94-46d260d270e0.png)

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
![delete_photo_success](https://user-images.githubusercontent.com/37493831/233639749-06f5896d-ea17-45e0-9b16-9127b600f323.png)

- failed response:
![delete_photo_failed_auth](https://user-images.githubusercontent.com/37493831/233639740-65ee490f-8fc9-4bc4-8fee-7d7f76ec73ea.png)
![delete_photo_failed_auth_photo_not_found](https://user-images.githubusercontent.com/37493831/233639743-c8099f14-a896-421f-bcb7-91c043420439.png)
![delete_photo_failed_unauth](https://user-images.githubusercontent.com/37493831/233639745-e3a965e8-bf19-4c35-9a66-2ac0a7982c60.png)

### `GET` photos
it's a get all photo endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_all_photo_success](https://user-images.githubusercontent.com/37493831/233639761-f0c41716-bb52-499d-bd32-b18ae5e6c326.png)

- failed response:
![get_all_photo_failed_auth](https://user-images.githubusercontent.com/37493831/233639758-8ef13763-63f9-4bb6-a329-870d79d53aa7.png)

### `GET` photos/me
it's a get all photo {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_all_photo_user_success](https://user-images.githubusercontent.com/37493831/233639771-d68e52c1-286e-49da-9381-9c1edb65aa5c.png)

- failed response:
![get_all_photo_user_failed_auth](https://user-images.githubusercontent.com/37493831/233639767-2ea776f3-f582-4741-8d93-cc108171c3b6.png)

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
![get_photo_success](https://user-images.githubusercontent.com/37493831/233639787-1de5ce7f-979d-457b-b2a9-9f930e17a0b7.png)

- failed response:
![get_photo_failed_auth](https://user-images.githubusercontent.com/37493831/233639778-c3f3628e-54c9-4103-b4c5-d2c073216cd1.png)
![get_photo_failed_not_found](https://user-images.githubusercontent.com/37493831/233639783-21d292ab-42de-4f95-b68a-e4acb945b5ec.png)

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
![create_comment_success](https://user-images.githubusercontent.com/37493831/233644065-4fbb3551-58f2-46aa-8fa5-670ec29e52d3.png)

- failed response:
![create_comment_failed_auth](https://user-images.githubusercontent.com/37493831/233644052-076962da-ebb2-4498-a194-70168eb1119b.png)
![create_comment_failed_not_found](https://user-images.githubusercontent.com/37493831/233644055-e7feafa2-96a9-49a5-b42e-dce6bedc69ce.png)

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
![update_comment_success](https://user-images.githubusercontent.com/37493831/233644047-f2dafc5d-cbfb-4b0b-959b-267258856da8.png)

- failed response:
![update_comment_failed_auth](https://user-images.githubusercontent.com/37493831/233644034-4e48f32e-cca9-47a8-834d-3d9c428223a1.png)
![update_comment_failed_not_found](https://user-images.githubusercontent.com/37493831/233644039-e5d22c93-8407-4b15-9647-36bdc390719c.png)
![update_comment_failed_unauth](https://user-images.githubusercontent.com/37493831/233644043-1c03cf81-64b7-43e1-92dc-d08bad64a20c.png)

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
![delete_comment_success](https://user-images.githubusercontent.com/37493831/233644078-dd98a3de-38ad-4d78-8c55-c341f02f579b.png)

- failed response:
![delete_comment_failed_auth](https://user-images.githubusercontent.com/37493831/233644071-8b8ab973-c563-4474-af21-ea3263449a07.png)
![delete_comment_failed_unauth](https://user-images.githubusercontent.com/37493831/233644074-a1ba94c3-bceb-41d1-a78a-5ca50237fa79.png)

### `GET` comments
it's a get all comment endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_all_comment_success](https://user-images.githubusercontent.com/37493831/233644001-5cc0aceb-0b9d-46a2-ab57-5f4365e85314.png)

- failed response:
![get_all_comment_failed_auth](https://user-images.githubusercontent.com/37493831/233643907-b0a47ab6-81c4-4431-b9b6-f6c5cfcafe50.png)

### `GET` comments/me
it's a get all comment {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_all_comment_user_success](https://user-images.githubusercontent.com/37493831/233644011-0f3e4b1f-5b77-4428-b98d-92d4a77dcd6a.png)

- failed response:
![get_all_comment_user_failed_auth](https://user-images.githubusercontent.com/37493831/233644006-cf3f2386-8443-49ca-b2a1-32ed4bfd3fc0.png)

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
![get_comment_success](https://user-images.githubusercontent.com/37493831/233644031-1cbd7b5e-0779-4778-9027-df95fc46c19f.png)

- failed response:
![get_comment_failed_auth](https://user-images.githubusercontent.com/37493831/233644018-9887ef00-e752-44d7-bd4b-1fc34735dd36.png)
![get_comment_failed_not_found](https://user-images.githubusercontent.com/37493831/233644023-f7c99b36-a24e-4943-97cf-5ddd934edf92.png)

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
![create_sosmed_success](https://user-images.githubusercontent.com/37493831/233647553-213c9d85-64c3-4d02-8cfc-fce42624b384.png)

- failed response:
![create_sosmed_failed_auth](https://user-images.githubusercontent.com/37493831/233647545-c9768115-3d2a-4106-8fa0-cdd41294f1ac.png)
![create_sosmed_failed_dup](https://user-images.githubusercontent.com/37493831/233647551-62e41c91-cd04-4227-a11b-fc9c7dfac762.png)

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
![update_sosmed_success](https://user-images.githubusercontent.com/37493831/233647541-3e052eec-8192-42b7-a922-6cc4bffccc83.png)

- failed response:
![update_sosmed_failed_auth](https://user-images.githubusercontent.com/37493831/233647536-814c84f4-d6a3-41b4-9526-e83616e4a0e2.png)
![update_sosmed_failed_unauth](https://user-images.githubusercontent.com/37493831/233647537-ee58f051-5333-451a-a7ce-a66e4b36f21a.png)

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
![delete_sosmed_success](https://user-images.githubusercontent.com/37493831/233647562-6bce671f-d5d2-4d40-aad4-0cf302aeafb0.png)


- failed response:
![delete_sosmed_failed_auth](https://user-images.githubusercontent.com/37493831/233647554-0258343c-8ef0-48e5-a91e-5a63962fbc3f.png)
![delete_sosmed_failed_unauth](https://user-images.githubusercontent.com/37493831/233647559-4f7bbde9-bd58-458a-af59-6b9cdea4ce0f.png)

### `GET` socialmedia
it's a get all social media endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_all_sosmed_success](https://user-images.githubusercontent.com/37493831/233647506-4a15be3f-46c6-4837-8bb0-9abc9c678e2c.png)

- failed response:
![get_all_sosmed_failed_auth](https://user-images.githubusercontent.com/37493831/233647565-9dfb4706-d763-4091-b50a-0ad72b741171.png)

### `GET` socialmedia/me
it's a get social media {by userId jwt} endpoint. This endpoint only can be accessed by `user` that logged in.

**Request Headers**
```bash
Authorization : Bearer <token>
```

**Response Body** :
- success response:
![get_sosmed_user_success](https://user-images.githubusercontent.com/37493831/233647531-1effc67e-58ef-4616-ace9-00cb017c9da0.png)

- failed response:
![get_sosmed_user_failed_auth](https://user-images.githubusercontent.com/37493831/233647527-10212045-204c-4c89-91a4-81e9fed082eb.png)

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
![get_sosmed_success](https://user-images.githubusercontent.com/37493831/233647522-e2e7cb92-b8c9-419a-8927-39e8b01eec49.png)

- failed response:
![get_sosmed_failed_auth](https://user-images.githubusercontent.com/37493831/233647517-a72868f3-6ea5-4a9e-91b4-8461bc97c4cd.png)