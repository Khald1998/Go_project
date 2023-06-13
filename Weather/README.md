# README.md

## 1. Title
Weather Search App

## 2. Background
This application is a simple web-based program for obtaining weather data based on the name of a city. The program is built using Go language and uses the [Gin](https://github.com/gin-gonic/gin) web framework. It uses the [OpenWeatherMap API](https://openweathermap.org/api) to retrieve weather data. It also utilizes the [go-randomdata](https://github.com/Pallinder/go-randomdata) package to generate random city names.

## 3. Prerequisites
Before you can run this application, you must have the following prerequisites:

- [Go](https://golang.org/) programming language installed on your machine.
- The following Go packages:
  - Gin web framework: Install by running `go get -u github.com/gin-gonic/gin`
  - go-randomdata: Install by running `go get -u github.com/Pallinder/go-randomdata`
  - gjson: Install by running `go get -u github.com/tidwall/gjson`
  - gin-contrib/static: Install by running `go get -u github.com/gin-contrib/static`
- API Key for OpenWeatherMap: Replace the `apiKey` constant in the code with your OpenWeatherMap API key.

## 4. Inputs
The application accepts city name as an input. This input can be given in two ways:
1. Search: Manually enter the city name in the search box.
2. Random: If the Random button is pressed, the application will generate a random city name.

## 5. Outputs
The application will display weather details of the input city. The output will include:
- City name
- Current Temperature
- Feels like Temperature
- Maximum Temperature
- Minimum Temperature
- Weather (e.g., Clear, Rain, etc.)
- Weather Description
- Weather Icon

If the city name cannot be found or if there is any error with the input, an error page will be displayed with the corresponding error message.

## 6. Steps to Run
1. Clone the repository to your local machine or download the `main.go` file.
2. Open your terminal and navigate to the folder where you cloned the repo or where the `main.go` file is located.
3. Run the application with the command `go run main.go`.
4. Open a web browser and navigate to `http://localhost:8080`.
5. You can now start searching for weather data by city name or use the random feature to get weather data for a randomly selected city.
