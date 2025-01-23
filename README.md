# AlloPopoT Email Notifier Service
A service to dispatch notification emails from any internal service via RabbitMQ (amqp).
Requires RabbitMQ to be installed.

## Build the Project
To build the project run:

    go build

## Message Structure
The message is a stringified JSON.

    {
        "to": [
            "<sender-email-address>"
        ],
        "subject": "<subject>",
        "body": "<body-in-html-or-plaintext>",
        "attachments": [
            {
                "filename": "<filename-with-extension>",
                "mimetype": "<file-mimetype>",
                "payload": "<file-payload-in-base64>"
            }
        ]
    }

## Environment Variables
| Name               | Required | Default                |
| ------------------ | -------- | ---------------------- |
| AMQP_HOST          | true     | -                      |
| AMQP_PORT          | false    | 5672                   |
| AMQP_USERNAME      | true     | -                      |
| AMQP_PASSWORD      | true     | -                      |
| AMQP_EXCHANGE_NAME | false    | EMAIL-SERVICE-EXCHANGE |
| SMTP_HOST          | true     | -                      |
| SMTP_PORT          | true     | -                      |
| SMTP_USERNAME      | true     | -                      |
| SMTP_PASSWORD      | true     | -                      |
| SMTP_SENDER        | true     | -                      |