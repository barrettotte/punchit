# punchit
A basic Go package to convert source to IBM 5081 (80 column) punch cards encoded by an IBM 029 keypunch. 

The IBM 029 encoding was introduced around 1964 alongside the IBM System/360 and was used in early FORTRAN, COBOL, and RPG programming.

This small weekend project was mainly to dip my toes into Go and learn more about how punchcards worked.
I'm around 75% sure I got the IBM 029 encoding and other terminology correct. 
Some of this stuff is a little confusing to me.


## Introduction to Punch Cards
```
IBM 5081 Punch Card (80 columns)

       .........1.........2.........3.........4.........5.........6.........7.........8  
      __________________________________________________________________________________ 
     / &-0123456789ABCDEFGHIJKLMNOPQR/STUVWXYZ:#@'="¢.<(+|!$*);¬ ,%_>?                  |
12  /  X           XXXXXXXXX                        XXXXXX                              |
11 |    X                   XXXXXXXXX                     XXXXXX                        |
 0 |     X                           XXXXXXXXX                   XXXXX                  |
 1 |      X        X        X        X                                                  |
 2 |       X        X        X        X       X     X     X                             |
 3 |        X        X        X        X       X     X     X     X                      |
 4 |         X        X        X        X       X     X     X     X                     |
 5 |          X        X        X        X       X     X     X     X                    |
 6 |           X        X        X        X       X     X     X     X                   |
 7 |            X        X        X        X       X     X     X     X                  |
 8 |             X        X        X        X XXXXXXXXXXXXXXXXXX XXXXX                  |
 9 |              X        X        X        X                                          |
   |____________________________________________________________________________________|

```

Each column in the punch card corresponds to an encoded character.
In this example, the punch card was punched using the IBM 029 keypunch with its own encoding method.

Valid characters are ```&-0123456789ABCDEFGHIJKLMNOPQR/STUVWXYZ:#@'="¢.<(+|!$*);¬ ,%_>?```

This character set was used in early FORTRAN, RPG, and COBOL programming (and probably more).

Please read more about the history of punch cards on the [wiki page](https://en.wikipedia.org/wiki/Punched_card).


## Usage
See example of punching a directory of sample files at [examples/main.go](examples/main.go)


## Examples
Examples of generated punchcards for some legacy programming languages.

- [hello.cob](examples/cards/hello.cob.txt) - hello world in COBOL
- [hello.f](examples/cards/hello.f.txt) - hello world in FORTRAN IV
- [hello.rpg](examples/cards/hello.rpg.txt) - hello world in RPG


## References
- https://golang.org/doc/code.html
- [IBM 29 Card Punch Reference Manual (1970)](http://bitsavers.org/pdf/ibm/punchedCard/Keypunch/029/GA24-3332-6_Reference_Manual_Model_29_Card_Punch_Jun70.pdf)
- [Virtual keypunch](https://www.masswerk.at/keypunch/)
- [Punched Card Typography IBM 026, 029, 129](https://www.masswerk.at/misc/card-punch-typography/)
- http://www.columbia.edu/cu/computinghistory/029.html
- https://twobithistory.org/2018/06/23/ibm-029-card-punch.html
- https://homepage.divms.uiowa.edu/~jones/cards/index.html