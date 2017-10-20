google.com:
    description: check google every 
    try: 4
    check: 
        cmd: touch /tmp/suiteA.check.google.com
        reset: 1h
    fix:
        cmd: touch /tmp/suiteA.fix.google.com
        after: 2
    warning: 
        cmd: touch /tmp/suiteA.warning.google.com
    critical:
        cmd: touch /tmp/suiteA.critical.google.com
    recover:
        cmd: touch /tmp/suiteA.recover.google.com