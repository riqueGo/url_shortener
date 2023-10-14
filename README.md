# url_shortener

Its a project from Devgym.

Create an http server that contains two endpoints:
- POST / - takes a url and returns a unique code
- GET /:code - uses the code to redirect to the original url
- The code is a single code, the same url sent several times generates different codes
- The code is up to 6 characters long

## Getting Started

1. Clone this repository:

   ```
   git clone https://github.com/yourusername/your-project.git
   cd your-project
   ```

2. Docker up:

   ```
   docker compose up
   ```

### API Endpoints

- `POST`
```
curl -X POST -d '{"url": "https://google.com"}' -H "Content-Type: application/json" http://localhost:8080/
```

- `GET`
```
curl http://localhost:8080/{code}
```
