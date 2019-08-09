# Github Repo Searcher

Given a query, it'll return the repositories where the code source matches the
query.

## Usage

### Create a Github token (personal or oauth)

https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line
Then you can export the token via a env variable called `TOKEN`:

`.token.sh`
``` sh
#!/usr/bin/env sh
export TOKEN=mytoken
```

### Build the tool
``` sh
make build
```

### Run

``` sh
source .token.sh && github-searcher --query "sarama language:go org:Shopify"
```

or you can also provide the token via parameters, but it'll be in your commands
history ¯\_(ツ)_/¯

``` sh
source .token.sh && ./github-searcher --query "sarama language:go org:Shopify"
--token mytoken
```
