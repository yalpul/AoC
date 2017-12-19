first :: String -> Int
first str = first' str "" 0 where
        first' ('(':xs) stack n = first' xs ('(':stack) (n+1)
        first' (')':xs) ('(':stack) n = first' xs stack (n+1)
        first' (')':_) _ n = n+1
