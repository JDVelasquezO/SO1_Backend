# Universidad de San Carlos de Guatemala - Sistemas Operativos 1

## Práctica 1: KVM, Docker y Containers
## 201800722 - José Daniel Velásquez Orozco

### Index
- [Overview](#overview)
- [Frontend](#frontend)
- [Backend](#backend)
- [Database](#database)
- [Containers e Images](#containers)
- [Docker-Compose](#compose)
- [DockerHub](#dockerhub)
- [KVM](#kvm)

### <a name="overview">Overview</a>
A continuación vamos a describir los componentes y herramientas utilizadas para la práctica que trata acerca de un sitio web construído en react que consume servicios de una api construída en golang la cual se comunica con una base de datos en mongoDB separados en contenedores diferentes.

La arquitectura utilizada en la práctica se muestra a continuación:

![](https://i.imgur.com/wl04CZP.png)


### <a name="frontend">Frontend</a>
[Link a repositorio](https://github.com/JDVelasquezO/SO1_Frontend)

Para el cliente se creó un proyecto utilizando ```create react app``` utilizando Typescript y JSX para una mejor escritura de código, se usaron técncias como hooks para manejar el estado de la aplicación de forma reactiva. Para el diseño gráfico se utilizó la librería Material UI y para el entorno de desarrollo se utilizó el IDE WebStorm. 

![](https://i.imgur.com/OtypE9L.png)

Se tiene la siguiente estructura de carpetas:

![](https://i.imgur.com/QXam6ql.png)

A continuación se muestran los componentes utilizados para la construcción de la aplicación:
![](https://i.imgur.com/TLDgdM6.png)

La vista final se muestra a continuación:
![](https://i.imgur.com/XP0X6lY.png)

Para poder conectar el frontend al backend debemos crear un proxy en el archivo ```package.json``` como sigue:
![](https://i.imgur.com/WLwaAt4.png)

Debemos tomar en cuenta el nombre del servicio en el archivo ```docker-compose.yml``` que veremos más adelante.

También debemos tomar en cuenta que cuando hacemos fetch a la petición, solo colocamos el endpoint:
![](https://i.imgur.com/BG6uo3j.png)


### <a name="backend">Backend</a>
Para el servidor se creó una api escrita en el lenguaje de programación Golang y la librería Fiber para realizar peticiones http. Se creó una arquitectura basada en tres capas para crear modelos, controladores y retornar resultados en formato JSON gracias a las rutas y para prueba de endpoints se utilizó el software postman. También se utilizó como entorno de desarollo el IDE Goland.
![](https://i.imgur.com/2VDgzAk.png)

La estructura de carpetas es la siguiente:

![](https://i.imgur.com/TOr3opD.png)

Resultado de una consulta hacia la base de datos:
![](https://i.imgur.com/a8AJUV4.png)


### <a name="database">Database</a>
La base de datos fue escrita en el gestor noSQL de MongoDB, creando colecciones tanto de Vehicle como de Record para almacenarlos en el volumen, este volumen es almacenado en el contenedor de la base de datos en el path: ```mongo-data:/data/db```.
Para visualizar gráficamente los datos se utilizó el IDE Datagrip y para ejecutar la base de datos, se descargó la imagen oficial de docker-hub.

![](https://i.imgur.com/nh8cAws.png)

La tabla Vehicle luce de la siguiente forma:
![](https://i.imgur.com/NTOfIEW.png)

### <a name="containers">Containers e Images</a>
Las imágenes se crearon usando un archivo ```Dockerfile``` tanto para el frontend como para el backend, mongoDB ya tiene su propia imagen subida en dockerhub.

#### Dockerfile para backend
Este nos muestra 7 pasos los cuales son:
- Obtener la última versión del lenguaje Go.
- Crear un directorio de trabajo raíz llamado /backend_app.
- Copiar los archivos del proyecto en este directorio.
- Descargar las librerías necesarias.
- Establecer la variable de entorno de la base de datos.
- Establecer puerto donde se ejecutará la api.
- Ejecutar por consola virtual el comando ```go run Main.go```

![](https://i.imgur.com/dszdxxS.png)

#### Dockerfile para Frontend
Este muestra 8 pasos importantes y son:
- Obtener la última versión de Node.
- Crear un nuevo directorio llamado /frontend_app.
- Copiar el package.json y package-lock.json para obtener las dependencias.
- Ejecutar el comando ```npm install``` para instalar los node_modules necesarios.
- Establecer la variable de entorno para el backend.
- Copiar el resto de archivos al directorio.
- Ejecutar el comando ```npm start```

![](https://i.imgur.com/ZXFoxKn.png)

### <a name="compose">Docker-Compose</a>
Este archivo fue creado específicamente para ejecutar las imágenes subidas a dockerhub y crear contenedores correspondientes localmente.

Este se compone de tres partes, la base de datos, backen y frontend:
#### Base de datos:
Obtenemos la imagen de mongo, creamos un contenedor, establecemos variables de entorno y también el volumen para la persistencia de datos así como el puerto donde deseamos que se ejecute.

![](https://i.imgur.com/gqfiakY.png)


#### Backend:
![](https://i.imgur.com/x0oKUz9.png)

##### Frontend:
![](https://i.imgur.com/Ozsw8B2.png)

#### Configuración de volúmenes:
Esto se hace con la intención de establecer una ruta con un directorio dentro del contenedor el cual guardará nuestros datos
![](https://i.imgur.com/hAQ30lU.png)

### <a name="dockerhub">Dockerhub</a>
Las imágenes respectivas del frontend y backend están disponibles en los siguientes links:

- [Frontend](https://hub.docker.com/repository/docker/jdveloper/so1practica1frontend)
- [Backend](https://hub.docker.com/repository/docker/jdveloper/so1practica1backend)


### <a name="kvm">KVM</a>
Adicionalmente se creó un servidor virtual de la imagen de Ubuntu Server 18.04, utilizando el hipervisor KVM y el entorno gráfico Virtual Machine Manager.

En este server existe una carpeta llamada practica1 y dentro tenemos el archivo ```docker_compose.yml```.
![](https://i.imgur.com/Wz0o8IU.png)

Para ejecutar nuestras imágenes simplemente debemos iniciar el siguiente comando:
```console
docker-compose up -d 
```

El resultado será así:
![](https://i.imgur.com/guOCCTC.png)

Ahora, necesitamos conocer la dirección IP de la máquina virtual, la cual podemos visualizar en la parte de información:
![](https://i.imgur.com/Fa54IcQ.png)
En este caso la Ip es 192.168.122.163.

Ahora debemos buscar en nuestro navegador de la máquina host la dirección Ip y conectarnos al contenedor:
![](https://i.imgur.com/7NAuaw5.png)

De esta forma nos podemos dar cuenta que tanto en máquina virtual como en máquina host tenemos los mismos datos de la misma base de datos almacenada en el contenedor con imagen de mongo.
