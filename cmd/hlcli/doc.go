package main

// Code generated by gen.go. DO NOT EDIT.

var methodParamMap = map[string][]string{
	"NewSessionAndTokenID": {},
	"SetSessionAndTokenID": {"sessionID", "tokenID"},
	"GlobalConfig":         {},
	"NetworkTypes":         {},
	"PCAssistantConfig":    {},
	"DeviceConfig":         {},
	"WebUIConfig":          {},
	"SmsConfig":            {},
	"WlanConfig":           {},
	"DhcpConfig":           {},
	"CradleStatusInfo":     {},
	"CradleMACSet":         {"addr"},
	"CradleMAC":            {},
	"AutorunVersion":       {},
	"DeviceBasicInfo":      {},
	"PublicKey":            {},
	"DeviceControl":        {"code"},
	"DeviceReboot":         {},
	"DeviceReset":          {},
	"DeviceBackup":         {},
	"DeviceShutdown":       {},
	"DeviceFeatures":       {},
	"DeviceInfo":           {},
	"DeviceModeSet":        {"mode"},
	"FastbootFeatures":     {},
	"PowerFeatures":        {},
	"TetheringFeatures":    {},
	"SignalInfo":           {},
	"ConnectionInfo":       {},
	"ConnectionProfile":    {"roaming", "maxIdleTime"},
	"GlobalFeatures":       {},
	"Language":             {},
	"LanguageSet":          {"lang"},
	"NotificationInfo":     {},
	"SimInfo":              {},
	"StatusInfo":           {},
	"TrafficInfo":          {},
	"TrafficClear":         {},
	"MonthInfo":            {},
	"WlanMonthInfo":        {},
	"NetworkInfo":          {},
	"WifiFeatures":         {},
	"ModeList":             {},
	"ModeInfo":             {},
	"ModeNetworkInfo":      {},
	"ModeSet":              {"netMode", "netBand", "lteBand"},
	"PinInfo":              {},
	"PinEnter":             {"pin"},
	"PinActivate":          {"pin"},
	"PinDeactivate":        {"pin"},
	"PinChange":            {"pin", "new"},
	"PinEnterPuk":          {"puk", "new"},
	"PinSaveInfo":          {},
	"PinSimlockInfo":       {},
	"Connect":              {},
	"Disconnect":           {},
	"ProfileInfo":          {},
	"ProfileAdd":           {"name", "apn", "user", "password", "isDefault"},
	"ProfileDelete":        {"index", "newDefault"},
	"SmsFeatures":          {},
	"SmsList":              {"boxType", "page", "count", "sortByName", "ascending", "unreadPreferred"},
	"SmsCount":             {},
	"SmsSend":              {"msg", "to"},
	"SmsSendStatus":        {},
	"SmsReadSet":           {"id"},
	"SmsDelete":            {"id"},
	"UssdStatus":           {},
	"UssdCode":             {"code"},
	"UssdContent":          {},
	"UssdRelease":          {},
	"DdnsList":             {},
	"LogPath":              {},
	"LogInfo":              {},
	"PhonebookGroupList":   {"page", "count", "sortByName", "ascending"},
	"PhonebookCount":       {},
	"PhonebookImport":      {"group"},
	"PhonebookDelete":      {"id"},
	"PhonebookList":        {"group", "page", "count", "sim", "sortByName", "ascending", "keyword"},
	"PhonebookCreate":      {"group", "name", "phone", "sim"},
	"FirewallFeatures":     {},
	"DmzConfig":            {},
	"DmzConfigSet":         {"enabled", "dmzIPAddress"},
	"SipAlg":               {},
	"SipAlgSet":            {"port", "enabled"},
	"NatType":              {},
	"NatTypeSet":           {"ntype"},
	"Upnp":                 {},
	"UpnpSet":              {"enabled"},
}

var methodCommentMap = map[string]string{
	"NewSessionAndTokenID": "NewSessionAndTokenID starts a session with the server, and returns the session and token.",
	"SetSessionAndTokenID": "SetSessionAndTokenID sets the sessionID and tokenID for the Client.",
	"GlobalConfig":         "GlobalConfig retrieves global Hilink configuration.",
	"NetworkTypes":         "NetworkTypes retrieves available network types.",
	"PCAssistantConfig":    "PCAssistantConfig retrieves PC Assistant configuration.",
	"DeviceConfig":         "DeviceConfig retrieves device configuration.",
	"WebUIConfig":          "WebUIConfig retrieves WebUI configuration.",
	"SmsConfig":            "SmsConfig retrieves device SMS configuration.",
	"WlanConfig":           "WlanConfig retrieves basic WLAN settings.",
	"DhcpConfig":           "DhcpConfig retrieves DHCP configuration.",
	"CradleStatusInfo":     "CradleStatusInfo retrieves cradle status information.",
	"CradleMACSet":         "CradleMACSet sets the MAC address for the cradle.",
	"CradleMAC":            "CradleMAC retrieves cradle MAC address.",
	"AutorunVersion":       "AutorunVersion retrieves device autorun version.",
	"DeviceBasicInfo":      "DeviceBasicInfo retrieves basic device information.",
	"PublicKey":            "PublicKey retrieves webserver public key.",
	"DeviceControl":        "DeviceControl sends a control code to the device.",
	"DeviceReboot":         "DeviceReboot restarts the device.",
	"DeviceReset":          "DeviceReset resets the device configuration.",
	"DeviceBackup":         "DeviceBackup backups device configuration and retrieves backed up configuration data as a base64 encoded string.",
	"DeviceShutdown":       "DeviceShutdown shuts down the device.",
	"DeviceFeatures":       "DeviceFeatures retrieves device feature information.",
	"DeviceInfo":           "DeviceInfo retrieves general device information.",
	"DeviceModeSet":        "DeviceModeSet sets the device mode (0-project, 1-debug).",
	"FastbootFeatures":     "FastbootFeatures retrieves fastboot feature information.",
	"PowerFeatures":        "PowerFeatures retrieves power feature information.",
	"TetheringFeatures":    "TetheringFeatures retrieves USB tethering feature information.",
	"SignalInfo":           "SignalInfo retrieves network signal information.",
	"ConnectionInfo":       "ConnectionInfo retrieves connection (dialup) information.",
	"ConnectionProfile":    "ConnectionProfile set connection (dialup) information for roaming and max idle time.",
	"GlobalFeatures":       "GlobalFeatures retrieves global feature information.",
	"Language":             "Language retrieves current language.",
	"LanguageSet":          "LanguageSet sets the language.",
	"NotificationInfo":     "NotificationInfo retrieves notification information.",
	"SimInfo":              "SimInfo retrieves SIM card information.",
	"StatusInfo":           "StatusInfo retrieves general device status information.",
	"TrafficInfo":          "TrafficInfo retrieves traffic statistic information.",
	"TrafficClear":         "TrafficClear clears the current traffic statistics.",
	"MonthInfo":            "MonthInfo retrieves the month download statistic information.",
	"WlanMonthInfo":        "WlanMonthInfo retrieves the WLAN month download statistic information.",
	"NetworkInfo":          "NetworkInfo retrieves network provider information.",
	"WifiFeatures":         "WifiFeatures retrieves wifi feature information.",
	"ModeList":             "ModeList retrieves available network modes.",
	"ModeInfo":             "ModeInfo retrieves network mode settings information.",
	"ModeNetworkInfo":      "ModeNetworkInfo retrieves current network mode information.",
	"ModeSet":              "ModeSet sets the network mode.",
	"PinInfo":              "PinInfo retrieves SIM PIN status information.",
	"PinEnter":             "PinEnter enters a SIM PIN.",
	"PinActivate":          "PinActivate activates a SIM PIN.",
	"PinDeactivate":        "PinDeactivate deactivates a SIM PIN.",
	"PinChange":            "PinChange changes a SIM PIN.",
	"PinEnterPuk":          "PinEnterPuk enters a SIM PIN puk.",
	"PinSaveInfo":          "PinSaveInfo retrieves SIM PIN save information.",
	"PinSimlockInfo":       "PinSimlockInfo retrieves SIM lock information.",
	"Connect":              "Connect connects the Hilink device to the network provider.",
	"Disconnect":           "Disconnect disconnects the Hilink device from the network provider.",
	"ProfileInfo":          "ProfileInfo retrieves profile information (ie, APN).",
	"ProfileAdd":           "Add connection profile and set new default profile",
	"ProfileDelete":        "Delete connection profile an set new default profile",
	"SmsFeatures":          "SmsFeatures retrieves SMS feature information.",
	"SmsList":              "SmsList retrieves list of SMS in an inbox.",
	"SmsCount":             "SmsCount retrieves count of SMS per inbox type.",
	"SmsSend":              "SmsSend sends an SMS.",
	"SmsSendStatus":        "SmsSendStatus retrieves SMS send status information.",
	"SmsReadSet":           "SmsReadSet sets the read status of a SMS.",
	"SmsDelete":            "SmsDelete deletes a specified SMS.",
	"UssdStatus":           "UssdStatus retrieves current USSD session status information.",
	"UssdCode":             "UssdCode sends a USSD code to the Hilink device.",
	"UssdContent":          "UssdContent retrieves content buffer of the active USSD session.",
	"UssdRelease":          "UssdRelease releases the active USSD session.",
	"DdnsList":             "DdnsList retrieves list of DDNS providers.",
	"LogPath":              "LogPath retrieves device log path (URL).",
	"LogInfo":              "LogInfo retrieves current log setting information.",
	"PhonebookGroupList":   "PhonebookGroupList retrieves list of the phonebook groups.",
	"PhonebookCount":       "PhonebookCount retrieves count of phonebook entries per group.",
	"PhonebookImport":      "PhonebookImport imports SIM contacts into specified phonebook group.",
	"PhonebookDelete":      "PhonebookDelete deletes a specified phonebook entry.",
	"PhonebookList":        "PhonebookList retrieves list of phonebook entries from a specified group.",
	"PhonebookCreate":      "PhonebookCreate creates a new phonebook entry.",
	"FirewallFeatures":     "FirewallFeatures retrieves firewall security feature information.",
	"DmzConfig":            "DmzConfig retrieves DMZ status and IP address of DMZ host.",
	"DmzConfigSet":         "DmzConfigSet enables or disables the DMZ and the DMZ IP address of the device.",
	"SipAlg":               "SipAlg retrieves status and port of the SIP application-level gateway.",
	"SipAlgSet":            "SipAlgSet enables/disables SIP application-level gateway and sets SIP port.",
	"NatType":              "NatType retrieves NAT type.",
	"NatTypeSet":           "NatTypeSet sets NAT type (values: 0, 1).",
	"Upnp":                 "Upnp retrieves the status of UPNP.",
	"UpnpSet":              "UpnpSet enables/disables UPNP.",
}
