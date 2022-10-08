## 接続先
```35.208.4.136:8080```

## API仕様
### POST /user/create
新規ユーザを作成する

#### Request
| name | type | memo |
|---|---|---|
| username | string | |
| password | string | |

#### Response
| name | type | memo |
|---|---|---|
| token | string | ユーザがログインした状態であることを認証するための値 |

### POST /login
ログインする(すでにユーザを作成している場合)

#### Request
| name | type | memo |
|---|---|---|
| username | string | |
| password | string | |

#### Response
| name | type | memo |
|---|---|---|
| token | string | ユーザがログインした状態であることを認証するための値 |

### POST /drivelog
違反時のドライブデータを保存する

#### Request
| name | type | memo |
|---|---|---|
| token | string | ログインしたときのtoken |
| date | string | 現在時刻 |
| speed | float | 速度 |
| acceleration | float | 加速度 |
| latitude | float | 緯度 |
| longtitude | float | 経度 |

#### Response
| name | type | memo |
|---|---|---|

### GET /drivelog
違反時のドライブデータを取得する

#### Request
| name | type | memo |
|---|---|---|
| token | string | ログインしたときのtoken |
| date | string | 現在時刻 |

#### Response
| name | type | memo |
|---|---|---|
| speed | float | 速度 |
| acceleration | float | 加速度 |
| latitude | float | 緯度 |
| longtitude | float | 経度 |


### 使用例
#### ユーザ作成
```
curl -d {username=taro, password=pass} http://127.0.0.1:8080/user/create
```
```
{
    "token": "8407d217-46fe-11ed-907b-84a938147d00"
}
```

