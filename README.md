# opengo

Automatically opens the url with the given list.

## Install

```
▶ go install github.com/alfiannurrizky/opengo@latest
```

## Basic Usage

opengo accepts line-delimited url on `stdin`:

```
▶ cat url.txt
http://example.com
http://example.net
http://example.edu

▶ cat url.txt | opengo
will automatically open the browser
```

## Extra

By default opengo will open all urls with the given list.
If you just want to open some urls, use `-t` flag.

```
▶ cat url.txt | opengo -t 10
```

## Tips

You can combo with [httpx](https://github.com/projectdiscovery/httpx) tool, like this.

```
▶ cat subdomain.txt
products.example.com
support.example.com
www.example.com

▶ cat subdomain.txt | httpx -sc -mc 200 | opengo -t 5
This will open a site whose status code is 200
```
