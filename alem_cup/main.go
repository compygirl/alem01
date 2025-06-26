package main

import (
    "fmt"
    "os"
)

func main() {
    for {
        var n, m, playerID, tick int
        fmt.Scan(&n, &playerID, &tick)

        // Список городов
        cities := make([]City, n)

        for i := 0; i < n; i++ {
            var owner string
            var units, x, y int
            var cityName string
            fmt.Scan(&owner, &units, &cityName, &x, &y)
            cities[i] = City{
                Owner:    owner,
                Units:    units,
                Name:     cityName,
                X:        x,
                Y:        y,
            }
        }

        // Число передвижений
        fmt.Scan(&m)

        // Обработка передвижений
        movements := make([]Movement, m)
        for i := 0; i < m; i++ {
            var fromCity, toCity, fromPlayer, toPlayer string
            var turns, units int
            fmt.Scan(&fromCity, &toCity, &fromPlayer, &toPlayer, &turns, &units)
            movements[i] = Movement{
                FromCity:  fromCity,
                ToCity:    toCity,
                FromPlayer: fromPlayer,
                ToPlayer:  toPlayer,
                Turns:     turns,
                Units:     units,
            }
        }

        // Ваша логика здесь
        action := TakeAction(cities, playerID, tick)
        
        // Вывод действия
        fmt.Println(action)
    }
}

// Структура для представления города
type City struct {
    Owner    string
    Units    int
    Name     string
    X, Y     int
}

// Структура для представления передвижения
type Movement struct {
    FromCity  string
    ToCity    string
    FromPlayer string
    ToPlayer  string
    Turns     int
    Units     int
}

func TakeAction(cities []City, playerID, tick int) string {
    // Пример логики:
    // Перебираем города, которые принадлежат игроку (playerID)
    for _, city := range cities {
        if city.Owner == fmt.Sprintf("p%d", playerID) {
            // Если у города достаточно юнитов и есть соседи с вражескими городами,
            // усилим его, добавив половину юнитов к его силе.
            if city.Units > 50 {
                return fmt.Sprintf("STRENGTHEN %s", city.Name)
            }
            // В противном случае, попробуем отправить часть юнитов в соседний вражеский город.
            else {
                for _, neighbor := range cities {
                    if neighbor.Owner != fmt.Sprintf("p%d", playerID) && areNeighbors(city, neighbor) {
                        return fmt.Sprintf("MOVE %s %s %d", city.Name, neighbor.Name, city.Units/2)
                    }
                }
            }
        }
    }

    // Если нет подходящего действия, просто ждем и ничего не делаем.
    return "WAIT"
}

// Функция, определяющая, являются ли два города соседями (находятся рядом друг с другом)
func areNeighbors(city1, city2 City) bool {
    // Здесь вы можете определить, какие условия делают два города соседями.
    // Например, если они находятся на расстоянии не более чем 1 клетка друг от друга.
    return abs(city1.X-city2.X) <= 1 && abs(city1.Y-city2.Y) <= 1
}

// Функция для получения абсолютного значения числа
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

