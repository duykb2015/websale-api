# API user

## API regiser
*Create user login system ...*


### Defined flow code

```JSON
I. Register
    1. Nhan thong tin tu body json
    2. Kiem tra input hop hay khong
        2.1. Khong hop le -> bao loi
            2.1.1: Kiem tra email hop le
            2.1.2: Kiem tra email ton tai trong he thong hay chua
            2.1.3: Kiem tra password hop le
            2.1.3: Kiem tra password == password_confirm
        2.2. Hop le -> Tiep tuc
    3. Goi function insert thong tin user
        3.1. Tao ket noi den database
        3.2. Tao cau query insert database
        3.3. Tra thong tin loi
    4. Kiem tra user insert
        4.1. That bai -> bao loi
        4.2. Thanh cong -> tiep tuc
    5. Goi function tao token
        5.1. That bai -> bao loi
        5.2 Thanh cong -> tra ve token
```


### Infomation

| Information | Description                         |
| ----------- | ----------------------------------- |
| Endpoint    | `https://<domain>/user/v1/register` |
| Method      | POST                                |

### Request Header


| Name         | Value              |
| ------------ | ------------------ |
| Content-Type | `application/json` |


### Payload

| Parameter    | Position     | Required | Nullable | Description                               |
| ------------ | ------------ | -------- | -------- | ----------------------------------------- |
| `authorized` | JSON Payload | Yes      | No       | To check if the user is authorized or not |
| `id_user`    | JSON Payload | Yes      | No       | The id of user                            |
| `exp`        | JSON Payload | Yes      | No       | Token life time                           |

#### Example Payload

```JSON
{
    "authorized": true,
    "id_user":1,
    "exp":"15 minute"
}
```

### Response

| Parameter | Position     | Required | Nullable | Description |
| --------- | ------------ | -------- | -------- | ----------- |
| `token`   | JSON Payload | Yes      | No       | Token ....  |


#### Example Response

```JSON
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjI5NTQ5NjcsInVzZXJfaWQiOjB9.imsyPkh4fLK-xe9JSW7hsCOS7S06wmpzTV0w4adcdqk"
}
```
### Prepare

```



```