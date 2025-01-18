# 🔗 URL Shortener

A simple URL shortener built with Go and the Gin framework.

## 🚀 How to Run

1. **Install dependencies:**

   ```bash
   cd Intermediate/url-shortener
   go mod init url-shortener
   go get github.com/gin-gonic/gin

2. **Run the server**
   ```bash
   go run main.go

3. **Usage**
Shorten a URL:
Send a POST request to http://localhost:8080/shorten with JSON data:

   ```json
   {
     "url": "https://example.com"
   }
Response
   ```json
   {
     "short_url": "http://localhost:8080/abc123"
   }
Access the shortened URL:
Open the generated short URL in your browser, and it will redirect to the original URL.