**HostGator** 

Brasil,

API - Domain Availability

**September 13, 2022**

# **OVERVIEW**
The purpose of this API is to validate the availability of a domain for registration, using a public endpoint. This document aims to present the requirements and procedures necessary to use the tool, which was developed using the Golang programming language.
# **Requirements**
1. To install Golang (GO), it is necessary that the user has a package manager on his machine, such as:  [Homebrew](https://brew.sh),[ ](https://chocolatey.org)[Chocolatey](https://chocolatey.org),[ ](https://help.ubuntu.com/community/AptGet/Howto)[Apt](https://help.ubuntu.com/community/AptGet/Howto) or[ ](https://access.redhat.com/solutions/9934)[yum](https://access.redhat.com/solutions/9934)
1. Have a source code editor that supports the GO language installed on your machine, such as: [Visual Studio Code](https://code.visualstudio.com/download), [GoLand](https://www.jetbrains.com/go/download/?source=google&medium=cpc&campaign=10156130867&gclid=Cj0KCQiAybaRBhDtARIsAIEG3kkQZeuiFCTmJsk2T-Sz_Ho4VqsyibW-a0i5MEgMlaVej0GQBNUPtREaAiq_EALw_wcB#section=windows), [Komodo](https://www.activestate.com/products/komodo-ide/download-ide/).
# **Introduction**
Below are three different installation ways, which are directed to the three main operating systems currently used.
## **Mac OSX:**

The first step is to install the homebrew (brew). To install it you must have xcode, it can be installed with the following command:

|**xcode-select --install**|
| :- |



Then, run the following command to install homebrew:

|**/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"**|
| :- |

With this, it is now possible to install GO, using the command:

|**brew install go**|
| :- |

## **Linux**
To perform the installation on Linux, it is necessary to download the desired version of GO. Just access the [official website](https://golang.org/) and copy the desired version (for example: <https://dl.google.com/go/go1.13.linux-amd64.tar.gz>).

It is highly recommended to always download the most current version available.

To download, just run the command in the terminal:

|<p>**cd ~**</p><p>**curl -O "https://dl.google.com/go/go${VERSAO\_GO}.linux-amd64.tar.gz"**</p>|
| :- |

Unzip the file with the command:


|**tar xvf "go${VERSAO\_GO}.linux-amd64.tar.gz"**|
| :- |

Move the files to your user's binary directory, with the command:


|**sudo mv go /usr/local**|
| :- |

To verify that the installation was successful, run:


|<p>**go version**</p><p>**go version go1.13 linux/amd64**</p>|
| :- |

## **Windows**
For Windows users, it is possible to perform the installation in two ways: through a ZIP file (it is necessary to configure some environment variables), or using an MSI file that performs the configuration automatically.

The desired version can be downloaded from the language's official website. It is recommended to always download the most current version.


**Installation using MSI**

Run the MSI file and follow the installation steps. The installer adds Go to the folder C:\Go.

The installer adds the path C:\Go\bin in the environment variable "Path" and create the user variable"GOPATH" with the way C:\Users\%USER%\go.
### **Installation via ZIP**
Extract the files to a directory of your choice.

Add in the environment variable "Path" the path to the "bin" folder of the Go files.

Open the Start menu, type "Variables". Then choose the option "Edit system environment variables''.

Click on the "Advanced" > "Environment Variables" tab. Locate the "Path" variable and click Edit > New and fill in the chosen path (Ex: C:/my-folder-go/bin).

# **Explaining the GO environment**

By convention all code is placed in a single folder (workspace). This folder can be created anywhere on the machine.

If the user does not specify this location, $HOME/go will be set as the default workspace. It is identified (and modified) by the environment variable GOPATH. The GOPATH environment variable lists the places to look for the Go code.

You need to set the environment variable so that you can use it in future scripts, shells, etc.

Update your .bash\_profile to contain the following exports.


|<p>export GOPATH=$HOME/go</p><p>export PATH=$PATH:$GOPATH/bin</p>|
| :- |

*Note: You must open a new terminal to set these environment variables.*

*Go assumes that your workspace contains a specific directory structure.*

*It places its files in three directories:*

- The source code is in src,
- Package objects are in pkg
- Compiled programs are placed in bin

You are now able to use go get so that src/package/bin is correctly installed in the appropriate $GOPATH/xxx directory.

# **Conclusion**

These are the procedures required to install GO on your machine.

All information for the elaboration of this README.md was taken from the official website of the Golang language and can be consulted [here](http://www.golangbr.org/doc/instalacao).



