# Guess-It-2


## 🎯 Objectives

The Guess-It-2 Project is a Go-based application designed to predict a range for the next input number using statistical analysis (Linear Regression). The program reads numerical data sequentially and calculates a probable range for the subsequent number based on previous inputs.


## 📜 How It Works

The application operates by processing a sequence of numbers from the standard input. Each input corresponds to the y-value on a graph where line numbers serve as the x-values. The goal is to predict the range in which the next input number will likely fall.

Program Flow

Here's an example interaction with the program:


    189      # standard input
    120 200  # predicted range for the next input (113)
    113      # standard input
    160 230  # predicted range for the next input (121)
    121      # standard input
    110 140  # predicted range for the next input (114)


The ranges are dynamically calculated.

## 🧪 Testing

The program's performance is evaluated based on its ability to accurately predict ranges with minimal range. For each correct range prediction, a score is awarded — smaller ranges yield higher scores.

To teat the program, a test script can be downloaded from here (https://assets.01-edu.org/guess-it/guess-it-dockerized.zip) which can then be run to verify the correctness of the range predictions. The script runs the program with predefined datasets and checks the output against expected results. You can comapare the guesses by comparing the percentage of correct guesses.


## ⚙️ Setup

Prerequisites

Ensure Go is installed on your system. 

    Clone the repository:
    git clone https://learn.zone01kisumu.ke/git/krodgers/guess-it-2
    cd guess-it-2

You can confirm the programs functionality on the terminal by running 
    go run main.go


Or running the program within the provided script with random data sets.

## 🚀 Usage

To use the program, run it from the command line or run it within the provided script with random data sets. 


## 🔍 Examples

Below is an example demonstrating the application’s functionality:

    $ 
    189      # input
    120 200  # predicted range
    113      # input
    160 230  # predicted range
    121      # input
    110 140  # predicted range
    ...

These predictions are generated using statistical calculations, ensuring that each range is as precise as possible given the data.

## 🌐 Chosen Language

This project is implemented in Go.

## 📄 License

This project is licensed under the MIT License. For more information, see the LICENSE file.

## ✍️ Author

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/rodgers-kaunda>
            <img src=https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=krodgers/>
            <br />
            <sub style="font-size:14px"><b>krodgers</b></sub>
        </a>
    </td>
</tr>
</table>