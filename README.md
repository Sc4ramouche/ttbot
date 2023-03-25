# ttbot

Experimenting with golang and telegram bot API.

### Notes

- Web scraping can be done quite easily with tools like [Colly](https://github.com/gocolly/colly).
  - There are lots of [official examples](https://github.com/gocolly/colly/tree/master/_examples) which are super handy for those who is neither experienced in golang nor web scraping.
  - Another solid open-source module that can be used for scraping is [goquery](https://github.com/PuerkitoBio/goquery).
- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) helped me to quickly test and launch first real integration.

### Todos

- At the current stage, this module is useless as it reads telgram user ID from the env variable:
  - I could store user IDs in a database, thus unhardcoding the bot.
  - [SQLite](https://sqlite.org/index.html) seems like a good candidate to try.
- Introduce the logic of checking if the "announcement" message has already been sent this week.
  - Q: could that be done by persisting sent messages in DB too, together with the time when they were sent?
- To really enhance my learning with this project, I could do the following:
  - Host the server on the web.
    - Q: what is a best practice when deploying go servers – uploading the compiled executable?
    - Q: what could be a good server provider to play around?
    - Q: How do I make my program to only run, say, every 2 hours on the remote machine?
  - Build a CI and explore what are the conventional tools in the golang world.
  - Host the database on the web.
    - Q: would it be alright to host both server and DB on the same machine for such a tiny project?
      - [Stack Overflow thread](https://stackoverflow.com/questions/659970/why-is-it-not-advisable-to-have-the-database-and-web-server-on-the-same-machine) on the topic.
