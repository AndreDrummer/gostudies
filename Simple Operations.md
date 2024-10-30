# Go Studies

### `2` GO Practice Challenge: Go Basics Challenge

#### **Challenge 1: Simple Math Operations**

1. **Goal**: Implement a program that performs basic math operations.
2. **Description**: Create a program called `math_operations.go` that takes two integers as inputs and performs addition, subtraction, multiplication, and division on them.

   - **Input**: Prompt the user to enter two numbers.
   - **Output**: Display the result of each operation with clear labels.

3. **Tasks**:
   - Implement the following operations:
     - **Addition**: Sum the two numbers.
     - **Subtraction**: Subtract the second number from the first.
     - **Multiplication**: Multiply the two numbers.
     - **Division**: Divide the first number by the second, but make sure to handle division by zero gracefully.
   - **Example Output**:
     ```
     Enter first number: 10
     Enter second number: 5

     Results:
     10 + 5 = 15
     10 - 5 = 5
     10 * 5 = 50
     10 / 5 = 2
     ```

4. **Bonus**: Use Go’s `fmt.Scan` to capture input from the user, and include error handling for cases like invalid inputs (e.g., strings instead of numbers).

---

#### **Challenge 2: Conditional Statements and Logic**

1. **Goal**: Practice using conditionals by implementing a program that determines the type of a given integer.
2. **Description**: Create a program called `number_type.go` that checks if a given integer is:
   - Positive or Negative
   - Even or Odd

3. **Tasks**:
   - **Input**: Prompt the user to enter a single integer.
   - **Output**: Display whether the number is positive/negative and even/odd.

   - **Example Output**:
     ```
     Enter a number: -4
     
     The number is negative.
     The number is even.
     ```

4. **Bonus**: Implement error handling for invalid inputs and make sure your program handles the case where the number is zero, printing `"The number is zero"` and `"The number is even"`.

---

#### **Challenge 3: Temperature Converter**

1. **Goal**: Create a program that converts temperatures between Celsius and Fahrenheit.
2. **Description**: Create `temperature_converter.go` that asks the user for a temperature value and the conversion direction (Celsius to Fahrenheit or Fahrenheit to Celsius).

3. **Tasks**:
   - **Input**: Prompt the user to enter the temperature value and the conversion choice (e.g., “C” for Celsius to Fahrenheit, “F” for Fahrenheit to Celsius).
   - **Output**: Display the converted temperature with an appropriate label.

4. **Conversion Formulas**:
   - Celsius to Fahrenheit: \( F = C \times \frac{9}{5} + 32 \)
   - Fahrenheit to Celsius: \( C = (F - 32) \times \frac{5}{9} \)

5. **Example Output**:
   ```
   Enter temperature value: 100
   Convert from (C/F): F

   100°F is 37.78°C
   ```

6. **Bonus**: Add input validation so that only valid numbers and conversion choices ("C" or "F") are accepted.

---

#### **Challenge 4: Leap Year Checker**

1. **Goal**: Write a program to determine if a given year is a leap year.
2. **Description**: Create a program called `leap_year_checker.go` that determines whether a year is a leap year.

3. **Leap Year Rules**:
   - A year is a leap year if it is divisible by 4, but not by 100, except if it is also divisible by 400.

4. **Tasks**:
   - **Input**: Prompt the user to enter a year.
   - **Output**: Display whether the given year is a leap year or not.

5. **Example Output**:
   ```
   Enter a year: 2000
   2000 is a leap year.
   ```

6. **Bonus**: Add error handling for invalid input (e.g., non-integer values) and handle negative years.