# logger-go

This is a log tool for create info/error/warn/test/debug log.

## Usage

### logger in STD_OUT

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go"
    ```

2. Use the logger in your project.

    Example:

    ```go
    loggergo.Info("tag", "message")
    loggergo.Error("tag", "message")
    loggergo.Warn("tag", "message")
    loggergo.Test("tag", "message")
    loggergo.Debug("tag", "message")
    ```

### logger in file

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go"
    ```

2. Declare a logger instance.

    ```go
    logger := loggergo.NewFileLogger("logger.log")
    defer logger.Close()
    ```

    At here, you can use the `WithFlag`, `WithPerm` and `WithColor` to set the flag and permission of the file.

    Example:

    ```go
    logger := loggergo.NewFileLogger("logger.log", loggergo.WithFlag(os.O_APPEND|os.O_CREATE|os.O_WRONLY), loggergo.WithPerm(os.FileMode(0644)), loggergo.WithColor(true))
    ```

3. Use the logger in your project.

    ```go
    logger.Info("tag", "message")
    logger.Error("tag", "message")
    logger.Warn("tag", "message")
    logger.Test("tag", "message")
    logger.Debug("tag", "message")
    ```

### Executable file

1. Download the executable file from [release](https://github.com/Alonza0314/logger-go/releases).

2. Use the executable file directly.

    Example:

    ```bash
    ./logger-go info [tag] [message]
    ./logger-go error [tag] [message]
    ./logger-go warn [tag] [message]
    ./logger-go test [tag] [message]
    ```

> [!NOTE]
> The executable file is now built for:
> - Linux-amd64
> - Windows-amd64
> - Darwin-amd64
> - Darwin-arm64
