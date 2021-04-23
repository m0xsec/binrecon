# binrecon

```
 ▄▄▄▄    ██▓ ███▄    █  ██▀███  ▓█████  ▄████▄   ▒█████   ███▄    █ 
▓█████▄ ▓██▒ ██ ▀█   █ ▓██ ▒ ██▒▓█   ▀ ▒██▀ ▀█  ▒██▒  ██▒ ██ ▀█   █ 
▒██▒ ▄██▒██▒▓██  ▀█ ██▒▓██ ░▄█ ▒▒███   ▒▓█    ▄ ▒██░  ██▒▓██  ▀█ ██▒
▒██░█▀  ░██░▓██▒  ▐▌██▒▒██▀▀█▄  ▒▓█  ▄ ▒▓▓▄ ▄██▒▒██   ██░▓██▒  ▐▌██▒
░▓█  ▀█▓░██░▒██░   ▓██░░██▓ ▒██▒░▒████▒▒ ▓███▀ ░░ ████▓▒░▒██░   ▓██░
░▒▓███▀▒░▓  ░ ▒░   ▒ ▒ ░ ▒▓ ░▒▓░░░ ▒░ ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒ 
▒░▒   ░  ▒ ░░ ░░   ░ ▒░  ░▒ ░ ▒░ ░ ░  ░  ░  ▒     ░ ▒ ▒░ ░ ░░   ░ ▒░
 ░    ░  ▒ ░   ░   ░ ░   ░░   ░    ░   ░        ░ ░ ░ ▒     ░   ░ ░ 
 ░       ░           ░    ░        ░  ░░ ░          ░ ░           ░ 
      ░                                ░                            
```

`binrecon` is a simple Pastebin collections tool written in Go and Elastic stack. It assumes you have [Pastebin Scraping API](https://pastebin.com/doc_scraping_api) access (which is based on IPv4 or IPv6 association). 

## Configuration

The `configs` folder contains the relevant configuration files for Elastic search, Kibana, Filebeat, and Logstash. These configuration files will be used by Docker.

## Usage

This is intended to be ran via Docker though you may also deploy it elsewhere.

If you have any questions, comments, or concerns, reach out to me on Twitter [@m0xxz](https://twitter.com/m0xxz) <3