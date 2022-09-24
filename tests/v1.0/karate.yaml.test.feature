
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def webhookurl = karate.properties['webhookurl']


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"webhook-url": #(webhookurl),
		"verbose": true,
		"content": {
			"type": "message",
			"attachments": [
				{
					"contentType": "application/vnd.microsoft.card.adaptive",
					"content": {
						"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
						"type": "AdaptiveCard",
						"version": "1.2",
						"body": [
							{
								"type": "TextBlock",
								"text": "Hello World!"
							},
												{
													"type": "TextBlock",
													"text": "2 Stops Test",
													"weight": "bolder",
													"spacing": "medium"
												}
						]
					}
				}
			]
		}
	}
	"""
	When method POST
	Then status 200
	# And match $ ==
	# """
	# {
	# "teams": [
	# {
	# 	"result": "#notnull",
	# 	"success": true
	# }
	# ]
	# }
	# """
	