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
