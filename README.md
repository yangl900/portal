# Portal
A simple tool that parses and converts stdin to clickable Azure portal hyperlink. Enable easier navigation between Azure CLI and Portal.

# How to use
Print the deep link url for my-vm. In cloud shell, this will open a new tab automatically with the deep link.
```
az vm show -g my-group -n my-vm | portal
```

Print all deep link urls of VMs in the subscription
```
az vm list | jq -r .[].id | portal
```

# How to install
For Linux:
```bash
curl -sL https://github.com/yangl900/portal/releases/download/0.1/portal_linux_64-bit.tar.gz | tar xz
```

For Windows (In PowerShell):
```powershell
curl https://github.com/yangl900/portal/releases/download/0.1/portal_windows_64-bit.zip -OutFile portal.zip
```
And unzip the file, the only binary needed is portal.exe.

For MacOS:
```bash
curl -sL https://github.com/yangl900/portal/releases/download/0.1/portal_macOS_64-bit.tar.gz | tar xz
```

# Azure Cloud Shell
It works in cloud shell beautifully.

[![Launch Cloud Shell](https://shell.azure.com/images/launchcloudshell.png "Launch Cloud Shell")](https://shell.azure.com)