import Data.List
import Data.Ord
import System.Environment


main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        let message = map getMax $ transpose . lines $ content where
                getMax = head . maximumBy (comparing length) . group . sort
        putStrLn message
