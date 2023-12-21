# Safe Trader written in Golang for Kraken API

## What is the purpose of this bot

I have started plenty of bots and all of them felt simple when I started.
This one is no exception, but I am limiting even more of what I want the bot to do.

So here is the idea:

I set the sell price of the token and the quantity and then the bot does:

- Confirms in slack that I have set a sell
- Asks if I want to set up a safety
  - If so it asks what the _safety_ should be and gives me defaults
- Asks if I want to buy back
  - If so it asks what the _target_ is
  - If it should buy back the _same number of coins_ or _max_

Now the bot waits till the token is (partly) sold. Once the sale is started:

- Creates a stop loss limit buy for all the euro's the trade makes
- Starts monitoring for the target buy price
- If we are below sell
- We remove the safety
- We reach the target price we create a buy limit for the number of coins
