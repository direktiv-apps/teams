
# teams 1.0

Send messages to Teams

---
- #### Categories: social
- #### Image: gcr.io/direktiv/functions/teams 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/teams/issues
- #### URL: https://github.com/direktiv-apps/teams
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About teams

This function can send messages and adaptove cards to Microsoft Teams via incoming webhook URL.

To create an incoming webhook URL go to `Manage Channel` in Teams. In the `Connector` section select `Incoming Webhook` and create a new webhook.  This URL is used as Direktiv secret to send the message to this channel.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: teams
  image: gcr.io/direktiv/functions/teams:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: teams
  type: action
  action:
    function: teams
    secrets: ["webhook-url"]
    input: 
      webhook-url: jq(.secrets."webhook-url")
      verbose: true
      content:
        text: Hello World!
```
   #### Adaptive card and verbose output
```yaml
- id: teams
  type: action
  action:
    function: teams
    secrets: ["webhook-url"]
    input: 
      webhook-url: jq(.secrets."webhook-url")
      verbose: true
      content:
        type: message
        attachments:
        - contentType: application/vnd.microsoft.card.adaptive
          content:
            "$schema": http://adaptivecards.io/schemas/adaptive-card.json
            type: AdaptiveCard
            version: '1.2'
            body:
            - type: TextBlock
              text: Textblock1
            - type: TextBlock
              text: 2 Stops Test
              weight: bolder
              spacing: medium
```

   ### Secrets


- **webhook-url**: Incoming webhook URL from Teams






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Result of teams message command
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": 1,
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| teams | [PostOKBodyTeams](#post-o-k-body-teams)| `PostOKBodyTeams` |  | |  |  |


#### <span id="post-o-k-body-teams"></span> postOKBodyTeams

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | integer| `int64` |  | |  |  |
| success | boolean| `bool` |  | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | [interface{}](#interface)| `interface{}` |  | |  |  |
| verbose | boolean| `bool` |  | | Enables verbose output |  |
| webhook-url | string| `string` |  | | URL for Teams's incoming webhook |  |

 
