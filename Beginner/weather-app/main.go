package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

const apiKey = "YOUR API KEY" // Replace with your OpenWeatherMap API key

type WeatherResponse struct {
    Name string `json:"name"`
    Main struct {
        Temp     float64 `json:"temp"`
        Pressure int     `json:"pressure"`
        Humidity int     `json:"humidity"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func getWeather(city string) {
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("❗ Error fetching data:", err)
        return
    }
    defer response.Body.Close()

    if response.StatusCode != 200 {
        fmt.Println("❗ Error: Invalid response from server.")
        return
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("❗ Error reading response:", err)
        return
    }

    var weatherData WeatherResponse
    err = json.Unmarshal(body, &weatherData)
    if err != nil {
        fmt.Println("❗ Error parsing JSON:", err)
        return
    }

    fmt.Printf("📍 Weather in %s:\n", weatherData.Name)
    fmt.Printf("🌡️ Temperature: %.2f°C\n", weatherData.Main.Temp)
    fmt.Printf("💧 Humidity: %d%%\n", weatherData.Main.Humidity)
    fmt.Printf("🌬️ Pressure: %d hPa\n", weatherData.Main.Pressure)
    fmt.Printf("🌥️ Condition: %s\n", weatherData.Weather[0].Description)
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("❗ Please provide a city name.")
        return
    }

    city := os.Args[1]
    getWeather(city)
}
