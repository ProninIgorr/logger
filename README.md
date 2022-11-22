# logger

Add methods:
-SetLogLevel(logLvl LogLevel);
-Infoln(msg string);
-Warnln(msg string);
-Errorln(msg string).
-LogLevel — an enumeration with such posibble value:
Info;
Warning;
Error.
Functions Infoln, Warnln и Errorln wrap method log.Println, adding the appropriate prefix.
