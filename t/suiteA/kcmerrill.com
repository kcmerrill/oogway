kcmerrill.com:
    summary: My description would go here
    try: 4
    check: 
        cmd: touch /tmp/suiteA.check.kcmerrill.com
        reset: 1h
    fix:
        cmd: touch /tmp/suiteA.fix.kcmerrill.com
        after: 2
    warning: 
        cmd: touch /tmp/suiteA.warning.kcmerrill.com
    critical:
        cmd: touch /tmp/suiteA.critical.kcmerrill.com
    recover:
        cmd: touch /tmp/suiteA.recover.kcmerrill.com