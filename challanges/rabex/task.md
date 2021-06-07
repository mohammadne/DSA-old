We have an imaginary language called "Xulu".
In this language, each verb comes after a set of names and
all the names have English equivalences in numbers.

Alphabets in *Xulu* are equivalent to numbers as well.
- a = 1
- b = 2
- c = 3
- d = 4
- e = 5

Xulu has only five alphabets, say [a,b,c,d,e]

* Names: 

A composite name in *Xulu* is being composed as follows:
(sum of each repeated number mod 5 )^2 

example: 
     aabbcccca
            - break this number to its repeated alphabets:
                          aa , bb , cccc , a
            - now compute its equivalents for every part 
                          aa = 1 + 1 = 2
                          bb = 2 + 2 = 4
                          cccc = 3 + 3 + 3 + 3 = 12
                          a = 1
           - compute mod 5 of each set : 
                         aa mod 5 = 2
                         bb mod 5 = 4
                         cccc mod 5 = 2
                         a mod 5 = 1
          - Now sqaure each number and add them up:
                         (2^2) + (4^2) + 2^2 +1 ^ 2 = 25

* Verbs: 

Xulu has only three verbs that can be translated to a mathematical operation:

abcd = addition
bcde = subtraction 
dede = multiplication
abab = division

* Sentence :
Any repetition of verbs and names in Xulu forms a sentence. A sentence begins with a verb.
A verb can be followed by another sentence to form a composite sentence:

* Grammar: 

The Xulu's grammar is not so difficult, in the grammar, 
the verbs come before the names and the names can be in any quantity right after the verb. 
Normally, two verbs do not happen to place after each other, if so, it means they are parts of separate sentences:
      example: 
                abcd bcde ab ac abab a b it means :
                abcd (bcde ab ac) (abab a b) = sentence 1 + sentence 2 

The Xulu grammar converts any sentences into a number :
     example:
               abcd abcd aabbc ab a c ccd dede cccd cd  
               The above sentence is equvalent to :
                     (29 + 5 + 1 + 9 + 17)  +  ( 32*25 ) = 61 + 800 = 861


A sample test set is here: https://gist.github.com/d-fal/a40984e381a4a56300ffa6cbb7f5c82f
