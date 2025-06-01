# brainfuck-go

**brainfuck-go** es un intérprete del lenguaje de programación esotérico [Brainfuck](https://esolangs.org/wiki/Brainfuck), desarrollado completamente en Go. Este proyecto busca ser simple, eficiente y educativo, ideal tanto para quienes desean ejecutar programas Brainfuck como para quienes quieren entender cómo se implementa un intérprete desde cero.

## Características

- Compatible con el conjunto estándar de instrucciones de Brainfuck (`+`, `-`, `>`, `<`, `[`, `]`, `,`, `.`)
- Lectura desde archivos o entrada directa
- Estructura clara y modular del código
- Escrita 100% en Go

## Uso

### Desde el código fuente

```bash
go run . archivo.bf
```

### Desde los ejecutables compilados

También puedes usar las versiones precompiladas ubicadas en el directorio `bin/`.

#### En Linux / macOS

```bash
./bin/bf archivo.bf
```

#### En Windows

```powershell
.\bin\bf.exe archivo.bf
```

> Asegúrate de que el archivo `.bf` esté en el mismo directorio o proporciona la ruta completa al archivo.

## 📄 Ejemplos

El repositorio incluye algunos programas de ejemplo escritos en Brainfuck, como el clásico "Hello World!". Puedes probarlos con:

```bash
./bin/bf ejemplos/helloworld.bf
```

## 🛠️ Requisitos

* Go 1.24.2 o superior (solo si deseas compilar desde el código fuente)

## 🧠 ¿Por qué Brainfuck?

Brainfuck es un lenguaje minimalista que permite explorar conceptos fundamentales de los lenguajes de programación, como la manipulación de memoria, bucles y estructuras de control. Implementarlo es una excelente forma de aprender sobre parsing, ejecución y estructuras de datos.

## 📦 Compilar manualmente

Para compilar los ejecutables tú mismo:

```bash
# Para Linux
GOOS=linux GOARCH=amd64 go build -o bin/bf main.go

# Para Windows
GOOS=windows GOARCH=amd64 go build -o bin/bf.exe main.go
```

## 📚 Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más información.
