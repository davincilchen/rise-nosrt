# rise-nostr

## Usage

* Run docker-compose to start db
* Run relay project at  ws://127.0.0.1:8100/
* Client will connect to ws://127.0.0.1:8100/ for default relay

## API

* POST   API by http://{{serve}}:8000/event
* POST   API by http://{{serve}}:8000/req
* DELETE API by http://{{serve}}:8000/req

Example of API body when POST http://{{serve}}:8000/event  

    {
        "pubKey":"1234567894b6881f61ab5116a52c72e161af9cdca5ca8fd59f296c3d94e8532a",
        "priKey":"1234567897de8e9ac5afcfc3eea195f21aa9940daa09614c5dd59260c1a77812",
        "msg":"good night"
    }

Use dafault key pairs when pubKey or priKey field in API is empty.

可以啟動多個client,
每個client預設都會執行一次REQ CMD

## Question Exercises

What are some of the challenges you faced while working on Phase 1?

* Lose something when read nostr protocol and assigment. So spend a lot time to fixed some bug.

* Try to find how to sign form 64 hex private key.

* Because of not use Nostr lib, spend a lot of time.

What kind of failures do you expect to a project such as DISTRISE to encounter?

* Lose some event if we need to aggregate a large number of user's events.

  * What kind of architecture is required when we need to aggregate a large number of user events. How to avoid system can't not work fine.

* Input speed greater than processing speed

  * How to coordinate multiple compute engines to work normally.

  * How to allocate compute engine subscribed their own objects evenly and dynamically

  * How to reduce repeat event if need.

* Input speed greater than store speed

  * How to store a large number of user's events and does not cause IO bottleneck.

Others

* Security of transferring private key. (TODO)

* Store the private key locally if only local users are supported. (Done)

* Retry if disconnect.  (TO Check)
  
  * Error when multi-client req (TODO)
