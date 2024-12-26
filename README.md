# Desarrollo

Para el desarrollo de la aplicacion se uso el lenguaje de programacion Golang, haciendo uso de Gin-gonic como libreria para levantar el REST API y para la base de datos se use SQLite.

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

