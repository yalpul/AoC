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

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        print $ length $ filter isTriangle
                (map (map read . words) (lines content))
