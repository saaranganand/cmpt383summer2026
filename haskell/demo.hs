-- demo.hs

-- equational definiton
middle :: [a] -> [a]
middle []     = []
middle [_]    = []
middle (_:xs) = reverse (tail (reverse xs))

pluralize :: String -> String
pluralize "" = ""
pluralize x  = x ++ (if last x == 's' 
                     then "" 
                     else "s")

-- guarded equations
sign n | n == 0    = "zero"
       | n < 0     = "negative"
       | otherwise = "positive"

coin :: Int -> String
coin 1  = "penny"
coin 5  = "nickel"
coin 10 = "dime"
coin 25 = "quarter"
coin _  = "unknown"

nand :: Bool -> Bool -> Bool
nand True True = False
nand _    _    = True 
-- nand False False = True
-- nand False True  = True
-- nand True  False = True
-- nand True  True  = False
-- truth table 

get_expr :: [String] -> String
get_expr [_,"+",_] = "addition"
get_expr [_,"-",_] = "subtraction"
get_expr _         = "unknown"

firsts :: [a] -> [b] -> (a, b)
firsts (x:_) (y:_) = (x, y)
