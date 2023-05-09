package bot

const (
	SetExample                = "Example \"/set\" command:\n\n \"/set <service> <login> <password>\" "
	GetExample                = "Example \"/get\" command:\n\n \"/get <service>\" "
	DelExample                = "Example \"/del\" command:\n\n \"/del <service>\" "
	ParamsCountErrorMsg       = "Incorrect params count!"
	SetDataErrorMsg           = "Failed to set data for the service"
	GetDataErrorMsg           = "Failed to get data for the service"
	DelDataErrorMsg           = "Failed to del data for the service"
	GetPasswordDecodeErrorMsg = "Failed to get password!"
	SetSuccessMsg             = "Data has been successfully saved for the service"
	GetSuccessMsg             = "Data has been successfully received for the service"
	DelSuccessMsg             = "Data has been successfully deleted for the service"
	StartMsg                  = "Hello! I am a bot for storing logins and passwords for various services created by @p_stpn.\n\n" +
		"You can use my commands:\n\n1) Setting the login and password for the service.\n" +
		"Example:\n\n`/set <service name> <login> <password>` \n\n2) Obtaining a login and" +
		" password by the name of the service.\nExample:\n\n`/get <service name>` \n\n" +
		"3) Removing the login and password for the service by its name:\n\n`/del <service name>` \n\n" +
		"You can find the source code for this project and more on GitHub: \nhttps://github.com/31steets31/VK-bot-storage"
)
