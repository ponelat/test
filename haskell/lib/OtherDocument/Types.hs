{-# LANGUAGE DeriveGeneric #-}
{-# LANGUAGE GeneralizedNewtypeDeriving #-}
{-# OPTIONS_GHC -fno-warn-unused-binds -fno-warn-unused-imports #-}

module OtherDocument.Types (
    ) where

import Data.List (stripPrefix)
import Data.Maybe (fromMaybe)
import Data.Aeson (Value, FromJSON(..), ToJSON(..), genericToJSON, genericParseJSON)
import Data.Aeson.Types (Options(..), defaultOptions)
import Data.Text (Text)
import qualified Data.Text as T
import qualified Data.Map as Map
import GHC.Generics (Generic)
import Data.Function ((&))


-- Remove a field label prefix during JSON parsing.
-- Also perform any replacements for special characters.
removeFieldLabelPrefix :: Bool -> String -> Options
removeFieldLabelPrefix forParsing prefix =
  defaultOptions
    { fieldLabelModifier = fromMaybe (error ("did not find prefix " ++ prefix)) . stripPrefix prefix . replaceSpecialChars
    }
  where
    replaceSpecialChars field = foldl (&) field (map mkCharReplacement specialChars)
    specialChars = [("@", "'At"), ("!", "'Exclamation"), ("<=", "'Less_Than_Or_Equal_To"), ("#", "'Hash"), ("$", "'Dollar"), ("%", "'Percent"), ("&", "'Ampersand"), ("*", "'Star"), ("+", "'Plus"), ("-", "'Dash"), (".", "'Period"), (":", "'Colon"), ("|", "'Pipe"), ("<", "'LessThan"), ("!=", "'Not_Equal"), ("=", "'Equal"), ("^", "'Caret"), (">", "'GreaterThan"), ("_", "'Underscore"), (">=", "'Greater_Than_Or_Equal_To")]
    mkCharReplacement (replaceStr, searchStr) = T.unpack . replacer (T.pack searchStr) (T.pack replaceStr) . T.pack
    replacer = if forParsing then flip T.replace else T.replace


