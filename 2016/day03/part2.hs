import Data.List
import System.Environment

isTriangle :: [Int] -> Bool
isTriangle [a, b, c] = let sum1= b + c
                           sum2 = a + c
                           sum3 = a + b
                           diff1 = abs $ b - c
                           diff2 = abs $ a - c
                           diff3 = abs $ a - b
                       in  (a > diff1) && (a < sum1) &&
                           (b > diff2) && (b < sum2) &&
                           (c > diff3) && (c < sum3)

splitEvery :: Int -> [a] -> [[a]]
splitEvery _ [] = []
splitEvery n xs = as : splitEvery n bs
        where (as,bs) = splitAt n xs

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        print $ sum $ map (length . filter isTriangle)
                (map (splitEvery 3) $ transpose $
                        map (map read . words) (lines content))
