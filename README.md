
# Send email 

    USAGE:
        mailer [global options] command [command options] [arguments...]
        
    GLOBAL OPTIONS:
      -debug, -d
        	flag set to debug (default false)
    
      -config, -c  string
        CONFIG_FILEPATH	(environment var)
        	config file path
    
      -proxyhttp  string
        HTTP_PROXY	(environment var)
        	Sets http_proxy for network connections
    
      -proxyhttps  string
        HTTPS_PROXY	(environment var)
        	Sets https_proxy for network connections
    
      -noproxy  string
        NO_PROXY	(environment var)
        	Sets no_proxy for network connections
    
      -skip_update, -skip
        SKIP_UPDATE	(environment var)
        	set flag to skip update check (default false)
    
      -log_dir, -ld  string
        LOG_DIR	(environment var)
        	set logging directory (default $HOME)
    
      -from  string
        FROM	(environment var)
        	Who is the email from `user@domain.com`
    
      -to  string
        TO	(environment var)
        	Whom to send the email to `user@domain.com`
    
      -cc  string
        CC	(environment var)
        	Whom to cc on the email `user@domain.com`
    
      -ccname  string
        CCNAME	(environment var)
        	Name to cc on the email
    
      -subject, -s  string
        SUBJECT	(environment var)
        	Subject of the Email (default Test Message)
    
      -html
        HTML	(environment var)
        	Message body default plain set for html (default false)
    
      -message, -m  string
        MESSAGE	(environment var)
        	Message body to send in either quoted plaintext or html (default Hello, test message)
    
      -file, -f  string
        FILE	(environment var)
        	Attach a file `FILE`
    
      -smtp_server  string
        SMTP_SERVER	(environment var)
        	SMTP Server (default localhost)
    
      -smtp_port  int
        SMTP_PORT	(environment var)
        	SMTP Port (default 587)
    
      -smtp_username  string
        SMTP_USERNAME	(environment var)
        	SMTP Username
    
      -smtp_password  string
        SMTP_PASSWORD	(environment var)
        	SMTP Password
    
    
    COMMANDS:
      update:       (check for updates)
      buildconfig:  (build generic configuration you can fill in)
      send:         (send email)