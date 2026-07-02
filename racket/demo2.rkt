#lang racket

;; returns the length of list L
(define (len L)
  (if (empty? L)
      0
      (+ 1 (len (rest L)))))

#;(define (contains x lst)
    (cond [test1 value1]
          [test2 value2]
          [else value3]
          )
    )

;; returns #t if one or more top-level items in lst satisfy
;; pred?, and #f otherwise; pred? is assumed to be a predicate,
;; i.e. a function that takes one input and returns either #t
;; or #f
(define (satisfies? pred? lst)
  (cond [(empty? lst)
         #f]
        [(pred? (first lst))
         #t]
        [else
         (satisfies? pred? (rest lst))]))

#;(define (contains? x lst)
    (cond [(empty? lst)
           #f]
          [(equal? x (first lst))
           #t]
          [else
           (contains? x (rest lst))]))

;; returns #t if x is lst (at the top-level),
;; and #f otherwise
(define (contains? x lst)
  (satisfies? (lambda (a) (equal? x a))
              lst))

#;(define (count-sym lst)
    (cond [(empty? lst)
           0]
          [(symbol? (first lst))
           (+ 1 (count-sym (rest lst)))]
          [else
           (count-sym (rest lst))]))

;; returns the number of top-level symbols in lst
(define (count-sym lst)
  (if (empty? lst)
      0
      (+ (if (symbol? (first lst)) 1 0)
         (count-sym (rest lst)))))

;; returns the number of top-level items on lst that satisfy pred?
(define (count-fn pred? lst)
  (cond [(empty? lst)
         0]
        [(pred? (first lst))
         (+ 1 (count-fn pred? (rest lst)))]
        [else
         (count-fn pred? (rest lst))]))

;; reverses list lst
(define (rev lst)
  (if (empty? lst)
      '()
      (append (rev (rest lst))
              (list (first lst)))))

;; appends list A and B
(define (my-append A B)
  (cond [(empty? A)
         B]
        [else
         (cons (first A)
               (my-append (rest A) B))
         ]))

;; returns the number of numbers that occur in the top-level
;; of lst
(define (deep-count-num lst)       
  (cond [(empty? lst) 
         0]
        [(list? (first lst))
         (+ (deep-count-num (first lst)) 
            (deep-count-num (rest lst)))]
        [(number? (first lst))
         (+ 1 (deep-count-num (rest lst)))]
        [else
         (deep-count-num (rest lst))]
        ))

;; returns a list of the non-list items in x, e.g.
;; (flatten '((a) (b (c) (d e)) f)) returns '(a b c d e f)
(define (my-flatten x)
  (cond [(empty? x) 
         x]
        [(not (list? x)) 
         x]
        [(list? (first x))
         (append (my-flatten (first x))
                 (my-flatten (rest x)))]
        [else ;; first element is not a list
         (cons (first x) 
               (my-flatten (rest x)))]
        ))

(define (my-map f lst)
  (cond [(empty? lst)
         '()]
        [else
         (cons (f (first lst))
               (my-map f (rest lst)))]))

(define (my-filter pred? lst)
  (cond [(empty? lst)
         '()]
        [(pred? (first lst))
         (cons (first lst)
               (my-filter pred? (rest lst)))]
        [else
         (my-filter pred? (rest lst))]))
  