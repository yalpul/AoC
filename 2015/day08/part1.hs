import System.Environment

cntChars :: String -> Int
cntChars ('\"':xs) = cntChars xs
cntChars ('\\':y:xs) = case y of '\\' ->  1 + cntChars xs
                                 '\"' ->  1 + cntChars xs
                                 'x'  -> -1 + cntChars xs
cntChars (x:xs) = 1 + cntChars xs
cntChars _ = 0

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        print $ sum $ map (\x -> length x - cntChars x) (lines content)
