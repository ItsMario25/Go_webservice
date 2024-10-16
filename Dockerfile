# Establecer la imagen base
FROM golang:1.20-alpine

# Crear y definir el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto donde Go estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
