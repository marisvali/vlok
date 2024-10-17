package gamelib

/*
Unit is used in order to get some advantages of floating Points when doing
integer-only calculations.

The problem:
I want to use integers for all my calculations. The reason and the solution are
explained in the ints package. However, this means that the smallest unit of
change is 1. If I have an object at position (10, 5) and I want it at position
(30, 5) I can only move with a minimum speed of 1 on the X axis. Which might be
too fast for my needs.

This problem comes up in all sorts of other places. For example whenever I do a
division I lose some precision because I round to the nearest integer. This can
quickly become a problem.

The solution proposed here:
Use a Unit everywhere instead of 1. So instead of an object at position 10 going
to position 30, it is an object at position (10*Unit, 5*Unit) going to position
(30*Unit, 5*Unit). If this Unit is 1000, the speed can be as low as 1
milli-Unit.

For convenience, you can express 10*Unit by U(10). and 10*milli-Units by MU(10).

This gives me two things:
1. It allows me to think in terms of reasonable dimensions, instead of using
positions like (10000, 5000) and (30000, 5000) everywhere. I can think in terms
of tens, hundreds and thousands and know that I can have subunit values when
I need them. For example: a pixel on my screen can represent a unit. This way
I can easily reason about sprites, their sizes and their positions when I'm
drawing and debugging them.
2. It allows me to experiment with how big a unit needs to be. I currently have
no idea what kind of operations I will need and how much leeway I need to give
myself. int64 is big, but it's not infinite. If I multiply 4 numbers together,
they must all be smaller than 55108, or I overflow. Suddenly int64 isn't that
big. So I need to be careful to have a Unit that's big enough to give me the
precision I want in my computations. But it needs to be small enough so that
I don't overflow in my computations. Since I won't know what my computations
will be until I finish the game, I need this flexibility established from the
start.

A useful way to choose Unit is to predict the largest values one will need for
a variable containing a number of units. Then, predict how many
multiplications we need. For example in order to work with 2D vectors, it is
often necessary to work with the result of multiplying 2 values. So a reasonable
first constraint is how many multiplications will be required?
Let's say we need at most 3 multiplications (only 2 seems very restrictive).
The limit of int64 is 2^63 = 9,223,372,036,854,775,808 ~= 10^18.
Let's say a value has a range in the interval [-vRange, vRange].
That means if we multiply 3 values, we have:
v1 * v2 * v3 = vRange * Unit * vRange * Unit * vRange * Unit
v1 * v2 * v3 = (vRange * Unit)^3
We need v1 * v2 * v3 <= MaxInt64
(vRange * Unit)^3 <= 10^18.
vRange * Unit <= 10^6
Which means:
Unit = 1 		=> vRange in [-1 000 000, 1 000 000]
Unit = 10 		=> vRange in [-100 000, 100 000]
Unit = 100 		=> vRange in [-10 000, 10 000]
Unit = 1000 	=> vRange in [-1 000, 1 000]
Unit = 10000 	=> vRange in [-100, 100]

Unit is primarily intended to express distances in the world of a game.
If the intention is to express the size of the game as 1-to-1 mapping to a
common screen resolution like 1920x1080, this requires at least a vRange of
[-1920, 1920]. It is often useful to spill outside the screen, which suggests
a vRange of [-10 000, 10 000]. Which requires a Unit of 100.
*/

const Unit = 100

func Units(numUnits Int) Int {
	return numUnits.Times(I(Unit))
}

func Milliunits(numMilliunits Int) Int {
	return numMilliunits.Times(I(Unit / 1000))
}

func U(numUnits int) Int {
	return I(numUnits).Times(I(Unit))
}

func UPt(xUnits int, yUnits int) Pt {
	return Pt{I(xUnits).Times(I(Unit)), I(yUnits).Times(I(Unit))}
}

func CU(numUnits int) Int {
	return I(numUnits).Times(I(Unit)).DivBy(I(100))
}
