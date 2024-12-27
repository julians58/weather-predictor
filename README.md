# Desarrollo

Para el desarrollo de la aplicacion se uso el lenguaje de programacion Golang, haciendo uso de Gin-gonic como libreria para levantar el REST API y para la base de datos se uso SQLite junto a gorm como ORM para practicidad con las queries SQL.

## Para tener en cuenta

Se tomo en cuenta que todos los planetas parten desde el instante cero en la misma posicion, es decir alineados sobre el eje x con 0 grados.

El primer dia (instante cero del que parten los planetas en su movimiento circular uniforme) es el dia 0 sin embargo en la API se cuentan como desde el dia 1, es decir el dia 1 seria el instante cero, con todos los planetas alineados y con condiciones climaticas de "Sequia".

Para obtener la alineacion de los planetas se uso el metodo de calcular las pendientes, en el escenario de calcular si los planetas estan alineados con el origen (el sol) calculamos la pendiente de cada planeta con respecto al sol y comparamos que sean iguales, para el escenario donde estan alineados entre ellos pero no con el sol usamos un metodo similar, calculamos la pendiente entre el planeta de los Ferengi con el de lo Vulcanos y luego el de los Vulcanos con el de los Betazoide y los comparamos tomando una constante de tolerancia, esta variable es editable para ajustar precision y es necesaria ya que practicamente no encontrariamos ningun dia donde se cumpliera que existe una recta exacta.

## Correr proyecto 
Instalar dependencias :

    go get github.com/gin-gonic/gin
    go get gorm.io/gorm
    go get gorm.io/driver/sqlite
Correr aplicacion:

    go run main.go
Esto creara una base de datos 'weather.db' local

Para ver el registro de todos los dias de los 10 años:

    http://localhost:8080/weather
Para ver la condicion climatica de un dia particular (En el ejemplo veriamos el primer dia de todos donde todos los planetas parten alineados con angulo 0 y en condicion de sequia):

    http://localhost:8080/weather/1

Para ver por climas:

    http://localhost:8080/days/Lluvia

Para ver las estadisticas:

    http://localhost:8080/weather-statistics

Tener en cuenta que existen los siguientes climas:
| Clima | Particularidad |
|--|--|
| Sequia |Planetas alineados entre si y con el sol  |
|Lluvia|Planetas formando triangulo con sol en su interior|
|Pico de lluvia|Perimetro maximo de un periodo de lluvia|
|Condiciones normales|Planetas formando un triangulo sin el sol en su interior|
|Condiciones optimas de presion y temperatura|Planetas alineado entre si pero no con el origen|

## Data en 10 años
    {"ID":1,"drought_periods":41,"rain_periods":61,"optimal_periods":122,"rain_peak_days":"[25,73,155,205,253,335,385,433,515,565,613,695,745,793,875,925,973,1055,1105,1153,1235,1285,1333,1415,1465,1513,1595,1645,1693,1775,1825,1873,1955,2005,2053,2135,2185,2233,2315,2365,2413,2495,2545,2593,2675,2725,2773,2855,2905,2953,3035,3085,3133,3215,3265,3313,3395,3445,3493,3575,3625]"}
Periodos de - Sequia: 41, Lluvia: 61, Optimas condiciones: 122
