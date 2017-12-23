import System.Environment

countChars :: String -> Int
countChars ('\"':xs) = 2 + countChars xs
countChars ('\\':xs) = 2 + countChars xs
countChars (x:xs) = 1 + countChars xs
countChars _ = 2

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        print $ sum $ map (\x -> countChars x - length x) (lines content)
