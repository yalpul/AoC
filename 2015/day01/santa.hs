santa :: String -> Int
santa ('(':xs) = 1 + santa xs
santa (')':xs) = -1 + santa xs
santa _ = 0
