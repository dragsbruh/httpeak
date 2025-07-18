# httpeak

i use this on my homelab to debug caddy or cloudflare or other transport stuff.
just visit it in your browser and all good.

all it does is return the entire http request (?) as plaintext body.

```
X-Forwarded-For: <stuff>
X-Forwarded-Proto: http
Accept-Encoding: gzip, br
Accept-Language: en-US,en;q=0.5
Cf-Connecting-Ip: <stuff>
Cf-Ray: <stuff>
Cf-Visitor: {"scheme":"https"}
Sec-Fetch-Mode: navigate
Upgrade-Insecure-Requests: 1
Via: 1.1 Caddy
User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:140.0) Gecko/20100101 Firefox/140.0 (arch btw)
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Cdn-Loop: cloudflare; loops=1
Cf-Warp-Tag-Id: <stuff>
X-Forwarded-Host: echo.vinster.xyz
X-Real-Ip: <stuff>
Alt-Used: echo.vinster.xyz
Priority: u=0, i
Sec-Fetch-Dest: document
Cf-Ipcountry: IN
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Sec-Gpc: 1
```

## TODO

- [ ] moar optimizations
