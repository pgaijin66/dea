
# Table of contents

- [Disposable Email Address ( DEA ) checker](#disposable-email-address--dea--checker)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [FAQs](#faqs)
  - [What are DEA ( Disposable Email Addresses ) ?](#what-are-dea--disposable-email-addresses--)
  - [Why should i use `dea` ?](#why-should-i-use-dea-)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Disposable Email Address ( DEA ) checker

A `dea` is a email verification package that can be used to determine if the email address provided in the input is a disposable email address (DEA) or not. 

## Requirements

Go 1.19 or above

## Installation

You can use `go get` to install this package:

```
go get -u github.com/pgaijin66/dea
```

## Usage

Here's an example of how to use this package:

```go

package main

import (
    "fmt"
    "github.com/pgaijin66/dea"
)

func main() {
	input := "demo@2wc.info"

	res1, _ := dea.IsDisposableEmail(input)

	// Check if email provided is a disposable email address
	if res1 {
		fmt.Printf("%v is a disposable email address \n", validEmail)
	}

}

```

## API Reference
`IsDisposableEmail(email string) bool`

This function takes an email address as input and returns a boolean value indicating whether the email address is a disposable email address or not.

`email` (string): The email address to check.

Returns:

`true` if the email address is a disposable email address.

`false` if the email address is not a disposable email address.


## FAQs



### What are DEA ( Disposable Email Addresses ) ?

Disposable email addresses (DEAs) are temporary email addresses that are created for the purpose of registering for online services, newsletters, or other websites that require email verification. These email addresses are designed to be discarded after a short period of use, typically after the initial registration process is completed.

DEAs are commonly used by spammers and scammers to bypass email filters and deliver unwanted or malicious content. Many organizations and websites block or flag emails from known disposable email domains to reduce the risk of spam or fraud.

### Why should i use `dea` ?

Using a disposable email address (DEA) checker in your project can provide several benefits:


1. Protect against fraud: DEAs are often used by scammers to sign up for fraudulent accounts or services. By checking for DEAs, you can reduce the risk of fraud in your project.

2. Reduce spam: DEAs are commonly used by spammers to bypass email filters and deliver unwanted or malicious content. By checking for DEAs, you can reduce the amount of spam in your project.

3. Enhance data quality: DEAs can provide inaccurate or incomplete data when used in forms or surveys. By checking for DEAs, you can ensure that you are collecting accurate and complete data from your users.


## Contributing

Contributing
Contributions are welcome! If you find any issues or want to enhance the functionality of this package, feel free to open an issue or submit a pull request on GitHub.

Please make sure to follow the existing coding style and add appropriate tests for any new features.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use and modify it according to your needs.

## Acknowledgments
The list of disposable email domains used in this package is sourced from [Disposable Domains](https://raw.githubusercontent.com/disposable-email-domains/disposable-email-domains/master/disposable_email_blocklist.conf).