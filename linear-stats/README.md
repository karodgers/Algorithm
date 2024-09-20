📊 Linear Stats

This project is a Go-based application that reads numerical data from a file and computes key statistical measures: the Linear Regression Line and the Pearson Correlation Coefficient. The program then outputs these calculations based on the requisite decimal precision.

🚀 Features

    Linear Regression Line: Calculates and displays the equation of the best-fit line.
    Pearson Correlation Coefficient: Computes the degree of linear relationship between the dataset points.
    File Input: Reads data directly from a specified file, allowing for flexible data handling.

📂 File Format

The input data should be stored in a text file where each line represents a y-value corresponding to its line number, which serves as the x-value.

Example:

    189
    113
    121
    114
    145
    110
    ...

📈 Calculations
Linear Regression Line

The Linear Regression Line is expressed as:
    y = mx + c


Where:

    y = Dependent variable (predicted value)
    x = Independent variable (line number in the file)
    m = Slope of the line.

The slope is calculated using the formula
    
    m = n(∑xy)−(∑x)(∑y) / n(∑x2)−(∑x)2

And the intercept is calculated as
   
    c=(∑y)(∑x2)−(∑x)(∑xy) / n(∑x2)−(∑x)2

Pearson Correlation Coefficient

The Pearson Correlation Coefficient, r, measures the strength and direction of the linear relationship between two variables. It is calculated by:
    r = n(∑xy)−(∑x)(∑y)/ sqrt[n(∑x2)−(∑x)2][n(∑y2)−(∑y)2]


Where:

    r = Pearson Correlation Coefficient (range: -1 to 1)
    n = Number of data points
    x = Independent variable (line number)
    y = Dependent variable (value in the file)

💻 Usage

To run the program, use the following command:

    $ go run main.go data.txt


🖥️ Output

After processing the data, the program will output:


    Linear Regression Line: y = <value>x + <value>
    Pearson Correlation Coefficient: <value>

Linear Regression Line: The equation of the best-fit line with 6 decimal places of precision.
Pearson Correlation Coefficient: The correlation value with 10 decimal places of precision.

This project demonstrates the power of Go in handling statistical calculations efficiently, making it a valuable tool for anyone dealing with numerical data and trends analysis.