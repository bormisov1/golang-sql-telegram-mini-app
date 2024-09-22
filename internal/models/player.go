package models

type Player struct {
    ID        string
    Position  float64 // Player's current position
    Direction bool    // Player's current direction (true for +1, false for -1)
    Speed     float64 // Player's speed
}
