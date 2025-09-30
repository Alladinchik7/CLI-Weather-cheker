# CLI Weather cheker

## 📋 Описание проекта

Основанный на Go инструмент получения данных о погоде через API.

## 🛠 Основные функции

- Использует API для получения данных о погоде
- Выводит данные о погоде ведённого города
- Выводит данные ввиде таблички

## 🚀 Установка и запуск

### Требования

- Go 1.22 или выше

### Установка

1. Клонируйте репозиторий

    ```markdown
        
        ```bash
        `git clone https://github.com/Alladinchik7/CLIWeather.git`
        ```

2. Перейдите в директорию проекта

    ```markdown

        ```bash
        `cd CLI-weather-cheker`
        ```

3. Настройте конфиги в .env орентируясь по .env.example(Используйте ключ от `https://www.visualcrossing.com/weather-api/`)

4. Запустите программу

    ```markdown

        ```bash
        `.\weather.exe`
        ```

## 🎮 Примеры использования

1. Получение данных о погоде в Москве

    ```markdown

        ```bash
        .\weather.exe Moscow,ru
        ```

        ```bash
        Weather for  Temperature  Humidity
        -----------  -----------  --------
        Moscow,ru    7            63.2
        ```
