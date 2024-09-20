## Table of Contents

- [Overview](#overview)
  - [Purpose](#purpose)
  - [Usage](#usage)
  - [Expected Output](#expected-output)
- [Notes](#notes)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Contributing](#contributing)
  - [License](#license)
  - [Changelog](#changelog)
- [Author](#author)

## Statistics Calculator

![Statistics Calculator Logo](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)

### Overview

#### Purpose

The Statistics Calculator is a user-friendly tool developed to efficiently compute essential statistical measures such as average, median, variance, and standard deviation from a dataset provided in a file. This project simplifies statistical analysis, offering a streamlined alternative to manual calculations or complex software.

#### Usage

To use the Statistics Calculator:

1. Prepare a `.txt` file named `data.txt` containing the dataset you wish to analyze. Each line in the file should represent one value of the statistical population, separated by a new line.

    Example:
    ```
    728
    763
    763
    764
    90.23
    74
    738
    ...
    ```

2. Run the program via the command line. Since the program is written in Go, execute the following command:

    ```bash
    $ go run main.go data.txt
    ```

    Replace `main.go` with your program file name and `data.txt` with the path to your dataset file.

#### Expected Output

The program will compute the average, median, variance, and standard deviation of the dataset and display the results in the following format:

```
Average: 560
Median: 738
Variance: 91545
Standard Deviation: 303
```

Note: Values are rounded to integers.

### Notes

#### Installation

No installation is necessary. Clone the repository, navigate to the folder containing the files, and execute the program using provided commands.

#### Configuration

There are no configuration options available for this project.

#### Contributing

Contributions to this project are welcomed. Fork the repository, implement your changes, and submit a pull request.

#### License

This project is licensed under the [MIT License](LICENSE).

#### Changelog

- First release

### Author

<table>
    <tr>
        <td align="center" style="word-wrap: break-word; width: 150px; height: 150px;">
            <a href="https://www.linkedin.com/in/rodgers-kaunda">
                <img src="https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870" width="100" style="border-radius: 50%;" alt="Emmanuel"/>
                <br />
                <sub style="font-size: 14px;"><b>krodgers</b></sub>
            </a>
        </td>
    </tr>
</table>