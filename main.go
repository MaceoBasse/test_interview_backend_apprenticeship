package main

import (
	"fmt"
	"testing"
)

/*
Fix the function "fix_this_function()"

Press RUN to execute the function and follow the result to correct it.

Your test is successful if the test result is set to "PASS"

	TEST Result:
	PASS

Once you succeed this test please send a copy paste by mail of your modified function "fix_this_function()"

Have fun :)
*/

/************************************************* EDIT THIS FUNCTION ONLY *************************************************/
func fix_this_function() (result []interface{}) {
	for x := 1; x <= 17; x++ {
		if x%4 == 0 {
			result = append(result, "Fizz")
		} else if x%6 == 0 {
			result = append(result, "Bozz")
		} else {
			result = append(result, x)
		}
	}

	return
}

/***************************************************************************************************************************/
/*						 		  ______  __    __   _______   ______  __  ___
/*								 /      ||  |  |  | |   ____| /      ||  |/  /
/*								|  ,----'|  |__|  | |  |__   |  ,----'|  '  /
/*								|  |     |   __   | |   __|  |  |     |    <
/*								|  `----.|  |  |  | |  |____ |  `----.|  .  \
/*								 \______||__|  |__| |_______| \______||__|\__\
/*
/*
/*			________  .__       .__  __         .__    __________                   .__        __
/*			\______ \ |__| ____ |__|/  |______  |  |   \______   \ ____  ____  ____ |__|______/  |_
/*			 |    |  \|  |/ ___\|  \   __\__  \ |  |    |       _// __ \/ ___\/ __ \|  \____ \   __\
/*			 |    `   \  / /_/  >  ||  |  / __ \|  |__  |    |   \  ___|  \__\  ___/|  |  |_> >  |
/*			/_______  /__\___  /|__||__| (____  /____/  |____|_  /\___  >___  >___  >__|   __/|__|
/*			        \/  /_____/               \/               \/     \/    \/    \/   |__|
/*
/*
/***************************************************************************************************************************/

/*********************************************** DON'T MODIFY THE CODE BELOW ***********************************************/

func main() {
	fmt.Println("Actual Result  : ", fix_this_function())
	fmt.Println("Expected Result: ", "[1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19 Buzz]")
	fmt.Println("\nTEST Result: ")

	testing.Main(
		nil,
		[]testing.InternalTest{
			{"Test The result length", TestResultLength},
			{"Test naming convention", TestNaming},
			{"Test The multiple of 3", TestFizzReturn},
			{"Test The multiple of 5", TestBuzzReturn},
			{"Test The multiple of 3 and 5", TestFizzBuzzReturn},
		},
		nil, nil,
	)
}

func TestResultLength(t *testing.T) {
	if len(fix_this_function()) != 20 {
		t.Error("your result should have a 20 length")
	}
}

func TestFizzBuzzReturn(t *testing.T) {
	if fix_this_function()[14] != "FizzBuzz" {
		t.Error("if (x%5 == 0 && x%3 == 0) are true the return should be \"FizzBuzz\"")
	}
}

func TestNaming(t *testing.T) {
	for _, v := range fix_this_function() {
		_, ok := v.(int)
		if !ok {
			switch v {
			case "Fizz":
				continue
			case "Buzz":
				continue
			case "FizzBuzz":
				continue
			default:
				t.Errorf("the naming convention is set to \"Fizz\", \"Buzz\" and \"FizzBuzz\"; name Provided : %v", v)
			}
		}
	}
}

func TestFizzReturn(t *testing.T) {
	for i, v := range fix_this_function() {
		_, ok := v.(int)
		if ok {
			x := i + 1
			if x%3 == 0 && v != "Fizz" {
				t.Error("every multiple of 3 should return \"Fizz\"")
			}
		}

	}
}

func TestBuzzReturn(t *testing.T) {
	for i, v := range fix_this_function() {
		_, ok := v.(int)
		if ok {
			x := i + 1
			if x%5 == 0 && v != "Buzz" {
				t.Error("every multiple of 5 should return \"Buzz\"")
			}
		}

	}
}
