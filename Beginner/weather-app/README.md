# 🌤️ Weather Application

A simple command-line tool to fetch and display current weather data for any city using the OpenWeatherMap API.

## 🔑 Prerequisites

- Install Go: [Download Go](https://golang.org/dl/)  
- Get a free API key from [OpenWeatherMap](https://openweathermap.org/)

## 🚀 How to Run

1. **Navigate to the project directory:**

   ```bash
   cd Beginner/weather-app

2. **Replace YOUR_API_KEY in main.go with your actual OpenWeatherMap API key.**

3. **Run the application**
   ```bash
   go run main.go Mumbai

4. **Example Output:**

   ```yaml
   📍 Weather in Mumbai:
   🌡️ Temperature: 29.34°C
   💧 Humidity: 74%
   🌬️ Pressure: 1012 hPa
   🌥️ Condition: scattered clouds

 **Customization**
Change the temperature unit by modifying the units parameter in the URL (metric for °C, imperial for °F).

Enjoy exploring the weather! ☀️🌧️