# Check BackEnd (apprentice-level) Take Home Exercise
Technical assessment for apprentice-level back-end developers.

# Instructions
In the file 'main.go' Fix the function "fix_this_function()"

Execute the code and correct the code that have failing tests.

Your test is successful if the test result is set to "PASS"

```
TEST Result:
PASS
```

After passing the test, please copy your modified "fix_this_function()" and email it.

This test can be done directly online in the go playground https://go.dev/play/p/J2HzkWWNr21?v=goprev  
`don't forger to send it by mail`

Have fun :)

## How to submit your exercise

- Include in the mail subject this sentence "TEST for apprentice-level back-end " with your name
- Paste your solution in the mail body and send it
- Send your mail to this address team@app-check.fr

## Explanation of the Correction

The `fix_this_function()` needed to be modified to correctly implement the FizzBuzz logic. Here are the key changes made:

1. Extended the loop to iterate from 1 to 20 (inclusive) to match the expected result length.
2. Implemented the correct FizzBuzz logic:
   - For numbers divisible by both 3 and 5, append "FizzBuzz"
   - For numbers divisible by 3, append "Fizz"
   - For numbers divisible by 5, append "Buzz"
   - For other numbers, append the number itself
3. Removed incorrect conditions for multiples of 4 and 6.
4. Corrected the spelling from "Bozz" to "Buzz"

These changes ensure that the function passes all provided tests and correctly implements the FizzBuzz game rules.