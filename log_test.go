package rslog

import (
        "testing"
)


func TestRegisterRsyslogHandler(t *testing.T) {
    if ok := RegisterRsyslogHandler("./conf/rsyslog.xml"); ok == false {
        t.Log("register rsyslog handler fail")
        t.Fail()
    }

}

