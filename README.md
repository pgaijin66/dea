

[![Go Report Card](https://goreportcard.com/badge/github.com/pgaijin66/dea)](https://goreportcard.com/report/github.com/liangyaopei/checker)
[![Go Reference](https://pkg.go.dev/badge/github.com/liangyaopei/checker.svg)](https://pkg.go.dev/github.com/liangyaopei/checker)
[![Build Status](https://github.com/pgaijin66/dea/actions/workflows/ci.yml/badge.svg)](https://github.com/pgaijin66/dea/actions/workflows/ci.yml)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Coverage Status](https://coveralls.io/repos/github/pgaijin66/dea/badge.svg?branch=main)](https://coveralls.io/github/pgaijin66/dea/checker?branch=main)


# Table of contents

- [Disposable Email Address ( DEA ) checker](#disposable-email-address--dea--checker)
- [How does `dea` works ?](#who-does-dea-works-)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [FAQs](#faqs)
  - [What are DEA ( Disposable Email Addresses ) ?](#what-are-dea--disposable-email-addresses--)
  - [Why should i use `dea` ?](#why-should-i-use-dea-)
  - [Will using `dea` be suffecient to check for email legitimacy ?](#will-using-dea-be-suffecient-to-check-for-email-legitimacy-)
  - [Will by test gmail account be considered as DEA ?](#will-by-test-gmail-account-be-considered-as-dea-)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Disposable Email Address ( DEA ) checker

A `dea` is a email verification package that can be used to determine if the email address provided in the input is a disposable email address (DEA) or not.

## How does `dea` works ?

`dea` works by checking email legitimacey using a 3 step process:

1. Checking if the email provider is in the list of blocklisted domains.

2. Checking if the `SPF` record for the email provider is legit.

3. Checking if the email addresses that contain random strings of characters or misspelled words 

## Requirements

Go 1.19 or above

## Installation

You can use `go get` to install this package:

```
go get -u github.com/pgaijin66/dea
```

## Usage

First of all you would need to create a 

Here's an example of how to use this package:

```go

package main

import (
    "fmt"
    "github.com/pgaijin66/dea"
)

func main() {
	input := "demo@2wc.info"

	res1, err := dea.IsDisposableEmail(input)
	if err != nil {
		fmt.Printf("%v", err)
	}
	// Check if email provided is a disposable email address
	if res1 {
		fmt.Printf("%v is a disposable email address \n", validEmail)
	}

}

```

## Reference
`IsDisposableEmail(email string) ( bool , error )`

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

### Will using `dea` be suffecient to check for email legitimacy ?

Accourding to my test of using `dea` and 1000+ disposable email address, it was able to reliably detect and flag around 95% of the disposable email addresses. Please be aware that just using `dea` alone will not be suffecient to be assured that user who signed up will be using legitimate email and not disposable one. But, if used along with other features like using activation email, sending random one-off email address after random number of days for legitimacy verification etc would work as a pretty reliable way to detect if the email is a legitimate users email address or a one-off disposable email address.

### Will my test gmail account be considered as DEA ?

Whether a test email account created in Gmail is considered a disposable email address (DEA) depends on the specific context and use case.

If the test email account is created for a legitimate and non-spammy purpose, such as testing a website or application's email functionality, then it would not be considered a DEA. In this case, the email address is being used temporarily and for a specific purpose, but it is not intended to be discarded after that purpose is fulfilled.

However, if the test email account is being used for spamming or other illegitimate purposes, or if it is created with the intent to be used for a short period of time and then discarded, then it would be considered a DEA.

In general, whether an email address is considered a DEA depends on how it is being used and for what purpose. If an email address is being used to deceive, spam, or otherwise abuse a system or service, it is likely to be considered a DEA.

## Contributing

Contributing
Contributions are welcome! If you find any issues or want to enhance the functionality of this package, feel free to open an issue or submit a pull request on GitHub.

Please make sure to follow the existing coding style and add appropriate tests for any new features.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use and modify it according to your needs.

## Acknowledgments
The list of disposable email domains used in this package is sourced from [Disposable Domains](https://raw.githubusercontent.com/disposable-email-domains/disposable-email-domains/master/disposable_email_blocklist.conf).