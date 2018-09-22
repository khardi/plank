package lang

import "testing"

//
// Num
//
func Test_Num_String(t *testing.T) {
	want := "5"
	x := Num(5)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Value(t *testing.T) {
	want := 5
	x := Num(5)

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Equal(t *testing.T) {
	want := true
	a := Num(5)
	b := Num(5)

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Foldable(t *testing.T) {
	want := false
	x := Num(5)

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Num_Fold(t *testing.T) {
	want := Num(5)
	x := Num(5)

	got := x.Fold()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Sum
//
func Test_Sum_String(t *testing.T) {
	want := "5 + 2"
	x := Sum(Num(5), Num(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Value(t *testing.T) {
	want := 7
	x := Sum(Num(5), Num(2))

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Equal(t *testing.T) {
	want := true
	a := Sum(Num(5), Num(2))
	b := Sum(Num(5), Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Foldable(t *testing.T) {
	want := true
	x := Sum(Num(5), Num(2))

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Nums(t *testing.T) {
	want := Num(7)
	x := Sum(Num(5), Num(2))

	got := x.Fold()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left(t *testing.T) {
	want := "7 + 3"
	x := Sum(
		Sum(Num(5), Num(2)),
		Num(3),
	).Fold()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_Value(t *testing.T) {
	want := 10
	x := Sum(
		Sum(Num(5), Num(2)),
		Num(3),
	).Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Right(t *testing.T) {
	want := "3 + 7"
	x := Sum(
		Num(3),
		Sum(Num(5), Num(2)),
	).Fold()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Right_Value(t *testing.T) {
	want := 10
	x := Sum(
		Num(3),
		Sum(Num(5), Num(2)),
	).Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_And_Right(t *testing.T) {
	want := "7 + 6"
	x := Sum(
		Sum(Num(5), Num(2)),
		Sum(Num(2), Num(4)),
	).Fold().Fold()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Sum_Fold_Left_And_Right_Value(t *testing.T) {
	want := 13
	x := Sum(
		Sum(Num(5), Num(2)),
		Sum(Num(2), Num(4)),
	).Fold().Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Mul
//
func Test_Mul_String(t *testing.T) {
	want := "5 * 2"
	x := Mul(Num(5), Num(2))

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Value(t *testing.T) {
	want := 10
	x := Mul(Num(5), Num(2))

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Equal(t *testing.T) {
	want := true
	a := Mul(Num(5), Num(2))
	b := Mul(Num(5), Num(2))

	got := a.Equal(b)

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Foldable(t *testing.T) {
	want := true
	x := Mul(Num(5), Num(2))

	got := x.Foldable()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Nums(t *testing.T) {
	want := Num(10)
	x := Mul(Num(5), Num(2))

	got := x.Fold()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left(t *testing.T) {
	want := "10 * 2"
	x := Mul(
		Mul(Num(5), Num(2)),
		Num(2),
	).Fold()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_Value(t *testing.T) {
	want := 20
	x := Mul(
		Mul(Num(5), Num(2)),
		Num(2),
	).Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Righ(t *testing.T) {
	want := Num(20)
	x := Mul(
		Num(2),
		Mul(Num(5), Num(2)),
	)

	got := x.Fold()

	if !got.Equal(want) {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Righ_Value(t *testing.T) {
	want := 20
	x := Mul(
		Num(2),
		Mul(Num(5), Num(2)),
	).Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_And_Righ(t *testing.T) {
	want := "12 * 10"
	x := Mul(
		Mul(Num(4), Num(3)),
		Mul(Num(5), Num(2)),
	).Fold().Fold()

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Mul_Fold_Left_And_Righ_Value(t *testing.T) {
	want := 120
	x := Mul(
		Mul(Num(4), Num(3)),
		Mul(Num(5), Num(2)),
	).Fold().Fold()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

//
// Simple Expression
//
func Test_Simple_Expression_String(t *testing.T) {
	want := "5 * 2 + 3 * 6"
	x := Sum(
		Mul(Num(5), Num(2)),
		Mul(Num(3), Num(6)),
	)

	got := x.String()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}

func Test_Simple_Expression_Fold(t *testing.T) {
	strigns := []string{
		"5 + 2 + 3 * 6",
		"7 + 3 * 6",
		"7 + 18",
		"25",
		"25",
	}

	x := Sum(
		Sum(Num(5), Num(2)),
		Mul(Num(3), Num(6)),
	)

	expstrings := []string{
		x.String(),
		x.Fold().String(),
		x.Fold().Fold().String(),
		x.Fold().Fold().Fold().String(),
		x.Fold().Fold().Fold().Fold().String(), // Num
	}

	for i, want := range strigns {
		got := expstrings[i]
		if got != want {
			t.Fatalf("\n got : %v \n want: %v", got, want)
		}
	}
}

//
// Machine
//
func Test_Machine_Fold(t *testing.T) {
	want := 25

	x := Machine(
		Sum(
			Sum(Num(5), Num(2)),
			Mul(Num(3), Num(6)),
		),
	).Run()

	got := x.Value()

	if got != want {
		t.Fatalf("\n got : %v \n want: %v", got, want)
	}
}
