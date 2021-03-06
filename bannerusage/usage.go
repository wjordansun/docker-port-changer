package bannerusage

var (
	Usage string =
	`Usage:
		build docker image before using init command
		
Commands:
		help -- show this message
		init -- open production and honeypot ports on 3000 and 4000
		start -- start watching packets
		list -- display log files
		ports -- show what ports are open and which ones are honeypots
		reset -- Reset docker containers to production1 on 3000 and honeypot on 4000
		exit/quit -- exit the shell
	`
)