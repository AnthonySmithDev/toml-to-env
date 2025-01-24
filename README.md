# TOML to ENV CLI

Este es un CLI simple en Go que convierte un archivo TOML en un archivo `.env`. Es útil para convertir configuraciones en formato TOML a variables de entorno en formato `.env`.

## Instalación

Para usar este CLI, primero asegúrate de tener Go instalado en tu sistema. Luego, puedes clonar este repositorio y compilar el código:

```bash
git clone https://github.com/tu-usuario/toml-to-env-cli.git
cd toml-to-env-cli
go build -o toml-to-env
```

Esto generará un binario llamado `toml-to-env` que puedes usar directamente.

## Uso

El CLI toma un archivo TOML como entrada y genera un archivo `.env` como salida. Puedes especificar la ruta del archivo TOML de entrada y, opcionalmente, la ruta del archivo `.env` de salida.

### Sintaxis básica

```bash
./toml-to-env -input <ruta_del_archivo_toml> [-output <ruta_del_archivo_env>]
```

### Ejemplo

Supongamos que tienes un archivo `config.toml` con el siguiente contenido:

```toml
[database]
host = "localhost"
port = 5432
user = "admin"
password = "secret"

[server]
host = "0.0.0.0"
port = 8080
```

Puedes convertir este archivo TOML a un archivo `.env` ejecutando:

```bash
./toml-to-env -input config.toml -output .env
```

Esto generará un archivo `.env` con el siguiente contenido:

```env
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=admin
DATABASE_PASSWORD=secret
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
```

### Opciones

- `-input`: Ruta del archivo TOML de entrada. Este parámetro es obligatorio.
- `-output`: Ruta del archivo `.env` de salida. Si no se especifica, por defecto se usará `.env`.

## Contribuciones

Si encuentras algún problema o tienes alguna sugerencia, no dudes en abrir un issue o enviar un pull request.

## Licencia

Este proyecto está licenciado bajo la [MIT License](LICENSE).
