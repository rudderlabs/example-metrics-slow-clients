# example-metrics-slow-clients

Code examples of how slow clients can effect HTTP metrics.

## Word Count Service

`wc` is a simple HTTP service that response with the word count of the request body.

run:

```bash
    go run ./wc
```

and test the service with:

```bash
    curl -d "@README.md" localhost:8080
```

Also, notice that the service logs the amount of time it took for a service to be processed.

## Slow Client

We can simulate a slow client by passing `--limit-rate` option to the curl command.

> --limit-rate <speed>
> Specify the maximum transfer rate you want curl to use.
> This feature is useful  if you  have a limited pipe and
> you'd like your transfer not to use your entire bandwidth.

> The given speed is measured in bytes/second, unless a suffix
> is appended.  Appending  'k'  or 'K' will count the number
> as kilobytes, 'm' or M' makes it megabytes, while 'g' or 'G'
> makes it gigabytes. Examples: 200K, 3m and 1G.

Let's try to simulate a really slow one:

```bash
    curl -d "@README.md" localhost:8080 --limit-rate 60
```

And notice what happens with measurement logs.

## Buffered Body Middleware

A solution to avoid measuring the time it takes to read body, is to use body buffering middleware.

Try the slow client against this implementation:

```bash
    go run ./wc-buffered
```
