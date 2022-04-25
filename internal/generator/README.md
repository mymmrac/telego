# Telego's generator

Telego has all methods and types as they are described in [Telegram Bot API](https://core.telegram.org/bots/api), all
these are generated from docs themselves. All content of the doc's page is parsed using regular expressions, divided 
into methods and types, and used for generation.

Also, `With...` methods are generated from types and most of the tests are generated too. These tests meant to be
regenerated only after previous tests pass in the new version of the library have been generated.

The process of generation is pretty simple:
1. Get raw HTML of docs
2. Parse blocks of HTML with tables, where types and methods defined
3. Parse resulted blocks into separate types/methods with individual parameters/fields
4. Generate Go code from parsed data
5. Write it into files and run `go fmt` on them

## Disclaimer

This code (code of generator itself) isn't meant to be reused by anyone (at least directly) and has quite high
complexity so there are no tests or linters for it, and it is located inside the `internal` folder, so be very 
thoughtful if you want to use it. 
