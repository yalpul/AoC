import Data.List

lookAndSay :: String -> String
lookAndSay x = concatMap (\x -> show (length x) ++ [head x]) $ group x

main :: IO ()
main = print $ length $ foldr (.) id (replicate 50 lookAndSay) "1113122113"
