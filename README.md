# 🧮 Order Pack Calculator

This is a web application and HTTP API written in **Golang** that calculates the optimal combination of product pack sizes to fulfill customer orders.

It respects the following business rules:

1. ✅ Only **whole packs** can be sent (no breaking packs).
2. ✅ Send the **least number of items** to fulfill the order (never send less).
3. ✅ Among solutions with the same total items, send **as few packs as possible**.
   <br></br>
## 🚀 Live Demo
You can visit below link to try out the calculator app:
👉 [live demo app](https://immense-mountain-86814-9a015e36c702.herokuapp.com/)
<br></br>

## 🎯 Example

If available pack sizes are `250`, `500`, `1000`, `2000`, `5000` and the customer orders `12001` items → the app returns:

```
Total Items: 12250
Packs:
2 x 5000
1 x 2000
1 x 250
```


✅ Minimal total items  
✅ Minimal packs for that total
<br></br>
## 🏗️ Features

- 🌐 Web-based user interface (HTML + JavaScript)
- 📡 HTTP API with JSON responses
- ⚙️ Pack sizes are **configurable dynamically via UI**
- ➕ Add pack sizes from UI
- ➖ Remove individual pack sizes
- 🧹 Remove **all pack sizes** with one click
- 📝 API returns optimal pack combination
- ✅ Fully deployable on **Heroku**
  <br></br>
## 📦 Available API Endpoints

| Method | Path                    | Description                            |
|--------|------------------------|---------------------------------------|
| GET    | `/api/packsizes`        | Get current pack sizes                 |
| POST   | `/api/packsizes/add`     | Add a new pack size `{ "size": N }`   |
| DELETE | `/api/packsizes/delete`  | Remove a pack size `?size=N`           |
| DELETE | `/api/packsizes/deleteall` | Remove **all** pack sizes            |
| GET    | `/api/calculate?quantity=N` | Calculate optimal packs for quantity |
<br></br>

## 🖥️ Running Locally

1. **Install Go (1.20+)**

2. Clone the repo:

```bash
git clone https://github.com/yourusername/my-order-app.git
cd my-order-app
```

3. Build the app:

```
go build -o my-order-app
```

4. Run the app

```
./my-order-app
```

5. Open the browser

```
http://localhost:8080
```

## 📝 Project Structure
```bash
my-order-app/
├── main.go
├── go.mod
├── config.json
├── Procfile
├── internal/
│   ├── config/
│   ├── calculator/
│   └── handler/
└── static/
├── index.html
└── app.js
```
<br></br>

## 🧑‍💻 Configuration
Initial pack sizes are stored in config.json, but pack sizes can be modified dynamically from the UI or API without restarting the app.

Example config.json:

```json
{
"pack_sizes": [250, 500, 1000, 2000, 5000]
}
```