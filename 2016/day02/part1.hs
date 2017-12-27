import System.Environment

type Pos = (Int, Int)
type Direction = Char
type Code = Int

move :: Pos -> Direction -> Pos
move (x,y) 'U' = if y == 1 then (x,y) else (x,y+1)
move (x,y) 'R' = if x == 1 then (x,y) else (x+1,y)
move (x,y) 'D' = if y == (-1) then (x,y) else (x,y-1)
move (x,y) 'L' = if x == (-1) then (x,y) else (x-1,y)

initPos :: Pos
initPos = (0,0)

posToCode :: Pos -> Code
posToCode (-1,-1) = 7
posToCode (-1, 0) = 4
posToCode (-1, 1) = 1
posToCode ( 0,-1) = 8
posToCode ( 0, 0) = 5
posToCode ( 0, 1) = 6
posToCode ( 1,-1) = 9
posToCode ( 1, 0) = 6
posToCode ( 1, 1) = 3

main :: IO ()
main = do
        args <- getArgs
        content <- readFile (args !! 0)
        let codes = lines content 
            code1 = foldl move initPos (codes !! 0)
            code2 = foldl move code1   (codes !! 1)
            code3 = foldl move code2   (codes !! 2)
            code4 = foldl move code3   (codes !! 3)
            code5 = foldl move code4   (codes !! 4)
        putStrLn $ concatMap (show.posToCode)
                        [code1, code2, code3, code4, code5]
