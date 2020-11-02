# punchit
A basic Go package to convert to/from IBM 5081 (80 column) punch cards encoded by an IBM 029 keypunch. 

The IBM 029 encoding was introduced around 1964 alongside the IBM System/360 and was used in early FORTRAN, COBOL, and RPG programming.

This small weekend project was mainly to dip my toes into Go and learn more about how punchcards worked.
I'm around 75% sure I got the IBM 029 encoding and other terminology correct. 
Some of this stuff is a little confusing to me.


## Features
- Convert source to punchcards (serialize)
- Convert punchcards to source (deserialize)


## Introduction to Punch Cards
```
IBM 5081 Punch Card (80 columns)

      .........1.........2.........3.........4.........5.........6.........7.........8  
      _________________________________________________________________________________ 
     / &-0123456789ABCDEFGHIJKLMNOPQR/STUVWXYZ:#@'="¢.<(+|!$*);^ ,%_>?                 |
12  /  ▮           ▮▮▮▮▮▮▮▮▮                        ▮▮▮▮▮▮                             |
11 |    ▮                   ▮▮▮▮▮▮▮▮▮                     ▮▮▮▮▮▮                       |
 0 |     ▮                           ▮▮▮▮▮▮▮▮▮                  ▮▮▮▮▮▮                 |
 1 |      ▮        ▮        ▮        ▮                                                 |
 2 |       ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮     ▮                |
 3 |        ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮                     |
 4 |         ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮                    |
 5 |          ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮                   |
 6 |           ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮                  |
 7 |            ▮        ▮        ▮        ▮       ▮     ▮     ▮     ▮                 |
 8 |             ▮        ▮        ▮        ▮ ▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮                 |
 9 |              ▮        ▮        ▮        ▮                                         |
   |___________________________________________________________________________________|

```

Each column in the punch card corresponds to an encoded character.
In this example, the punch card was punched using the IBM 029 keypunch with its own encoding method.

Valid characters are ```&-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ&-/:#@'="¢.<(+|!$*);¬ ,%_>?```

This character was used in early FORTRAN, RPG, and COBOL programming (and probably more).

Please read more about punch cards on the [wiki page](https://en.wikipedia.org/wiki/Punched_card).


## Examples

### RPG
```RPG
     C* HELLO WORLD IN RPG                             
     C                     MOVEL'HELLO'   HELLO  11    
     C                     MOVE 'WORLD'   HELLO        
     C           HELLO     DSPLY                       
     C                     SETON                     LR
```

### FORTRAN IV
```fortran
C     HELLO WORLD IN FORTRAN IV
      PROGRAM HELLO
      WRITE (*,100)
      STOP
  100 FORMAT (' HELLO WORLD! ' /)
      END
```

### COBOL
```COBOL
      * HELLO WORLD IN COBOL 
       IDENTIFICATION DIVISION.
       PROGRAM-ID. HELLO.
       ENVIRONMENT DIVISION.
       DATA DIVISION.
       PROCEDURE DIVISION.
           DISPLAY "HELLO WORLD"
           STOP RUN.
```


## References
- https://golang.org/doc/code.html
- [IBM 29 Card Punch Reference Manual (1970)](http://bitsavers.org/pdf/ibm/punchedCard/Keypunch/029/GA24-3332-6_Reference_Manual_Model_29_Card_Punch_Jun70.pdf)
- [Virtual keypunch](https://www.masswerk.at/keypunch/)
- http://www.columbia.edu/cu/computinghistory/029.html
- https://twobithistory.org/2018/06/23/ibm-029-card-punch.html
- https://homepage.divms.uiowa.edu/~jones/cards/index.html
