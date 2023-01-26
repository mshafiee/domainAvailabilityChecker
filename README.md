# Domain Availability Checker
This code is a command-line tool that allows you to check the availability of a domain name across multiple top-level domains (TLDs). The user provides the domain name as a command-line argument and the tool returns a list of available and registered domains.

## Getting Started
### Prerequisites
 1. Go version 1.13 or higher
### Installing
 2. Clone this repository
 3. Copy code
```sh
$ git clone https://github.com/mshafiee/domainAvailabilityChecker.git
```
 4. Change directory to the repository folder
```sh
$ cd domainAvailabilityChecker
```
 5. Run the code
```sh
$ go run main.go <domain_name>
```
 6. (Optional) If you want to check the availability of the domain for
    specific TLDs, you can pass a comma-separated list of TLDs as a
    second argument.

```sh
$ go run main.go <domain_name> <comma_separated_list_of_TLDs>
```
- Sample:
```sh
$ go run main.go github com,org,net,de,uk,ir

Checking availability for domain 'github' with 6 TLDs
-------------------------------------
Available domains:
        github.uk is available.
        github.ir is available.
-------------------------------------
Registered domains:
        github.com is registered.
        github.org is registered.
        github.net is registered.
        github.de is registered.

```
 7. Run the following command to build the executable binary:
 ```sh
$ go build -o domainAvailabilityChecker
```
 8. This will create a binary called domainAvailabilityChecker in the root directory of the project. You can then move this binary to a directory that is in your system's PATH environment variable, such as /usr/local/bin, so that you can run the tool from anywhere on your system.
 ```sh
$ sudo mv domainAvailabilityChecker /usr/local/bin/
```
