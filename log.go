package rslog

import (
        "log/syslog"
        log "github.com/cihub/seelog"
)


var levelToSyslogSeverity = map[log.LogLevel]syslog.Priority{
        // Mapping to rsyslog level
        log.TraceLvl:    syslog.LOG_DEBUG,
        log.DebugLvl:    syslog.LOG_DEBUG,
        log.InfoLvl:     syslog.LOG_INFO,
        log.WarnLvl:     syslog.LOG_WARNING,
        log.ErrorLvl:    syslog.LOG_ERR,
        log.CriticalLvl: syslog.LOG_CRIT,
        log.Off:         syslog.LOG_DEBUG,
}


type CustomReceiver struct { // implements seelog.CustomReceiver
        net        string // using net(tcp, udp) way to send log to rsyslog
        addr       string // rsyslog addrs default "127.0.0.1:514"
        tag        string // a tag mark log info default "rslog"
}


func (rsyslog *CustomReceiver) ReceiveMessage(message string, level log.LogLevel, context log.LogContextInterface) error {
    writer, err := syslog.Dial(rsyslog.net, rsyslog.addr, levelToSyslogSeverity[level], rsyslog.tag)
    if err != nil {
        panic(err)
    }

    switch level {
        case log.TraceLvl, log.DebugLvl, log.Off:
            writer.Debug(message)
        case log.InfoLvl:
            writer.Info(message)
        case log.WarnLvl:
            writer.Warning(message)
        case log.ErrorLvl:
            writer.Err(message)
        case log.CriticalLvl:
            writer.Crit(message)
        default:
            writer.Err("May be it's a new level: " + message)
    }

    writer.Close()

    return nil
}

//Parse some args
func (rsyslog *CustomReceiver) AfterParse(initArgs log.CustomReceiverInitArgs) error {
    var ok bool

    rsyslog.net, ok = initArgs.XmlCustomAttrs["net"]
    if !ok {
        rsyslog.net = "tcp"
    }

    rsyslog.addr, ok = initArgs.XmlCustomAttrs["addr"]
    if !ok {
        rsyslog.addr = "127.0.0.1:514"
    }

    rsyslog.tag, ok = initArgs.XmlCustomAttrs["tag"]
    if !ok {
        rsyslog.tag = "rslog"
    }

    return nil
}

func (rsyslog *CustomReceiver) Flush() {

}

func (rsyslog *CustomReceiver) Close() error {
    return nil
}

func RegisterRsyslogHandler(loggerconf string) bool{
    var err error

    log.RegisterReceiver("rsyslog", &CustomReceiver{})
    logger, err := log.LoggerFromConfigAsFile(loggerconf)
    if err != nil {
        panic(err)
        return false
    }

    defer logger.Flush()
    err = log.ReplaceLogger(logger)
    if err != nil {
        panic(err)
        return false
    }

    return true
}

// Tracef formats message according to format specifier
// and writes to default logger with log level = Trace.
func Tracef(format string, params ...interface{}){
    log.Tracef(format, params)
}

// Debugf formats message according to format specifier
// and writes to default logger with log level = Deb.
func Debugf(format string, params ...interface{}){
    log.Debugf(format, params)
}

// Infof formats message according to format specifier
// and writes to default logger with log level = Inf.
func Infof(format string, params ...interface{}){
    log.Infof(format, params)
}

// Warnf formats message according to format specifier
// and writes to default logger with log level = War.
func Warnf(format string, params ...interface{}){
    log.Warnf(format, params)
}

// Errorf formats message according to format specifier
// and writes to default logger with log level = Err.
func Errorf(format string, params ...interface{}){
    log.Errorf(format, params)
}

//Criticalf formats message according to format specifier
// and writes to default logger with log level = Cri.
func Criticalf(format string, params ...interface{}){
    log.Criticalf(format, params)
}

// Trace formats message according to format specifier
// and writes to default logger with log level = Trace.
func Trace(v ...interface{}){
    log.Trace(v)
}

// Debug formats message according to format specifier
// and writes to default logger with log level = Deb.
func Debug(v ...interface{}){
    log.Debug(v)
}

// Info formats message according to format specifier
// and writes to default logger with log level = Inf.
func Info(v ...interface{}){
    log.Info(v)
}

// Warn formats message according to format specifier
// and writes to default logger with log level = War.
func Warn(v ...interface{}){
    log.Warn(v)
}

// Error formats message according to format specifier
// and writes to default logger with log level = Err.
func Error(v ...interface{}){
    log.Error(v)
}

// Critical formats message according to format specifier
// and writes to default logger with log level = Cri.
func Critical(v ...interface{}){
    log.Critical(v)
}
