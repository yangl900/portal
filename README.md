# portal
Simple tool to convert stdin to Azure portal hyperlink. Enable easier navigation between CLI and Portal.

# How to use
```
az vm show -g my-group -n my-vm | portal
```

```
az vm list | jq .[].id | portal
```