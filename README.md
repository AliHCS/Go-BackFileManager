# Go File Manager

## Descripción

Go File Manager es una aplicación de gestión de archivos desarrollada en Go. Permite a los usuarios subir, descargar y gestionar archivos de manera eficiente.

## Características

- Subida de archivos.
- Descarga de archivos.
- Eliminación de archivos.
- Gestión de usuarios.
- Soporte para archivos de diferentes tipos.

## Requisitos

Antes de comenzar, asegúrate de tener instalados los siguientes programas en tu sistema:

- [Go](https://golang.org/dl/) (versión 1.17 o superior)
- Git

## Instalación

1. Clona el repositorio:

   git clone https://github.com/AliHCS/Go-BackFileManager.git

2. Navega al directorio del proyecto:
   
   cd Go-BackFileManager
  
3. Copia el archivo de ejemplo de configuración:

  env.example a .env
  
4. Abre el archivo .env y establece el puerto deseado. Por ejemplo:

  PORT=8080

5.Instala las dependencias

  go mod tidy


Para levantar el servidor, ejecuta el siguiente comando:
  go run main.go


# Licencia
Este proyecto está bajo la Licencia MIT - consulta el archivo LICENSE para más detalles.
