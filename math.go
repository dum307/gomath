package gomath

// Add теперь складывает ТРИ числа (изменено в v2!)
func Add(a, b, c int) int {
    return a + b + c
}

// Multiply умножает два числа  
func Multiply(a, b int) int {
    return a * b
}

// Sub вычитает b из a
func Sub(a, b int) int {
    return a - b
}

// Divide делит a на b
func Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}