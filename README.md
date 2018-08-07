# OpsGenie Go Integration

## Requirements

Go

````
brew install go
```

```
go get github.com/weberling/opsgenie-go-integration
```

Get the key on MBACk_Toolkit
```
export OPSGENIE_API_KEY={opsgenie_api_key}
``` 

## Run

Go to user root folder
cd ~/go/src/github.com/weberling/opsgenie-go-integration
git checkout easy-runnable

### List alerts

List alerts in an interval using the opsgenie createdAt attribute and generate a csv file **report.csv**

```
go run main.go list-alerts {start_date} {end_date}(optional)
```

**pattern date**: yyyy-MM-dd. *Ex: 2018-07-10*

**interval rule**: >= start_date AND < end_date