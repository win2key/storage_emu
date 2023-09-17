package main

import (
	"fmt"
)

func cmd_help() {
	fmt.Println(`addhostport - Adds worldwide port names (WWPNs) or iSCSI names to a host object.
applysoftware - Updates the system code.
cancellivedump - Cancels a livedump.
catauditlog - Displays the in-memory contents of the audit log.
charray - Changes the attributes of an array.
charraymember - Changes an array member's attributes, or swaps a member of a RAID array.
chauthservice - Configures the remote authentication service of the system.
chbanner - Changes the login banner message.
chcurrentuser - Changes the attributes of the current user.
chdnsserver - Modify a DNS server.
chdrive - Change the attributes of a drive.
chemail - Sets or changes contact information for email event notifications.
chemailserver - Changes the attributes of an email server.
chemailuser - Changes the attributes for an email recipient.
chenclosure - Changes enclosure properties.
chenclosurecanister - Changes the properties of an enclosure canister.
chenclosurepsu - Internal use only command.
chenclosureslot - Changes the properties of an enclosure slot.
chencryption - Changes encryption settings.
cherrstate - Marks an unfixed error as fixed or a fixed error as unfixed.
cheventlog - Changes events in the event log.
chhost - Changes the attributes of a host object.
chiogrp - Changes the name and other attributes of an I/O group.
chkeyserver - Changes the attributes of a key server object.
chkeyserverisklm - Changes the system-wide IBM Security Key Lifecycle Manager (ISKLM) key server configuration.
chldap - Changes the system-wide Lightweight Directory Access Protocol (LDAP) configuration.
chldapserver - Changes a Lightweight Directory Access Protocol (LDAP) server.
chnodecanister - Changes the name and other attributes of a node canister.
chnodecanisterhw - Updates the hardware configuration for a node canister.
chportfc - Changes properties of an FC port.
chportib - Changes properties of an IB port.
chportip - Changes properties of an IP port.
chsecurity - Changes security settings.
chsnmpagent - Changes the settings for a Simple Network Management Protocol (SNMP) agent server.
chsnmpserver - Changes the attributes of an SNMP server.
chsyslogserver - Changes the attributes of a syslog server.
chsystem - Changes the attributes of an existing system.
chsystemcert - Modify the system SSL certificate.
chsystemip - Changes the IP configuration parameters for the system.
chsystemsupportcenter - Changes properties of a Support Center server or proxy.
chuser - Changes the attributes of a user.
chusergrp - Changes the attributes of a user group.
chvdisk - Changes the attributes of a volume.
cleardumps - Cleans the dump directories on a specified node canister.
clearerrlog - Clears all entries from the event log.
cpdumps - Copies dump files to the configuration node canister from another node canister.
dumpauditlog - Resets or clears the contents of the in-memory audit log.
dumperrlog - Dumps the contents of the event log to a text file.
enablecli - Enables the CLI for a recovered system.
expandvdisksize - Expands the size of a volume by a given capacity.
finderr - Finds the highest severity unfixed error in the event log.
lsarray - Displays array managed disks
lsarrayinitprogress - Displays the progress of array background initialization.
lsarraymember - Displays the member drives of one or more array managed disks.
lsarraymemberprogress - Displays array member background process status.
lscopystatus - Displays whether any file copies are currently in progress.
lscurrentuser - Displays the name and role of the logged-in user.
lsdependentvdisks - Displays which volumes will go offline if you remove a piece of hardware from the system.
lsdnsserver - Displays the DNS servers configured on the system.
lsdrive - Displays configuration information and drive VPD.
lsdriveclass - Displays drive class information
lsdriveprogress - Displays the progress of various drive tasks.
lsdumps - Displays files in a dumps directory on one of the node canisters in the system.
lsemailserver - Displays the email servers configured on the system.
lsemailuser - Displays email event notification settings for email recipients.
lsenclosure - Displays a summary of the enclosures.
lsenclosurebattery - Displays information about the batteries in the enclosure PSUs.
lsenclosurecanister - Displays status for each canister in an enclosure.
lsenclosurechassis - Displays information about the enclosure chassis.
lsenclosurepsu - Displays information about each power-supply unit (PSU) in the enclosure.
lsenclosureslot - Displays information about each drive slot in the enclosure.
lsenclosurestats - Displays enclosure statistics.
lsencryption - Displays encryption information.
lseventlog - Displays system event log information.
lsfabric840 - Displays the port mapping for FlashSystem 840
lsfcportcandidate - Displays unconfigured, logged-in SAS host bus adapter (HBA) ports.
lshbaportcandidate - Displays unconfigured, logged-in host bus adapter (HBA) ports.
lshost - Displays the hosts visible to the system.
lshostiogrp - Displays the I/O groups associated with a specified host.
lshostvdiskmap - Displays volumes that are mapped to a given host.
lsibportcandidate - Displays unconfigured, logged-in Infiniband host bus adapter (HBA) ports.
lsiogrp - Displays the I/O groups on the system.
lsiogrpcandidate - Displays I/O groups that can have node canisters added to them.
lsiogrphost - Displays the hosts mapped to an I/O group.
lsiscsiauth - Displays the Challenge Handshake Authentication Protocol (CHAP) secret configured for authenticating an entity to the system.
lskeyserver - Displays the key servers on the system.
lskeyserverisklm - Displays the system-wide IBM Security Key Lifecycle Manager (ISKLM) key server configuration.
lsldap - Displays the system-wide Lightweight Directory Access Protocol (LDAP) configuration.
lsldapserver - Displays all configured Lightweight Directory Access Protocol (LDAP) servers.
lslicense - Displays the current license settings for system features.
lslivedump - Displays the livedump state of a node canister.
lsmdisk - Displays the managed disks visible to the system.
lsmdiskgrp - Displays the storage pools on the system.
lsnodecanister - Displays the node canisters that are part of the system.
lsnodecanisterhw - Displays the hardware configuration of node canisters in the system.
lsnodecanisterstats - Displays node canister statistics.
lsnodecanistervpd - Displays the vital product data (VPD) for each node canister.
lspartnership - Displays the current systems partnered with the local system.
lsportfc - Displays information about the Fibre Channel I/O ports.
lsportib - Displays information about the InfiniBand ports.
lsportip - Displays the iSCSI IP addresses assigned for each port on each node canister in the system.
lsportsas - Displays information about the SAS ports.
lsportusb - Displays information about the USB ports used for encryption.
lsquorum - Displays the managed disks or drives the system is currently using to store quorum data.
lsrepairsevdiskcopyprogress - Displays the repair progress for space-efficient volume copies.
lsrepairvdiskcopyprogress - Displays the progress of volume repairs and validations.
lsrmvdiskdependentmaps - Displays all mappings that must be stopped for a volume to be deleted.
lsroute - Displays the IP routing table.
lssasportcandidate - Displays unconfigured, logged-in FC host bus adapter (HBA) ports.
lssecurity - Internal use only.
lssevdiskcopy - Displays the space-efficient copies of volumes.
lssnmpagent - Displays the Simple Network Management Protocol (SNMP) agent server configured on the system.
lssnmpserver - Displays the SNMP servers configured on the system.
lssoftwareupgradestatus - Displays the status of a code update.
lssra - Displays the secure remote access (SRA) status and the time of the last SRA login.
lssyslogserver - Displays the syslog servers configured on the system.
lssystem - Displays the local system.
lssystemcert - Displays the system certificate.
lssystemip - Displays the system management IP addresses configured for each port.
lssystemstats - Displays the most recent values or a history of canister statistics.
lssystemsupportcenter - Displays the Support Center or proxy servers configured on the system.
lstimezones - Displays the time zones that are available on the system.
lsupdate - Displays information about system, enclosure, and drive update.
lsuser - Displays the users created on the system.
lsusergrp - Displays the user groups created on the system.
lsvdisk - Displays the volumes on the system.
lsvdiskaccess - Displays which IO groups volumes are accessible through.
lsvdiskanalysis - List single vdisk or multiple vdisks thin provisioning and compression estimation analysis report.
lsvdiskanalysisprogress - List overall space analysis progress for cluster.
lsvdiskcopy - Displays volume copy information.
lsvdiskhostmap - Displays volume to host mappings.
lsvdiskprogress - Displays the progress during new volume formatting.
lsvdisksyncprogress - Displays the progress of volume copy synchronization.
mkarray - Creates a managed disk array and adds it to a storage pool.
mkdnsserver - Creates a DNS server.
mkemailserver - Creates an email server.
mkemailuser - Adds an email recipient of event and inventory notifications.
mkhost - Creates a logical host object.
mkkeyserver - Creates a key server.
mkldapserver - Creates a Lightweight Directory Access Protocol (LDAP) server.
mksnmpserver - Creates an SNMP server to receive notifications.
mksyslogserver - Creates a syslog server to receive notifications.
mksystemsupportcenter - Creates a new Support Center server or proxy.
mkuser - Creates a local or a remote user to access a system.
mkusergrp - Creates a user group.
mkvdisk - Creates a sequential, striped, or image mode volume.
mkvdiskhostmap - Creates a mapping between a volume and a host.
ping - Diagnoses IP configuration.
preplivedump - Reserves system resources required for a livedump.
recoverarray - Recovers a corrupt array.
resetleds - Simultaneously switches off all LEDs in the system.
rmarray - Removes a managed disk array from the configuration.
rmdnsserver - Deletes a DNS server.
rmemailserver - Deletes an email server.
rmemailuser - Remove an email recipient from your system.
rmhost - Deletes a host object.
rmhostport - Deletes worldwide port names (WWPNs) or iSCSI names from a host object.
rmkeyserver - Deletes a key server object.
rmldapserver - Deletes a Lightweight Directory Access Protocol (LDAP) server.
rmnodecanister - Deletes a node canister from the system.
rmportip_tms - Removes properties of an IP port.
rmsnmpserver - Delete an SNMP server.
rmsyslogserver - Deletes a syslog server.
rmsystemsupportcenter - Deletes a Support Center server or proxy.
rmuser - Deletes a user.
rmusergrp - Deletes a user group.
rmvdisk - Deletes a volume.
rmvdiskhostmap - Deletes a volume-to-host mapping.
sainfo - Displays service action views.
satask - Performs service action tasks.
sendinventoryemail - Sends an inventory email to all recipients enabled to receive inventory email notifications.
setlocale - Changes the locale setting and command output to the chosen language.
setpwdreset - Views or changes the status of the password-reset feature.
setsystemtime - Sets the time for the system.
settimezone - Sets the time zone for the system.
showtimezone - Displays the current time zone settings for the system.
shrinkvdisksize - Reduces the size of a volume by the specified capacity.
startemail - Activates the email and inventory notification function.
stopemail - Stops the email and inventory notification function.
stopsystem - Shuts down a single node canister or the entire system.
svc_livedump - Triggers a livedump of the system.
svc_snap - Collects log files used by support for troubleshooting.
svcconfig - Recovers, restores, or creates a backup of the system configuration.
testemail - Sends a test email notification to one user or to all users of the email notification function.
testkeyserver - Tests a key server.
testldapserver - Tests a Lightweight Directory Access Protocol (LDAP) server.
triggerlivedump - Captures a livedump to the internal disk on the node canister.`)
}
