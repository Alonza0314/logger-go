# logger-go

This is a log tool for create error/warn/info/debug/trace/test log.

## Usage

### logger directly

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go/v2"
    ```

2. Use the logger in your project.

    ```go
    loggergo.Info("tag", "message")
    loggergo.Error("tag", "message")
    loggergo.Warn("tag", "message")
    loggergo.Test("tag", "message")
    loggergo.Debug("tag", "message")
    ```

### logger in detail

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go/v2"
    ```

2. Declare the base logger structure.

    ```go
    debugMode := true // if true, log will only output in terminal
    filePath := "logger.log" // the destination log file
    logger := NewLogger(filePath, debugMode)
    defer logger.Close()
    ```

3. Set the log level you want. You can call the util package for pre-declared const log level string.

    ```go
    // valid levels: error, warn, info, debug, trace, test
    logger.Setlevel(util.LEVEL_STRING_INFO)
    ```

4. Set the target tag or tags, this will return an instance for you to use later.

    ```go
    demoSingleTag := logger.WithTag("TAG1")
    demoMultiTags := logger.WithTags("TAG1", "TAG2")
    ```

5. Use logger instance with "f" and "ln".

- Errorf
- Warnf
- Infof
- Debugf
- Tracef
- Testf
- Errorln
- Warnln
- Infoln
- Debugln
- Traceln
- Testln

    ```go
    demoSingleTag.Infof("%s %s", "msg1", "msg2")
    demoMultiTags.Infoln("msg1", "msg2")
    ```

## Gin Engine

Gin writes its own request logging (and panic/error output from `Recovery()`) straight to `os.Stdout`, via the package-level `gin.DefaultWriter` / `gin.DefaultErrorWriter` variables. `loggergo.NewGinWriter` adapts any tagged logger instance into an `io.Writer`, so you can route that output through logger-go instead -- same tags, same file, same level handling as the rest of your app. This does not require importing `github.com/gin-gonic/gin` in your own code beyond what you already use to build the engine.

1. Import the logger-go package in your project.

    ```go
    import loggergo "github.com/Alonza0314/logger-go/v2"
    ```

2. Create (or reuse) a tagged logger instance for gin's own output.

    ```go
    ginLog := logger.WithTag("API")
    ```

3. Wrap it with `NewGinWriter` and assign it to `gin.DefaultWriter` (and `gin.DefaultErrorWriter`, if you also want panic/error output routed the same way) **before** the `gin.Engine` is constructed.

    ```go
    gin.DefaultWriter = loggergo.NewGinWriter(ginLog)
    gin.DefaultErrorWriter = loggergo.NewGinWriter(ginLog)

    router := gin.Default() // or gin.New() -- must come after the two lines above
    ```