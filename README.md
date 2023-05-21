# cbison

cbison is a command line tool that interacts with Google's Generative AI to generate responses for given prompts using the chat-bison-001 model. The tool allows users to send a message to the AI model and receive a generated reply.

## Prerequisites

Ensure that you have the following installed in your environment:

* [Go](https://golang.org/doc/install)

## Usage

First, clone the repository and navigate to the directory:

```bash
git clone https://github.com/tmc/cbison.git
cd cbison
```

Next, install the binary:

```bash
go install
```

Make sure you have a valid Palm API key: https://makersuite.google.com/app/apikey

```bash
export GOOGLE_GENERATIVE_AI_API_KEY=your-api-key
```

To use cbison, prepare your input as text file:

```bash
echo "Hello, how are you?" > input.txt
```

Then, run the command to generate a reply:

```bash
cat input.txt | cbison
```

Note: Ensure you use the `cbison` command (with the dot and slash) to run the binary built in the earlier step. Replace `cbison` with the binary filename (e.g. `cbison.exe` on Windows) as needed.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

