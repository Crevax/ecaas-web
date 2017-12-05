# ecaas-web
Example Go Web API using the [Ecaas Go SDK](https://github.com/syllabix/ecaas)

## Running on IIS
The simplest way to run a Go application through IIS is with the [HttpPlatformHandler](https://www.iis.net/downloads/microsoft/httpplatformhandler) module.

Once that is installed, the path for the website should contain a `web.config` file similar to the following:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
    <system.webServer>
        <handlers>
            <add name="httpplatformhandler" path="*" verb="*" modules="httpPlatformHandler" resourceType="Unspecified" />
        </handlers>
        <httpPlatform processPath=".\app\ecaas-web.exe"
                      arguments=""
                      startupTimeLimit="60"
                      stdoutLogEnabled="true"
                      stdoutLogFile="c:\sitelogs\ecaas-web.txt">
            <environmentVariables>
                <environmentVariable name="FOO" value="BAR" />
            </environmentVariables>
        </httpPlatform>>
        <httpProtocol>
            <customHeaders>
                <remove name="X-Powered-By" />
            </customHeaders>
        </httpProtocol>
    </system.webServer>
</configuration>

```

N.B. The `system.webServer/handlers` section needs to be unlocked within the site's configuration. And it must be done before the `web.config` file is read from. IIS won't allow that change to be made while the block exists.