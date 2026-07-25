-- demo.hs

-- import Data.List (sort)

-- -- equational definiton
-- middle :: [a] -> [a]
-- middle []     = []
-- middle [_]    = []
-- middle (_:xs) = reverse (tail (reverse xs))

-- pluralize :: String -> String
-- pluralize "" = ""
-- pluralize x  = x ++ (if last x == 's' 
--                      then "" 
--                      else "s")

-- -- guarded equations
-- sign n | n == 0    = "zero"
--        | n < 0     = "negative"
--        | otherwise = "positive"

-- coin :: Int -> String
-- coin 1  = "penny"
-- coin 5  = "nickel"
-- coin 10 = "dime"
-- coin 25 = "quarter"
-- coin _  = "unknown"

-- nand :: Bool -> Bool -> Bool
-- nand True True = False
-- nand _    _    = True 
-- nand False False = True
-- nand False True  = True
-- nand True  False = True
-- nand True  True  = False
-- truth table 

-- get_expr :: [String] -> String
-- get_expr [_,"+",_] = "addition"
-- get_expr [_,"-",_] = "subtraction"
-- get_expr _         = "unknown"

-- firsts :: [a] -> [b] -> (a, b)
-- firsts (x:_) (y:_) = (x, y)


-- twice f x = f (f x)

-- makeAdder :: Int -> (Int -> Int)
-- makeAdder n = (+) n

-- makeAdder' :: Int -> Int -> Int
-- makeAdder' n = (+) n

-- myfilter _ []     = []
-- myfilter p (x:xs) = if p x
--                     then x : (filter p xs)
--                     else filter p xs

-- mytakeWhile _ []     = []
-- mytakeWhile p (x:xs) = if p x
--                        then x:(mytakeWhile p xs)
--                        else []

-- mysum :: Num a => [a] -> a
-- mysum = foldr (+) 0  -- pointfree style
-- mysum []     = 0
-- mysum (x:xs) = x + mysum xs

-- myprod :: Num a => [a] -> a
-- myprod = foldr (*) 1
-- myprod []     = 1
-- myprod (x:xs) = x * myprod xs

-- mymap :: (a -> b) -> [a] -> [b]
mymap f xs = foldr op [] xs
             where op x accum = (f x) : accum

myreverse :: [a] -> [a]
myreverse xs = foldr op [] xs
               where op x accum = accum ++ [x]

myfilter :: (a -> Bool) -> [a] -> [a]
myfilter p xs = foldr op [] xs
                where op x accum = if p x
                                   then x : accum
                                   else accum

f n = n^2
g n = 3*n+1
h = f . g

twice :: (a -> a) -> a -> a
-- twice f x = f (f x)
-- twice f x = (f . f) x
twice f = f . f

sumSquaresEven = sum . map (^2) . filter even

-- do f, then g, then h
-- h . g . f

type Predicate a = a -> Bool

remove_if :: Predicate a -> [a] -> [a]
remove_if p lst = filter (not . p) lst

-- data MyBool = True | False
-- my_and :: MyBool -> MyBool -> MyBool
-- my_and False False = False
-- my_and False True = False
-- my_and True False = False
-- my_and True True = False

data Light = Red | Yellow | Green
  deriving (Eq, Show)

change :: Light -> Light
change Red    = Green
change Yellow = Red
change Green  = Yellow

-- sum type
data Shape = Circle Float | Rect Float Float
    deriving Show

area :: Shape -> Float
area (Circle r) = pi * r^2
area (Rect w h) = w * h

perimeter :: Shape -> Float
perimeter (Circle r) = 2 * pi * r
perimeter (Rect w h) = 2 * (w + h)

-- data Maybe a = Nothing | Just a

safeHead :: [a] -> Maybe a
safeHead []    = Nothing
safeHead (x:_) = x
