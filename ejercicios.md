Aquí tienes una lista de ejercicios progresivos para que tus compañeros practiquen lógica de programación en Go, comenzando desde lo más básico hasta ejercicios un poco más complejos:

### 1. **Hola Mundo Personalizado**
**Objetivo:** Familiarizarse con la estructura básica y entrada/salida
```go
// Escribe un programa que pida el nombre al usuario y muestre:
// "Hola, [nombre]! Bienvenido a Go."
// Usa fmt.Scanln() para la entrada
```

### 2. **Calculadora Simple**
**Objetivo:** Operadores básicos y manejo de variables
```go
// Crea una calculadora que pida dos números y muestre:
// - Suma
// - Resta
// - Multiplicación
// - División (con 2 decimales)
```

### 3. **Clasificador de Números**
**Objetivo:** Uso de condicionales
```go
// Pide un número y determina si es:
// - Positivo/Negativo/Cero
// - Par/Impar
// - Múltiplo de 5
```

### 4. **Generador de Secuencias**
**Objetivo:** Bucles básicos
```go
// Pide un número n y muestra:
// a) Todos los números del 1 al n
// b) Los primeros n números pares
// c) La secuencia 2, 4, 8, 16... hasta superar n
```

### 5. **Buscador en Array**
**Objetivo:** Manejo de slices y búsqueda lineal
```go
// Crea un programa que:
// 1. Genere un slice con 10 números aleatorios (1-100)
// 2. Pida un número al usuario
// 3. Diga si está en el slice y en qué posición(es)
```

### 6. **Conversor de Temperatura**
**Objetivo:** Funciones y conversión de unidades
```go
// Crea dos funciones:
// - CelsiusToFahrenheit(c float64) float64
// - FahrenheitToCelsius(f float64) float64
// Programa principal que permita elegir conversión y muestre resultado
```

### 7. **Validador de Contraseña**
**Objetivo:** Manipulación de strings y condiciones compuestas
```go
// Valida que una contraseña tenga:
// - Mínimo 8 caracteres
// - Al menos 1 mayúscula
// - Al menos 1 número
// - Al menos 1 carácter especial (!@#$%^&*)
```

### 8. **Tablero de Ajedrez**
**Objetivo:** Bucles anidados y lógica de patrones
```go
// Imprime un tablero de ajedrez 8x8 usando:
// 'B' para casillas blancas
// 'N' para casillas negras
// Bonus: Usar runas/emojis para representar piezas
```

### 9. **Calculadora de Promedio con Slices**
**Objetivo:** Manejo de slices dinámicos
```go
// Programa que:
// 1. Permita ingresar números hasta que se escriba "fin"
// 2. Calcule y muestre:
//    - Promedio
//    - Número mayor
//    - Número menor
```

### 10. **Juego: Adivina el Número**
**Objetivo:** Bucles y números aleatorios
```go
// Implementa un juego donde:
// 1. La computadora genera un número (1-100)
// 2. El usuario tiene 5 intentos para adivinarlo
// 3. En cada intento dar pistas "muy alto" o "muy bajo"
```

### 11. **Manipulador de Texto**
**Objetivo:** Trabajar con strings y funciones
```go
// Crea un programa con funciones para:
// - Invertir un string (ej: "hola" -> "aloh")
// - Contar vocales
// - Verificar si es palíndromo
// Menú principal para seleccionar operación
```

### 12. **Gestor de Contactos Simple**
**Objetivo:** Structs y slices de structs
```go
// Define un struct Contacto con:
// Nombre, Teléfono, Email
// Programa que permita:
// - Agregar contactos
// - Buscar por nombre
// - Mostrar todos ordenados alfabéticamente
```

### Consejos para los ejercicios:
1. **Testeo Simple:** Enséñales a usar `fmt.Println()` para depuración
2. **Manejo de Errores Básico:** Introducir `if err != nil` en inputs
3. **Refactorización:** Primero que hagan funcionar el código, luego optimicen
4. **Comparación con TypeScript:** Resaltar diferencias clave como tipado fuerte y manejo de errores

### Ejercicio Bonus (Para los más avanzados):
```go
// Implementa un programa concurrente que:
// 1. Haga ping a 3 URLs simultáneamente
// 2. Muestre "OK" o "Fallo" según respuesta
// Usar goroutines y channels
```

¿Quieres que profundice en algún ejercicio en particular o prefieres que agregue más ejercicios de algún área específica?