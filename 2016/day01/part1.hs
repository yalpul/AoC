import System.Environment

type Pos = (Char, Int, Int)
type Direction = (Char, Int)

initPos :: Pos
initPos = ('N', 0, 0)

move :: Pos -> Direction -> Pos
move ('N', x, y) ('R', num) = ('E', x+num, y)
move ('N', x, y) ('L', num) = ('W', x-num, y)

move ('S', x, y) ('R', num) = ('W', x-num, y)
move ('S', x, y) ('L', num) = ('E', x+num, y)

move ('W', x, y) ('R', num) = ('N', x, y+num)
move ('W', x, y) ('L', num) = ('S', x, y-num)

move ('E', x, y) ('R', num) = ('S', x, y-num)
move ('E', x, y) ('L', num) = ('N', x, y+num)

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        let (_,x,y) = foldl move initPos $
                map ((\(x:xs) -> (x, read xs)) . init)
                    (words $ init content ++ ",")
        print $ abs x + abs y
