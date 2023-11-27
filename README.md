# ush - Link Shortener 🔪

```bash
git clone https://github.com/sotchenkov/ush.git
go mod download
go run cmd/ush/main.go 
```

```
[POST] http://host:port/url - to save url

body: {
    "url": "https://google.com",
    "alias": "gg"
}

✅ Basic Auth
```

```
[GET] http://host:port/{alias} - Redirect by alias
```
[reference](https://github.com/GolangLessons/url-shortener)

