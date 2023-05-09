# <span style="color:#C0BFEC">🦔 ***VK bot storage***</span>

## <span style="color:#C0BFEC">📑 ***Description:*** </span>

### <span>🧰 ***Telegram bot has commands for interaction:***</span>

1) `Start-command` to display a welcome message and information about the bot
2) `Set-command` for setting the login and password for the service
3) `Get-command` for getting login and password by service name
4) `Del-command` for deleting login and password by service name

Templates for commands:
```
1. /start
2. /set <service name> <login> <password>
3. /get <service name>
4. /del <service name>
```

### <span>💾 ***Data storage:***</span>

The repository uses a `PostgreSQL` database that runs in a `docker` container

### <span>💻 ***Other:***</span>

To ensure security, all messages containing 
passwords are deleted after a certain time 
set as constant `DeletePasswordDeadline` in `config.go`

## <span style="color:#C0BFEC">🏃🏻‍♂️ ***Run:*** </span>

1) To run the application, you need to specify 
the token of your telegram bot in the `config.yaml` 
configuration file (field `bot.token`)
2) Then you need to run the command to run the 
application in docker containers:
```shell
sudo docker-compose up
```
