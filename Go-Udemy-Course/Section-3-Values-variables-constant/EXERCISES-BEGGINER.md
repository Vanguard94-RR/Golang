# Exercises - Section 3: Values, Variables and Constants

## Exercise 1: Temperature Conversion

**Description:**
Create a program that converts temperature from Celsius to Fahrenheit. Use constants for conversion values (C to F: multiply by 9/5 and add 32).

**Expected Input and Output:**

```
Celsius: 0°C → Fahrenheit: 32°F
Celsius: 25°C → Fahrenheit: 77°F
Celsius: 100°C → Fahrenheit: 212°F
```

**Hint:**
Define the formula as constants. Declare temperatures as variables and print the results.

---

## Exercise 2: BMI Calculator (Body Mass Index)

**Description:**
Program that calculates BMI given weight (kg) and height (m). Formula: BMI = weight / (height²).

**Expected Input and Output:**

```
Weight: 70 kg, Height: 1.75 m → BMI: 22.86
Weight: 80 kg, Height: 1.80 m → BMI: 24.69
```

**Hint:**
Declare variables for weight and height. Multiply height by itself to get the square, or use `math.Pow()`.

---

## Exercise 3: Time Period in Seconds

**Description:**
Create a program that calculates how many seconds are in a given time period (hours, minutes, seconds). Use constants for conversions.

**Expected Input and Output:**

```
2 hours, 30 minutes, 45 seconds → 9045 total seconds
1 hour, 0 minutes, 0 seconds → 3600 total seconds
0 hours, 5 minutes, 30 seconds → 330 total seconds
```

**Hint:**
Define constants: SecondsPerMinute = 60, SecondsPerHour = 3600. Declare variables for each unit and sum the totals.
