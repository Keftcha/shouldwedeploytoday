# Should we deploy today ?

API and web page inspired by [estcequonmetenprodaujourdhui](https://www.estcequonmetenprodaujourdhui.info/).

## Server

The server listen on port `5461` on all interfaces.

### API

The api route is `/api/{weekday}`.
You should replace `{weekday}` by an integer that is the current day (ex: `/api/2`).  
0 is for Sunday, 1 for Monday, ...

The server will return you a json like this (for Tuesday):

```json
{
    "title": "On Tuesday ?",
    "image": "/images/sounds_good_to_me.gif",
    "text": "Let's go !",
    "is_a_deployment_day": true
}
```

### Web page

At the root URL you'll get a web page that tell you if you should deploy today or not.
