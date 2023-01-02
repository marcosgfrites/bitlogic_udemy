# ASP.NET Core MVC
_<link> Curso: https://www.udemy.com/course/aprende-aspnet-core-mvc-haciendo-proyectos-desde-cero/_
---

## Sección 01 - Introducción

    - ¿Qué es .NET?
    Es una plataforma de desarrollo que sirve para desarrollar distintos tipos de aplicaciones para diferentes entornos (SOs, etc).
    No es un lenguaje de programación, sino un ambiente en donde podemos ejecutar nuestras apps.
    El lenguaje que se utiliza puede ser C# (lenguaje multiparadigmas aunque mayormente orientado a objetos), F# (lenguaje funcional) o Visual Basic (orientado a objetos).
    .NET fue creado por Microsoft en 2002 para correr en Windows.
    .NET Framework llegó hasta 4.8
    .NET Core es la evolución de .NET Framework
    .Net Core salió en el año 2016. Desde la versión 1 hasta la 3, luego saltaron a la versión 5: .NET 5 (se quitó el `Core`).

    - ¿Qué es C#?
    Es un lenguaje de programación multiparadigmas, creado en el año 2000 por Microsoft.
    Permite el desarrollo de distintos tipos de aplicaciones: webapps, mobiles, destock apps, etc.
    Permite desarrollar programación orientada a objetos o programación funcional.
    Es un lenguaje fuertemente tipado.
    En cuanto a la POO, está basado en clases.
    Nos permite trabajar con funciones de primera clase: funciones como parámetros, retornar funciones, varaibles con asignación de funciones, funciones anónimas, etc.
    LINQ y expresiones lambda: característica de la programación funcional para realizar consultas a bases de datos apoyándonos en una librería llamada `Entity Framework Core`.

    - ¿Qué es ASP.NET Core?
    Es un framework para desarrollar webapps multiplataformas.
    ASP = Active Server Pages -> `Dinamismo`
    Sistema de ruteo: nos permite indicar qué función queremos ejecutar ante una petición `HTTP`.
    Sistema de usuarios: nos permite registrar a nuestros usuarios en nuestras apps.
    Sistema de inyección de dependencias: nos permite aplicar principios de desarrollo de software como el `principio de inversión de dependencias` donde podemos hacer que nuestras clases dependan de abstracciones y no de tipos concretos.

    - ASP.NET Core MVC
    MVC se refiere al `Pattern Model View Controller`: nos permite crear páginas web cuyo contenido va a ser renderizado (`Razor`/`Blazor`) por un servidor. El servidor genera el HTML que va a ser servido al usuario. Se centra en separar la manipulación de datos de la visualización de la misma.
    `Model`: se refiere a los datos de nuestra app. Ejemplo, usuarios, productos, etc, es decir cualquier concepto concreto que contenga datos a ser mostrados. Los modelos de una app dependen del negocio que maneja.
    `View`: se refiere a la plantilla que muestra la data de un `model`. En resumen, es la parte visual de nuestra app.
    `Controller`: se encarga de actualizar el `model` y le pasa los datos a la `view` para que muestr el contenido final al usuario. En resumen, es quien responde a las acciones del usuario; recibe las peticiones.

    - Razor Pages
    Es similar a `MVC`. Todo lo que se realiza con MVC puede hacerse con `Razor Pages` y viceversa. Aunque la diferencia radica en que en lugar de tener un gran `controller` que maneje muchas `views`, contamos con `mini controllers` que se enfocan, cada uno, en una `view` particular.
    Quien recibe las `request` es un `page handler`, la maneja y completa cualquier información que sea necesario de la página en cuestión; genera la vista y es servida al usuario (similar a `MVC`).

    - Web APIs
    Se utilizan cuando no queremos generar interfaces de usuario. Sino que sólo manejamos una parte lógica, puesto que el `frontend` lo manejaremos con diferentes `frameworks`.
    En términos de MVC, una `Web API` elimina la `view`.
    Se pueden combinar MVC con Web APIs.

    - Blazor
    Envía el código de C# al navegador del usuario y allí generar las `views`.
    Nos permite crear aplicaciones web interactivas.
    Existen 2 modelos de `Blazor`:
    - `WebAssembly`: podemos enviar el código de C# al navegador del usuario y ejecutarlo ahí mismo.
    - `Server`: podemos ejecutar la aplicación remotamente desde un servidor, así el usuario interactúa con ella en una conexión de tiempo real.

    - gRPC
    Es un framework que nos permite hacer llamadas de procedimiento remoto. Es proveniente de `google`.
    Nos permite comunicarnos de una manera muy rápida entre distintas aplicaciones. Como por ejemplo, en el uso de `microservices`.

## Sección 02 - Introducción a Bootstrap

## Sección 03 - Proyecto 1 - Portafolio

## Sección 04 - Proyecto 2 - Manejo de Presupuestos - Base de Datos

## Sección 05 - Proyecto 2 - Manejo de Presupuestos - CRUD con Dapper

## Sección 06 - Proyecto 2 - Manejo de Presupuestos - Los Otros CRUDs
