# Auto-Generated Swagger Bindings to `Uber`

The library in `lib` provides auto-generated-from-Swagger bindings to the Uber API.

## Installation

Installation follows the standard approach to installing Stack-based projects.

1. Install the [Haskell `stack` tool](http://docs.haskellstack.org/en/stable/README).
2. Run `stack install` to install this package.

## Main Interface

The main interface to this library is in the `Uber.API` module, which exports the UberBackend type. The UberBackend
type can be used to create and define servers and clients for the API.

## Creating a Client

A client can be created via the `createUberClient` function, which, if provided with a hostname and a port, will generate
a client that can be used to access the API if it is being served at that hostname / port combination. For example, if
`localhost:8080` is serving the Uber API, you can write:

```haskell
{-# LANGUAGE RecordWildCards #-}

import Uber.API

main :: IO ()
main = do
  UberBackend{..} <- createUberClient (ServerConfig "localhost" 8080)
  -- Any Uber API call can go here.
  return ()
```

## Creating a Server

In order to create a server, you must use the `runUberServer` function. However, you unlike the client, in which case you *got* a `UberBackend`
from the library, you must instead *provide* a `UberBackend`. For example, if you have defined handler functions for all the
functions in `Uber.Handlers`, you can write:

```haskell
{-# LANGUAGE RecordWildCards #-}

import Uber.API

-- A module you wrote yourself, containing all handlers needed for the UberBackend type.
import Uber.Handlers

-- Run a Uber server on localhost:8080
main :: IO ()
main = do
  let server = UberBackend{..}
  runUberServer (ServerConfig "localhost" 8080) server
```

You could use `optparse-applicative` or a similar library to read the host and port from command-line arguments:
```
{-# LANGUAGE RecordWildCards #-}

module Main (main) where

import Uber.API (runUberServer, UberBackend(..), ServerConfig(..))

import Control.Applicative ((<$>), (<*>))
import Options.Applicative (execParser, option, str, auto, long, metavar, help)

main :: IO ()
main = do
  config <- parseArguments
  runUberServer config UberBackend{}

-- | Parse host and port from the command line arguments.
parseArguments :: IO ServerConfig
parseArguments =
  execParser $
    ServerConfig
      <$> option str  (long "host" <> metavar "HOST" <> help "Host to serve on")
      <*> option auto (long "port" <> metavar "PORT" <> help "Port to serve on")
```
