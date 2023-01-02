# Go desde 0
_<link> Curso: https://www.udemy.com/course/lenguaje-go/_
---

## Sección 01 - Introducción

#### Características de Go (Golang)
- Es un lenguaje de programación creado por Google.

- Fue creado para resolver problemas internos de Google y luego escaló al universo del desarrollo.

- Es un lenguaje fuertemente tipado.

- Es un lenguaje que hereda su core de C++ aunque su sintaxis es mejorada y muy amigable.

- Fue pensado para aprovechas los últimos avances en hardware, los multiplocesadores, enormes cantidades de memoria y el paralelismo.

- Es un lenguaje compilado, generando archivos `*.exe` portables a cualquier sistema operativo y versión (su `*.exe` contiene todo lo necesario para ejecutarse)

- Obliga a los desarrolladores a llevar buenas prácticas.

- Demostró ser el lenguaje ideal para grandes desarrollos con miles y miles de usuarios concurrentes y millones de transacciones.

#### Convensiones de Go (Golang)
- Es fácil de entender. Su sintaxis es clara y mejorada con respecto a otros lenguajes.

- No es necesario usar `;` (punto y coma).

- El compilador arroja advertencias, ante variables declaradas no utilizadas y paquetes importados no utilizados.

- Las funciones pueden devolver más de un valor.

- Se pueden desarrollar tanto instrucciones sincrónicas como asíncronicas.

- Sólo existe la instrucción `for{...}` para iteraciones.

- No es un lenguaje orientado a objetos y resuelve lo que conocemos como POO, con estructuras, funciones, métodos e interfaces.

- El scope de variables, métodos y funciones se determinan con nombres en mayúsculas y minúsculas. Por ejemplo, si el nombre de una función comienza con mayúscula, el `scope` de esa función es `public`, permitiendo así que se exporte a otros `pkg`. En cambio, si el nombre comienza con minúscula, estamos en presencia de una función `private`.

---

## Sección 02 - Creando el entorno de trabajo
Links útiles:
- _<link> Go (Golang): https://golang.org/dl/_
- _<link> VS Code (IDE): https://code.visualstudio.com/docs/?dv=win_

---

## Sección 03 - Hola Mundo en Go

- Al momento de crear un nuevo archivo de Go, por convensión el nombre del archivo debe ser igual al nombre del `package`. Esto nos facilita la importación/exportación de funcionalidades.

- Todo el inicio de ejecución de un programa en Go comenzará o debe comenzar por el `package main`.

- Todo en Go son funciones (`func`).

- Sintaxis general de una función:
```go
func nameFunction(nameParameters ...typeParameters) (nameReturns ...typeReturns) {
    //TODO: add code for function
}
```

- Comando para correr el programa:
```go
go run main.go
```

- Comando para crear el `*.exe` para correr el programa en otro dispositivo. Es importante recordar que este archivo va a contener todo lo necesario para que se ejecute en cualquier sitio sin generar ninguna dependencia.
```go
go build main.go
```

---

## Sección 04 - Convenciones y sintaxis de Go

#### Variables

- Las variables se declaran de la siguiente forma:
```go
var variableName dataType

variableName := variableValue
```

- Se permite encadenar variables o declaración simultanea de variables sin necesidad que sean del mismo tipo de variables:
```go
var variableOne, variableTwo, variableThree, variableFour, variableFive dataType

variableOne, variableTwo, variableThree, variableFour, variableFive := valueOne, valueTwo, valueThree, valueFour, valueFive
```

- Toda variable declarada, aún sin haber sido asignada, se inicializa en su valor `zero`:
    - Las variables de tipo `string` se inicializan en `""`.
    - Las variables de tipo `int, float, uint` se inicializan en `0`.
    - Las variables de tipo `bool` se inicializan en `false`.

- No se permite cambiar el tipo de datos de variables de manera explícita. En este caso se debe `parsear` al tipo que necesitamos, por ejemplo, entre otras opciones:
```go
// método 1
var numero int = 2
var moneda int64 = 100
numero = int(moneda)

// método 2
texto = strconv.Itoa(int(moneda))
```

- Para que una variable sea accedida desde otro `package` el nombre de ésta debe comenzar con letra mayúscula.

- En Go existen 3 (tres) tipos de variables:
    - Numéricas:
        - Enteras
        - Enteras sin signo
        - Con coma flotante
    - Cadena:
    - Booleanas:


#### Condicionales

##### Conficional `if`

```go
package main

import (
    "fmt"
)

func main() {
    estado = true
    if estado == true {
        fmt.Println(estado)
    } else {
        fmt.Println("Es falso")
    }
}
```

- Particularmente en `go`, en medio de un condicional se permite asignar un valor a una variable afectada a la condición:
```go
estado = true

if estado = false; estado == true {
    fmt.Println(estado)
} else {
    fmt.Println("Es falso")
}
``` 

##### Conficional `switch`

```go
switch numero := 5; numero {
    case 1:
        fmt.Println(1)
    case 2:
        fmt.Println(2)
    case 3:
        fmt.Println(3)
    case 4:
        fmt.Println(4)
    case 5:
        fmt.Println(5)
    default:
        fmt.Println("Mayor a 5")
}
```

#### Mostrar y aceptar datos por teclado

```go
package main

import (
    "fmt"
    "os"
    "bufio"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
    fmt.Println("Ingrese número 1:")
    //fmt.Scanf("%d", &numero1) >> esta función tiene un error para Windows, ya que toma el 'enter' como un segundo valor
    fmt.Scanln(&numero1)

    fmt.Println("Ingrese número 2:")
    //fmt.Scanf("%d", &numero2) >> esta función tiene un error para Windows, ya que toma el 'enter' como un segundo valor
    fmt.Scanln(&numero2)

    fmt.Println("Descripción:")

    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        leyenda = scanner.Text()
    }

    resultado = numero1 + numero2
    fmt.Println(leyenda, resultado)
}
```

#### Ciclos

##### Ciclo `for`

```go
package main

import (
    "fmt"
)

func main() {
    //forma 1 de implementación
    i := 1
    for i < 10 {
        fmt.Println(i)
        i++
    }

    // forma 2 de implementación 
    for i := 1; i < 10; i++ {
        fmt.Println(i)
    }


    // forma 3 de implementación
    var numero int
    for {
        fmt.Println("Continuo")
        fmt.Println("Ingrese el número secreto:")
        fmt.Scanln(&numero)
        if numero == 100 {
            break
        }
    }

    // forma 4 de implementación
    var i = 0
    for i < 10 {
        fmt.Printf("\n Valor de i: %d", i)
        if i == 5 {
            fmt.Printf(" multiplicamos por 2 \n")
            i = i*2
            continue
        }
        fmt.Printf("             Pasó por aquí \n")
        i++
    }
}
```
- Notemos que existen 2 palabras reservedas:
    - `break`: utilizada para romper el ciclo, es decir directamente salir de donde estamos.
    - `continue`: utilizada para volver a evaluar inmediatamente la condición del ciclo y proseguir según corresponda.

- En Go se pueden asignar etiquetas de secciones, a modo de indicar el 'inicio de código nuevo'. Esto se puede utilizar junto con la sentencia `goto` para indicarle a qué linea de código queremos ir. Para que esto sea posible, debemos asignarle una sección a esta línea dónde queremos ir. Por ejemplo, a continuación se crea la sección _RUTINA_
```go
package main

import (
    "fmt"
)

func main() {
    var i int = 0

    RUTINA:
        for i < 10 {
            if i == 4 {
                i = i + 2
                fmt.Println("voy a RUTINA sumando 2 a i")
                goto RUTINA
            }
            fmt.Printf("Valor de i: %d\n", i)
            i++
        }
}
```

#### Funciones

- En Go es preciso recordar que no existen los métodos. Pero bien son tipos de funciones que no devuelven ningún valor o tipo de dato.
```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(uno(5))
    fmt.Println(unoBis(5))

    numero, estado := dos(1)
    fmt.Println(numero)
    fmt.Println(estado)
}

func uno(numero int) int {
    return numero*2
}

func unoBis(numero int) (resultado int) {
    resultado = numero*2 
    return
}

func dos(numero int) (int, bool) {
    if numero == 1 {
        return 5, true
    } else {
        return 10, false
    }
}
```

- En Go se hace presente una falencia y es que no existe la sobrecarga de métodos/funciones. Por ello, se implementa las funciones con `parámetros variables`:
```go
func calculo(numero ...int) int {
    total := 0
    for _, valor := range numero {
        total += valor
    }
    fmt.Printf("\n Total acumulado: %d", total)
    return total
}
```
- La instrucción `range` siempre debe utilizarse con listas, arreglos, mapas, colecciones, etc. Devuelve siempre 2 valores: el índice y el elemento contenido en la colección.

- En Go para no manipular una variable por no interesarnos se la nombra simplemente `_`. Así, se permite que el lenguaje guarde esa variable pero no reserve espacio en memoria, esto significa que reconoce el valor correspondiente pero lo elimina instantáneamente.

#### Funciones anónimas y closures

- `Funciones anónimas`: son funciones que no tienen nombre. Se utilizan para crear operaciones en las que no necesitamos llamar a una función, sino que vamos a isolar todo su código dentro de una variable y al que podemos mofificar en tiempo de ejecución.
```go
package main

import (
    "fmt"
)

var Calculo func(int, int) int = func(numero1 int, numero2 int) int {
    return numero1 + numero2
}

func main() {
    fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5,7))

    // redefinimos la devolución de la función
    Calculo = func(num1 int, num2 int) int {
        return num1 - num2
    }

    // redefinimos la devolución de la función
    Calculo = func(num1 int, num2 int) int {
        return num1 / num2
    }
}

func Operaciones() {
    resultado := func() int {
        a := 23
        b := 14
        return a + b
    }
    fmt.Println(resultado())
}
```

- `Closures`: es un concepto de funciones anónimas que se están utilizando mucho. Están relacionados con la isolación y protección de código. Pueden acceder a variables creadas fuera de la función, es decir en la función original que los llama.
```go
package main

import (
    "fmt"
)

func main() {
    /* Closures */
    tablaDel := 2
    MiTabla := Tabla(tablaDel)
    for i:= 1; i < 11; i++ {
        fmt.Println(MiTabla)
    }
}

func Tabla(valor int) func() int {
    numero := valor
    secuencia := 0

    return func() int {
        secuencia++
        return numero * secuencia
    }
}
```

#### Arreglos estáticos y slices

- `Vectores` o `arrays`: es una colección de datos en memoria.
```go
package main

import (
    "fmt"
)

var tabla1 [10]int

func main() {
    // forma 1 para declarar arrays
    tabla1[0] = 1
    tabla1[5] = 15

    fmt.Println(tabla1)

    // forma 2 para declarar arrays
    tabla2 := [10]int{1,2,3,4,5,6,7,8,9,10}

    for i := 0; i < len(tabla2); i++ {
        fmt.Println(tabla2[i])
    }

    // forma para declarar arrays multidimensionales (matriz)
    matriz := [10][7]int
    matriz[3][5] = 1
    fmt.Println(matriz)
}
```

- `Slices`: son `vectores` dinámicos, es decir que puedo variar las dimensiones en tiempo real o tiempo de ejecución.
```go
package main

import (
    "fmt"
)

func main() {
    // forma 1 de declarar un slice
    slice := []int{2,3,4}
    fmt.Prinln(slice)

    // forma 2 de declarar un slice
    elementos := [5]int{1,2,3,4,5}
    porcion1 := elementos[3:]
    porcion2 := elementos[:4]
    porcion3 := elementos[2:4]

    // forma 3 de declar un slice
    otroSlice := make([]int, 5, 20) // []int es el tipo de dato; 5 es la cantidad de elementos que va a tener; 20 es la capacidad máxima de guardar elementos
    fmt.Printf("Largo %d, Capacidad %d", len(otroSlice), cap(otroSlice))

    // forma 4 de declar un slice
    numeros := make([]int, 0, 0)
    for i := 0; i < 100; i++ {
        numeros = append(numeros, i) // a cada 'append' reasigna nuevamente el slice y lo redimensiona
    }
    fmt.Printf("\n Largo %d, Capacidad %d", len(numeros), cap(numeros))
```

#### Mapas

- `Map`: es una estructura que se puede serializar, es decir, crear elementos en serie. A diferencia de un `array`, donde el índice es numérico, en un `map` podemos especificar un par _[key : value]_ donde la _key_ puede ser de tipo `string`, `int` o `type`.
```go
package main

import (
    "fmt"
)

func main() {
    // forma 1
    paises := make(map[string]string)
    fmt.Println(paises)

    paises["Mexico"] = "D.F."
    paises["Argentina"] = "Buenos Aires"

    fmt.Println(paises)

    // forma 2
    seleccionesFutbol := make(map[string]string, 5)

    // forma 3
    equiposFutbol := map[string]int{
        "Barcelona": 39,
        "Real Madrid": 42,
        "Chivas": 37,
        "Boca Juniors": 30,
    }

    delete(equiposFutbol, "Chivas")

    // buscar en un map -> retorna 2 valores: el resultado obtenido, si existe, sino devuelve el valor 'zero' correspondiente; y el booleano true/false según encontró la 'key' especificada
    valor, existe := map["Chivas"]
}
```

#### Estructuras (POO)

- En Go no existe la programación orientada a objetos tal y como la hemos conocido con los conceptos de herencia, polimorfismo, clases, métodos, propiedades.
- Go implementa una nueva forma basada en `struct` (estructuras).

```go
package main

import (
    "fmt"
    "time"
)

type usuario struct {
    Id          int
    Nombre      string
    FechaAlta   time.Time
    Status      bool
}

func main() {
    User := new(Usuario)
    User.Id = 10
    User.Nombre = "Marcos"
    fmt.Println(User)
}
```

- _Ver código completo en "section\_04/estructuras"_

#### Interfaces

- Las `interfaces` definen conductas, operaciones y comportamientos.
- Una misma `struct` puede pertenecer a varias `interfaces` simpre que implemente todos los métodos de cada `interface`
- 

```go
```

- _Ver código completo en "section\_04/interfaces"_

#### Manejo de archivos

```go
```

- _Ver código completo en "section\_04/archivos"_

#### Recursión

- La `recursión` se usa más que nada en cálculos matemáticos. Trata de una función que se llama a sí misma repetidas veces una cantidad limitada de ellas.

```go
package main

import (
    "fmt"
)

func main() {

}

func exponencia(numer int) {
    if numero > 100000 {
        return
    }
    fmt.Println(numero)
    exponencia(numero * 2)
}
```

#### Defer, panic y recover

- Estos conceptos implican manejar `exceptions` y `errors`.
- La palabra reservada `defer` indica que la instrucción que la acompaña se ejecutará sí o sí pero al finalizar todas las ejecuciones de la función que contiene a la palabra `defer` mencionada. Es importante destacar que un `defer` ejecuta una sola instrucción; por lo que en caso de necesitar ejecutar varias, deberíamos implementar una función anónima.
- La palabra reservada `panic` interrumpe la ejecución en curso y se realiza el aborto del programa. Para poder controlarlo, se utiliza la función `recover`.
- La función `recover` retorna el último `panic` arrojado. Si no existió un `panic`, `recover` devolverá un `nil`.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)
	defer f.Close()

	if err != nil {
		fmt.Println("Error opening")
		os.Exit(1)
	}

    ejemploPanic()
}

func ejemploPanic() {
    defer func() {
        reco := recover()
        if reco == nil {
            log.Fatalf("ocurrió un error que generó Panic \n %v", reco)
        }
    }()

    a := 1
    if a == 1 {
        panic("se encontró el valor de 1")
    }
}

```

#### GoRoutines

- Estamos hablando de `programación asíncrona` o `hyperthreading`.
- Las `GoRoutines` son el equivalente a las `promesas` y a los `asyncs awaits`.
- Para indicar que una función va a ejecutarse de manera asíncrona, sólo debemos anteponer la palabra reservada `go`.

```go
package main

import (
    "fmt"
    "strings"
    "time"
)

func main() {
    go miNombreLento("marcos")
    fmt.Println("Estoy aquí")
    var x string
    fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
    letras := strings.Split(nombre, "")
    for _, letra := range letras {
        time.Sleep(1000 * time.Millisecond)
        fmt.Println(letra)
    }
}
```

#### Channels (diálogo entre GoRoutines)

- Los `channels`, permite que nuestras rutinas (paralelismo entre la ejecución de funciones) se comuniquen entre sí. En efecto, permite que cada ejecución paralela que se esté ejecutando en el procesador pueda dialogar con otra enviando información.
- Un `channel`, en definitiva, es un espacio de memoria, de dialogo, y cuando se aloje un valor en ese espacio de memoria la `GoRoutine` que está pidiendo ese valor va a actuar en consecuencia.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    canal1 := make(chan time.Duration)
    go bucle(canal1)
    fmt.Println("LLegué hasta aquí")

    msg := <- canal1
    fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
    inicio := time.Now()
    for i := 0; i < 100000000000; i++ {

    }

    final := time.Now()
    canal <- final.Sub(inicio)
}
```

#### Servidor Web

- Un `servidor web` (_http_) permite recibir `request` (_peticiones_) y poder generar `response` (_respuestas_).

```go
```

- _Ver código completo en "section\_04/server"_

#### Middlewares

- Un `middleware` es un interceptor. Nos permite encapsular la ejecución de una función pre-existente con otra función nueva. Va a ejecutar cierto código previamente al llamado de la función ya grabada.
- Son _interceptores_ que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables.
- Los `middlewares` hoy en día se utilizan para agregarle a todas las rutinas, a todas las funciones, todo lo que tiene que ver con funcionalidad, seguridad, encriptación. Se hacen muchas operaciones como capas de seguridad a través de los _middlewares_.

```go
package main

import (
    "fmt"
)

var result int

func main() {
    fmt.Println("Inicio")

    result = operacionesMidd(sumar)(2,3)
    fmt.Prinln(result)

    result = operacionesMidd(restar)(10,6)
    fmt.Prinln(result)

    result = operacionesMidd(multiplicar)(4,2)
    fmt.Prinln(result)
}

func sumar(a, b int) int {
    return a + b
}

func restar(a, b int) int {
    return a - b
}

func multiplicar(a, b int) int {
    return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
    return func(a, b int) int {
        fmt.Println("Inicio de Operación")
        return f(a, b)
    }
}
```

#### Descarga & Bonus de regalo

```go
```

---

## Sección 05 - Aprendemos a crear toda una aplicación Backend

---

## Sección 06 - Backend - Instalaciones y registros necesarios

---

## Sección 07 - Backend - Configurando MongoDB Atlas

---

## Sección 08 - Backend - Golang (Estructura de aplicación en nuestra PC)

---

## Sección 09

---

## Sección 10

---

## Sección 11

---

## Sección 12

---

## Sección 13

---

## Sección 14

---

## Sección 15

---

## Sección 16

---

## Sección 17

---

## Sección 18

---

## Sección 19

---

## Sección 20

---

## Sección 21

---

## Sección 22

---

## Sección 23

---

## Sección 24

---

## Sección 25

---

## Sección 26
