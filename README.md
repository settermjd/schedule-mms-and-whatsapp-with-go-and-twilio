# Schedule MMS and WhatsApp with Go and Twilio

This is a small repository that shows how to schedule an MMS and a WhatsApp message with Go and [Twilio's Programmable Messaging API](https://www.twilio.com/docs/sms/api/message-resource#schedule-a-message-resource).
It supports an upcoming Twilio tutorial.

## Requirements/Prerequisites

You will need the following to use this codebase:

- A Twilio account, either free or paid. 
  If you are new to Twilio, [click here to create a free account](http://www.twilio.com/referral/QlBtVJ)
- A Twilio phone number and [Messaging Service SID](https://support.twilio.com/hc/en-us/articles/223134387-What-is-a-Message-SID-)
- A mobile phone number that can receive MMS
- [WhatsApp](https://www.whatsapp.com/), whether on your mobile phone or desktop
- A recent version of [Go](https://go.dev/)
- Some prior experience with Go would be beneficial, though isn't really necessary

## Usage

To use the code, first, clone it locally, then copy _.env.example_ as _.env_ and set the required variables in place of the placeholders.
Specifically, after logging in to [the Twilio Console](https://console.twilio.com):

- Retrieve your **Twilio Auth Token**, **Account SID**, and **phone number** from the dashboard and set them in place of `TWILIO_AUTH_TOKEN`, `TWILIO_ACCOUNT_SID`, and `TWILIO_PHONE_NUMBER`. 
- Retrieve your Twilio Messaging Service SID, by navigating to **Explore Products > Messaging >** [Services](https://console.twilio.com/us1/develop/sms/services). 
  There, copy the value of the `Sid` column of your Messaging Service. 
  Set the copied value as the value of `TWILIO_MESSAGING_SERVICE_SID`.
- Navigate to **Explore Products > Messaging > Try it out >** [Send a WhatsApp Message](https://console.twilio.com/us1/develop/sms/try-it-out/whatsapp-learn?frameUrl=%2Fconsole%2Fsms%2Fwhatsapp%2Flearn%3Fx-target-region%3Dus1). 
  Follow the instructions under **Connect to WhatsApp Sandbox**. then, copy the contents of the `To` and `From` fields as the values of `TWILIO_WHATSAPP_NUMBER` and `RECIPIENT_WHATSAPP_NUMBER` respectively.
- Finally, replace `RECIPIENT_PHONE_NUMBER` with your mobile phone number. 
  Make sure that it is [E.164-formatted](https://www.twilio.com/docs/glossary/what-e164)!
  
### Update the scheduling delay

By default, messages are scheduled for 17 minutes from when the code is run. 
To change this, update the following line in *main.go*, changing `17m` to [the duration](https://pkg.go.dev/maze.io/x/duration#ParseDuration) that you prefer:

```go
duration, err := time.ParseDuration("17m")
```

### Schedule an MMS

Run the following command to run the code and schedule an MMS using Twilio's Programmable Messaging API:

```bash
go run main.go
```

### Schedule a WhatsApp Message

To schedule a WhatsApp message instead of an MMS, change the following line in main.go...

```go
params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
params.SetTo(os.Getenv("RECIPIENT_PHONE_NUMBER"))
```

...to the following:

```go
params.SetFrom(os.Getenv("TWILIO_WHATSAPP_NUMBER"))
params.SetTo(os.Getenv("RECIPIENT_WHATSAPP_NUMBER"))
```

Then, run the code as before.

## Got Stuck?

If you have any issues running the code, please [create an issue](https://github.com/settermjd/schedule-mms-and-whatsapp-with-go-and-twilio/issues/new/choose).  