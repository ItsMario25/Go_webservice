# Establecer la imagen base
FROM golang:1.23-alpine

# Crear y definir el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar el código fuente al contenedor
COPY . .

# Descargar las dependencias
RUN go mod download
RUN go mod tidy

# Compilar la aplicación Go
RUN go build -o main .

# Exponer el puerto donde Go estará escuchando
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
