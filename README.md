# logger-go

This is a log tool for simply create info/error/warn/test log.

## Usage

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go"
    ```

    Example:

    ```go
    loggergo.Info("tag", "message")
    loggergo.Error("tag", "message")
    loggergo.Warn("tag", "message")
    loggergo.Test("tag", "message")
    ```

2. Use the executable file directly.

    Download the executable file from [release](https://github.com/Alonza0314/logger-go/releases) and use it directly.

    Example:

    ```bash
    logger-go info [tag] [message]
    logger-go error [tag] [message]
    logger-go warn [tag] [message]
    logger-go test [tag] [message]
    ```

> [!NOTE]
> The executable file is now built for:
> - Linux-amd64
> - Windows-amd64
