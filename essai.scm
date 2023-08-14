(define l '(1 2 3 4))
(define f (lambda (n) (display n)(newline)))
(do
  (
   [ll l (cdr ll)]
   [nl '() (cons (f (car ll)) nl)]
  )
  ((null? ll) nl)
)
(newline)
(display l)
(newline)
(map f l)

(define rot (lambda (l) (append (cdr l) (cons (car l) '()))))

(define btree
  (lambda (lt t la n)
    (display lt)
    (display t)
    (display la)
    (display n)
    (newline)
    (cond
      ((= (length la) 2) (cons (min (car la) (cadr la)) (append t (cons (max (car la) (cadr la)) '()))))
      ((zero? n) lt)
      (else (btree
	(cons
          (btree
	    '()
	    (append t (cons (car la) '()))
	    (cdr la)
	    (length (cdr la))
	  )
	  lt
	)
	t (rot la) (- n 1)))
    )
  )
)
