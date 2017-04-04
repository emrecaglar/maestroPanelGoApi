package maestropanel

type action struct {
	URL    string
	Method string
}

//WEB ACTIONS
var getListAction = action{"/domain/getlist", "GET"}
var createDomainAction = action{"/domain/create", "POST"}
var deleteDomainAction = action{"/domain/delete", "DELETE"}
var stopDomainAction = action{"/domain/stop", "POST"}
var passwordChangeDomainAction = action{"/domain/password", "POST"}
var addDominAliasAction = action{"/domain/adddomainalias", "POST"}
var deleteDominAliasAction = action{"/domain/deletedomainalias", "DELETE"}
var getDomainAliasesAction = action{"/domain/getdomainaliases", "GET"}
var addSubdomainAction = action{"/domain/addsubdomain", "POST"}
var deleteSubdomainAction = action{"/domain/deletesubdomain", "DELETE"}
var subDomainsAction = action{"/domain/getsubdomains", "GET"}
var setSubDomainFTPAccountAction = action{"/domain/setsubdomainftpaccount", "POST"}
var changeIPAddressAction = action{"/domain/changeipaddr", "POST"}
var getDomainListItemAction = action{"/domain/getlistitem", "GET"}
var getLimitsAction = action{"/domain/getlimits", "GET"}
var forwardingAction = action{"/domain/forwarding", "POST"}
var changeResellerAction = action{"/domain/changereseller", "POST"}
var setDomainPlanAction = action{"/domain/setdomainplan", "POST"}
var changeNETRuntimeVersionAction = action{"/domain/changedotnetruntimeversion", "POST"}
var getNETRuntimeVersionAction = action{"/domain/getdotnetruntimeversion", "GET"}

//FILE MANAGER ACTIONS
var setWriteAccessAction = action{"/domain/setwriteaccess", "POST"}
var revokeWriteAccessAction = action{"/domain/revokewriteaccess", "POST"}
var createDirectoryAction = action{"/domain/createdirectory", "POST"}
var getItemsAction = action{"/domain/getitems", "GET"}
var deleteItemsAction = action{"/domain/deleteitems", "DELETE"}
var zipItemAction = action{"/domain/zipitem", "POST"}
var unZipItemAction = action{"/domain/unzipitem", "POST"}

//WEB STATS ACTIONS
var protectStatsAreaAction = action{"/domain/protectstatsarea", "POST"}
var unProtectStatsArea = action{"/domain/unprotectstatsarea", "POST"}
var enableStatsProtection = action{"/domain/enablestatsprotection", "POST"}
var disableStatsProtection = action{"/domain/disablestatsprotection", "POST"}
var setFtpUserStatsArea = action{"/domain/setftpuserstatsarea", "POST"}

//MAIL ACTIONS
var addMailBoxAction = action{"/domain/addmailbox", "POST"}
var deleteMailBoxAction = action{"/domain/deletemailbox", "POST"}
var changeMailBoxPasswordAction = action{"/domain/changemailboxpassword", "POST"}
var getMailListAction = action{"/domain/getmaillist", "GET"}

//DATABASE ACTIONS
var addDatabaseAction = action{"/domain/adddatabase", "POST"}
var deleteDatabaseAction = action{"/domain/deletedatabase", "DELETE"}
var addDatabaseUserAction = action{"/domain/adddatabaseuser", "POST"}
var deleteDatabaseUserAction = action{"/domain/deletedatabaseuser", "DELETE"}
var changeDatabaseUserPasswordAction = action{"/domain/changedatabaseuserpassword", "POST"}
var getDatabaseListAction = action{"/domain/getdatabaselist", "GET"}
var setDatabaseQuotaAction = action{"/domain/setdatabasequota", "POST"}
var setDatabaseUserPermissions = action{"/domain/setdatabaseuserpermissions", "POST"}

//FTP ACTIONS
var addFtpAccountAction = action{"/domain/addftpaccount", "POST"}
var deleteFtpAccountAction = action{"/domain/deleteftpaccount", "DELETE"}
var changeFtpAccountPasswordAction = action{"/domain/changeftppassword", "POST"}
var getFtpAccountsAction = action{"/domain/getftpaccounts", "GET"}

//DNS ACTIONS
var setDNSZoneAction = action{"/domain/setdnszone", "POST"}
var addDNSRecordAction = action{"/domain/adddnsrecord", "POST"}
var deleteDNSRecordAction = action{"/domain/deletednsrecord", "DELETE"}
var getDNSRecordsAction = action{"/domain/getdnsrecords", "GET"}

//SSL ACTIONS
var createSSLRequestAction = action{"/domain/createsslrequest", "POST"}
var getSSLAction = action{"/domain/getsslcert", "GET"}
var deleteSSLAction = action{"/domain/deletesslrequest", "DELETE"}
var completeSSLAction = action{"/domain/completesslrequest", "POST"}
