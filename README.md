# Slack To Discord Webhook
 forward a slack webhook to discord

 
![image](https://github.com/PaulLocksley/Slack-To-Discord-Webhook/assets/96610143/04e195c0-4459-4526-8b23-a448e71ce673)


So.... you can just add slack to the end of a webhook and discord handles it.....

Anyway.
Super basic, just built to handle truenas alerts. 

if you actually use it simply deploy the related package with the following:
- env var s2dwebhook set to your discord webhook.
- container port 8080 exposed 
